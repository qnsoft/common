// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cache

import (
	"time"

	"github.com/qnsoft/common/container/qn_list"
	"github.com/qnsoft/common/container/qn_map"
	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/os/qn_timer"
)

// LRU cache object.
// It uses list.List from stdlib for its underlying doubly linked list.
type memCacheLru struct {
	cache   *memCache   // Parent cache object.
	data    *qn_map.Map   // Key mapping to the item of the list.
	list    *qn_list.List // Key list.
	rawList *qn_list.List // History for key adding.
	closed  *qn_type.Bool // Closed or not.
}

// newMemCacheLru creates and returns a new LRU object.
func newMemCacheLru(cache *memCache) *memCacheLru {
	lru := &memCacheLru{
		cache:   cache,
		data:    qn_map.New(true),
		list:    qn_list.New(true),
		rawList: qn_list.New(true),
		closed:  qn_type.NewBool(),
	}
	qn_timer.AddSingleton(time.Second, lru.SyncAndClear)
	return lru
}

// Close closes the LRU object.
func (lru *memCacheLru) Close() {
	lru.closed.Set(true)
}

// Remove deletes the <key> FROM <lru>.
func (lru *memCacheLru) Remove(key interface{}) {
	if v := lru.data.Get(key); v != nil {
		lru.data.Remove(key)
		lru.list.Remove(v.(*qn_list.Element))
	}
}

// Size returns the size of <lru>.
func (lru *memCacheLru) Size() int {
	return lru.data.Size()
}

// Push pushes <key> to the tail of <lru>.
func (lru *memCacheLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop deletes and returns the key from tail of <lru>.
func (lru *memCacheLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.Remove(v)
		return v
	}
	return nil
}

// Print is used for test only.
//func (lru *memCacheLru) Print() {
//    for _, v := range lru.list.FrontAll() {
//        fmt.Printf("%v ", v)
//    }
//    fmt.Println()
//}

// SyncAndClear synchronizes the keys from <rawList> to <list> and <data>
// using Least Recently Used algorithm.
func (lru *memCacheLru) SyncAndClear() {
	if lru.closed.Val() {
		qn_timer.Exit()
		return
	}
	// Data synchronization.
	for {
		if v := lru.rawList.PopFront(); v != nil {
			// Deleting the key from list.
			if v := lru.data.Get(v); v != nil {
				lru.list.Remove(v.(*qn_list.Element))
			}
			// Pushing key to the head of the list
			// and setting its list item to hash table for quick indexing.
			lru.data.Set(v, lru.list.PushFront(v))
		} else {
			break
		}
	}
	// Data cleaning up.
	for i := lru.Size() - lru.cache.cap; i > 0; i-- {
		if s := lru.Pop(); s != nil {
			lru.cache.clearByKey(s, true)
		}
	}
}
