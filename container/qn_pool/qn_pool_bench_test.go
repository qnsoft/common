// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_pool_test

import (
	"sync"
	"testing"
	"time"

	"github.com/qnsoft/common/container/gpool"
)

var pool = gpool.New(time.Hour, nil)
var syncp = sync.Pool{}

func BenchmarkGPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
}

func BenchmarkGPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Put(i)
	}
}

func BenchmarkSyncPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Get()
	}
}
