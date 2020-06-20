// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// package qn_snotify provides a platform-independent interface for file system notifications.
package qn_snotify

import (
	"errors"
	"fmt"
	"time"

	"github.com/qnsoft/common/container/qn_set"
	"github.com/qnsoft/common/internal/intlog"

	"github.com/fsnotify/fsnotify"
	"github.com/qnsoft/common/container/gqueue"
	"github.com/qnsoft/common/container/qn_list"
	"github.com/qnsoft/common/container/qn_map"
	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/os/qn_cache"
)

// Watcher is the monitor for file changes.
type Watcher struct {
	watcher   *fsnotify.Watcher // Underlying fsnotify object.
	events    *gqueue.Queue     // Used for internal event management.
	cache     *qn_cache.Cache   // Used for repeated event filter.
	nameSet   *qn_set.StrSet    // Used for AddOnce feature.
	callbacks *qn_map.StrAnyMap // Path(file/folder) to callbacks mapping.
	closeChan chan struct{}     // Used for watcher closing notification.
}

// Callback is the callback function for Watcher.
type Callback struct {
	Id        int                // Unique id for callback object.
	Func      func(event *Event) // Callback function.
	Path      string             // Bound file path (absolute).
	name      string             // Registered name for AddOnce.
	elem      *qn_list.Element   // Element in the callbacks of watcher.
	recursive bool               // Is bound to path recursively or not.
}

// Event is the event produced by underlying fsnotify.
type Event struct {
	event   fsnotify.Event // Underlying event.
	Path    string         // Absolute file path.
	Op      Op             // File operation.
	Watcher *Watcher       // Parent watcher.
}

// Op is the bits union for file operations.
type Op uint32

const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)

const (
	repeatEventFilterDuration = time.Millisecond // Duration for repeated event filter.
	callbackExitEventPanicStr = "exit"           // Custom exit event for internal usage.
)

var (
	defaultWatcher      *Watcher                    // Default watcher.
	callbackIdMap       = qn_map.NewIntAnyMap(true) // Id to callback mapping.
	callbackIdGenerator = qn_type.NewInt()          // Atomic id generator for callback.
)

func init() {
	var err error
	defaultWatcher, err = New()
	if err != nil {
		// Default watcher object must be created, or else it panics.
		panic(fmt.Sprintf(`creating default fsnotify watcher failed: %s`, err.Error()))
	}
}

// New creates and returns a new watcher.
// Note that the watcher number is limited by the file handle setting of the system.
// Eg: fs.inotify.max_user_instances system variable in linux systems.
func New() (*Watcher, error) {
	w := &Watcher{
		cache:     qn_cache.New(),
		events:    gqueue.New(),
		nameSet:   qn_set.NewStrSet(true),
		closeChan: make(chan struct{}),
		callbacks: qn_map.NewStrAnyMap(true),
	}
	if watcher, err := fsnotify.NewWatcher(); err == nil {
		w.watcher = watcher
	} else {
		intlog.Printf("New watcher failed: %v", err)
		return nil, err
	}
	w.startWatchLoop()
	w.startEventLoop()
	return w, nil
}

// Add monitors <path> using default watcher with callback function <callbackFunc>.
// The optional parameter <recursive> specifies whether monitoring the <path> recursively, which is true in default.
func Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	return defaultWatcher.Add(path, callbackFunc, recursive...)
}

// AddOnce monitors <path> using default watcher with callback function <callbackFunc> only once using unique name <name>.
// If AddOnce is called multiple times with the same <name> parameter, <path> is only added to monitor once. It returns error
// if it's called twice with the same <name>.
//
// The optional parameter <recursive> specifies whether monitoring the <path> recursively, which is true in default.
func AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	return defaultWatcher.AddOnce(name, path, callbackFunc, recursive...)
}

// Remove removes all monitoring callbacks of given <path> from watcher recursively.
func Remove(path string) error {
	return defaultWatcher.Remove(path)
}

// RemoveCallback removes specified callback with given id from watcher.
func RemoveCallback(callbackId int) error {
	callback := (*Callback)(nil)
	if r := callbackIdMap.Get(callbackId); r != nil {
		callback = r.(*Callback)
	}
	if callback == nil {
		return errors.New(fmt.Sprintf(`callback for id %d not found`, callbackId))
	}
	defaultWatcher.RemoveCallback(callbackId)
	return nil
}

// Exit is only used in the callback function, which can be used to remove current callback
// of itself from the watcher.
func Exit() {
	panic(callbackExitEventPanicStr)
}
