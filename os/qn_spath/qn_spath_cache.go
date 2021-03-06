// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// package qn_spath implements file index and search for folders.
//

package qn_spath

import (
	"runtime"
	"strings"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_snotify"
)

// updateCacheByPath adds all files under <path> recursively.
func (sp *SPath) updateCacheByPath(path string) {
	if sp.cache == nil {
		return
	}
	sp.addToCache(path, path)
}

// formatCacheName formats <name> with following rules:
// 1. The separator is unified to char '/'.
// 2. The name should be started with '/' (similar as HTTP URI).
func (sp *SPath) formatCacheName(name string) string {
	if runtime.GOOS != "linux" {
		name = qn.str.Replace(name, "\\", "/")
	}
	return "/" + strings.Trim(name, "./")
}

// nameFromPath converts <filePath> to cache name.
func (sp *SPath) nameFromPath(filePath, rootPath string) string {
	name := qn.str.Replace(filePath, rootPath, "")
	name = sp.formatCacheName(name)
	return name
}

// makeCacheValue formats <filePath> to cache value.
func (sp *SPath) makeCacheValue(filePath string, isDir bool) string {
	if isDir {
		return filePath + "_D_"
	}
	return filePath + "_F_"
}

// parseCacheValue parses cache value to file path and type.
func (sp *SPath) parseCacheValue(value string) (filePath string, isDir bool) {
	if value[len(value)-2 : len(value)-1][0] == 'F' {
		return value[:len(value)-3], false
	}
	return value[:len(value)-3], true
}

// addToCache adds an item to cache.
// If <filePath> is a directory, it also adds its all sub files/directories recursively
// to the cache.
func (sp *SPath) addToCache(filePath, rootPath string) {
	// Add itself firstly.
	idDir := qn_file.IsDir(filePath)
	sp.cache.SetIfNotExist(
		sp.nameFromPath(filePath, rootPath), sp.makeCacheValue(filePath, idDir),
	)
	// If it's a directory, it adds its all sub files/directories recursively.
	if idDir {
		if files, err := qn_file.ScanDir(filePath, "*", true); err == nil {
			//fmt.Println("qn.spath add to cache:", filePath, files)
			for _, path := range files {
				sp.cache.SetIfNotExist(sp.nameFromPath(path, rootPath), sp.makeCacheValue(path, qn_file.IsDir(path)))
			}
		} else {
			//fmt.Errorf(err.Error())
		}
	}
}

// addMonitorByPath adds qn_snotify monitoring recursively.
// When the files under the directory are updated, the cache will be updated meanwhile.
// Note that since the listener is added recursively, if you delete a directory, the files (including the directory)
// under the directory will also generate delete events, which means it will generate N+1 events in total
// if a directory deleted and there're N files under it.
func (sp *SPath) addMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_, _ = qn_snotify.Add(path, func(event *qn_snotify.Event) {
		//qn_log.Debug(event.String())
		switch {
		case event.IsRemove():
			sp.cache.Remove(sp.nameFromPath(event.Path, path))

		case event.IsRename():
			if !qn_file.Exists(event.Path) {
				sp.cache.Remove(sp.nameFromPath(event.Path, path))
			}

		case event.IsCreate():
			sp.addToCache(event.Path, path)
		}
	}, true)
}

// removeMonitorByPath removes qn_snotify monitoring of <path> recursively.
func (sp *SPath) removeMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_ = qn_snotify.Remove(path)
}
