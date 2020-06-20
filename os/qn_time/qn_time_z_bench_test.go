// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_time_test

import (
	"testing"

	"github.com/qnsoft/common/os/qn_time"
)

func Benchmark_Timestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.Timestamp()
	}
}

func Benchmark_TimestampMilli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.TimestampMilli()
	}
}

func Benchmark_TimestampMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.TimestampMicro()
	}
}

func Benchmark_TimestampNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.TimestampNano()
	}
}

func Benchmark_StrToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.StrToTime("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_ParseTimeFromContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.ParseTimeFromContent("2018-02-09T20:46:17.897Z")
	}
}

func Benchmark_NewFromTimeStamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.NewFromTimeStamp(1542674930)
	}
}

func Benchmark_Date(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.Date()
	}
}

func Benchmark_Datetime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.Datetime()
	}
}

func Benchmark_SetTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qn_time.SetTimeZone("Asia/Shanghai")
	}
}
