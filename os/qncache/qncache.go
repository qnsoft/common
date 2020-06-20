// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// Package gfcache provides reading and caching for file contents.
package qncache

import (
	"time"

	"github.com/qnsoft/common/internal/cmdenv"
	"github.com/qnsoft/common/os/gcache"
	"github.com/qnsoft/common/os/gfsnotify"
	"github.com/qnsoft/common/os/qn_file"
)

const (
	// Default expire time for file content caching in seconds.
	gDEFAULT_CACHE_EXPIRE = time.Minute
)

var (
	// Default expire time for file content caching.
	cacheExpire = cmdenv.Get("qn.qncache.expire", gDEFAULT_CACHE_EXPIRE).Duration()
)

// GetContents returns string content of given file by <path> from cache.
// If there's no content in the cache, it will read it from disk file specified by <path>.
// The parameter <expire> specifies the caching time for this file content in seconds.
func GetContents(path string, duration ...time.Duration) string {
	return string(GetBinContents(path, duration...))
}

// GetBinContents returns []byte content of given file by <path> from cache.
// If there's no content in the cache, it will read it from disk file specified by <path>.
// The parameter <expire> specifies the caching time for this file content in seconds.
func GetBinContents(path string, duration ...time.Duration) []byte {
	key := cacheKey(path)
	expire := cacheExpire
	if len(duration) > 0 {
		expire = duration[0]
	}
	r := gcache.GetOrSetFuncLock(key, func() interface{} {
		b := qn_file.GetBytes(path)
		if b != nil {
			// Adding this <path> to gfsnotify,
			// it will clear its cache if there's any changes of the file.
			_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
				gcache.Remove(key)
				gfsnotify.Exit()
			})
		}
		return b
	}, expire)
	if r != nil {
		return r.([]byte)
	}
	return nil
}

// cacheKey produces the cache key for gcache.
func cacheKey(path string) string {
	return "qn.qncache:" + path
}
