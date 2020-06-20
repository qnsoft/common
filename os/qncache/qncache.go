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
	"github.com/qnsoft/common/os/qn_cache"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_snotify"
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
	r := qn_cache.GetOrSetFuncLock(key, func() interface{} {
		b := qn_file.GetBytes(path)
		if b != nil {
			// Adding this <path> to qn_snotify,
			// it will clear its cache if there's any changes of the file.
			_, _ = qn_snotify.Add(path, func(event *qn_snotify.Event) {
				qn_cache.Remove(key)
				qn_snotify.Exit()
			})
		}
		return b
	}, expire)
	if r != nil {
		return r.([]byte)
	}
	return nil
}

// cacheKey produces the cache key for qn_cache.
func cacheKey(path string) string {
	return "qn.qncache:" + path
}
