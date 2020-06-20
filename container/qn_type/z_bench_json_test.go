// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".+\_Json" -benchmem

package qn_type_test

import (
	"testing"

	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/internal/json"
)

var (
	vBool      = qn_type.NewBool()
	vByte      = qn_type.NewByte()
	vBytes     = qn_type.NewBytes()
	vFloat32   = qn_type.NewFloat32()
	vFloat64   = qn_type.NewFloat64()
	vInt       = qn_type.NewInt()
	vInt32     = qn_type.NewInt32()
	vInt64     = qn_type.NewInt64()
	vInterface = qn_type.NewInterface()
	vString    = qn_type.NewString()
	vUint      = qn_type.NewUint()
	vUint32    = qn_type.NewUint32()
	vUint64    = qn_type.NewUint64()
)

func Benchmark_Bool_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBool)
	}
}

func Benchmark_Byte_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vByte)
	}
}

func Benchmark_Bytes_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vBytes)
	}
}

func Benchmark_Float32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat32)
	}
}

func Benchmark_Float64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vFloat64)
	}
}

func Benchmark_Int_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt)
	}
}

func Benchmark_Int32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt32)
	}
}

func Benchmark_Int64_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInt64)
	}
}

func Benchmark_Interface_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vInterface)
	}
}

func Benchmark_String_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vString)
	}
}

func Benchmark_Uint_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint)
	}
}

func Benchmark_Uint32_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(vUint64)
	}
}
