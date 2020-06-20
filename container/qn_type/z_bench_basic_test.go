// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qn_type_test

import (
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/qnsoft/common/container/qn_type"

	"github.com/qnsoft/common/encoding/qn_binary"
)

var (
	it     = qn_type.NewInt()
	it32   = qn_type.NewInt32()
	it64   = qn_type.NewInt64()
	uit    = qn_type.NewUint()
	uit32  = qn_type.NewUint32()
	uit64  = qn_type.NewUint64()
	bl     = qn_type.NewBool()
	vbytes = qn_type.NewBytes()
	str    = qn_type.NewString()
	inf    = qn_type.NewInterface()
	at     = atomic.Value{}
)

func BenchmarkInt_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Set(i)
	}
}

func BenchmarkInt_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Val()
	}
}

func BenchmarkInt_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Add(i)
	}
}

func BenchmarkInt_Cas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		it.Cas(i, i)
	}
}

func BenchmarkInt32_Set(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Set(i)
	}
}

func BenchmarkInt32_Val(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Val()
	}
}

func BenchmarkInt32_Add(b *testing.B) {
	for i := int32(0); i < int32(b.N); i++ {
		it32.Add(i)
	}
}

func BenchmarkInt64_Set(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Set(i)
	}
}

func BenchmarkInt64_Val(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Val()
	}
}

func BenchmarkInt64_Add(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		it64.Add(i)
	}
}

func BenchmarkUint_Set(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Set(i)
	}
}

func BenchmarkUint_Val(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Val()
	}
}

func BenchmarkUint_Add(b *testing.B) {
	for i := uint(0); i < uint(b.N); i++ {
		uit.Add(i)
	}
}

func BenchmarkUint32_Set(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Set(i)
	}
}

func BenchmarkUint32_Val(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Val()
	}
}

func BenchmarkUint32_Add(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		uit32.Add(i)
	}
}

func BenchmarkUint64_Set(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Set(i)
	}
}

func BenchmarkUint64_Val(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Val()
	}
}

func BenchmarkUint64_Add(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		uit64.Add(i)
	}
}

func BenchmarkBool_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Set(true)
	}
}

func BenchmarkBool_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Val()
	}
}

func BenchmarkBool_Cas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bl.Cas(false, true)
	}
}

func BenchmarkString_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.Set(strconv.Itoa(i))
	}
}

func BenchmarkString_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str.Val()
	}
}

func BenchmarkBytes_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.Set(qn_binary.EncodeInt(i))
	}
}

func BenchmarkBytes_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vbytes.Val()
	}
}

func BenchmarkInterface_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.Set(i)
	}
}

func BenchmarkInterface_Val(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inf.Val()
	}
}

func BenchmarkAtomicValue_Store(b *testing.B) {
	for i := 0; i < b.N; i++ {
		at.Store(i)
	}
}

func BenchmarkAtomicValue_Load(b *testing.B) {
	for i := 0; i < b.N; i++ {
		at.Load()
	}
}
