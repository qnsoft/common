// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_hash_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_hash"
)

var (
	str = []byte("This is the test string for hash.")
)

func BenchmarkBKDRHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.BKDRHash(str)
	}
}

func BenchmarkBKDRHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.BKDRHash64(str)
	}
}

func BenchmarkSDBMHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.SDBMHash(str)
	}
}

func BenchmarkSDBMHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.SDBMHash64(str)
	}
}

func BenchmarkRSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.RSHash(str)
	}
}

func BenchmarkSRSHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.RSHash64(str)
	}
}

func BenchmarkJSHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.JSHash(str)
	}
}

func BenchmarkJSHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.JSHash64(str)
	}
}

func BenchmarkPJWHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.PJWHash(str)
	}
}

func BenchmarkPJWHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.PJWHash64(str)
	}
}

func BenchmarkELFHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.ELFHash(str)
	}
}

func BenchmarkELFHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.ELFHash64(str)
	}
}

func BenchmarkDJBHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.DJBHash(str)
	}
}

func BenchmarkDJBHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.DJBHash64(str)
	}
}

func BenchmarkAPHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.APHash(str)
	}
}

func BenchmarkAPHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_hash.APHash64(str)
	}
}
