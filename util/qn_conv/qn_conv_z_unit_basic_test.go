// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		f32 := float32(123.456)
		i64 := int64(1552578474888)
		t.AssertEQ(qn_conv.Int(f32), int(123))
		t.AssertEQ(qn_conv.Int8(f32), int8(123))
		t.AssertEQ(qn_conv.Int16(f32), int16(123))
		t.AssertEQ(qn_conv.Int32(f32), int32(123))
		t.AssertEQ(qn_conv.Int64(f32), int64(123))
		t.AssertEQ(qn_conv.Int64(f32), int64(123))
		t.AssertEQ(qn_conv.Uint(f32), uint(123))
		t.AssertEQ(qn_conv.Uint8(f32), uint8(123))
		t.AssertEQ(qn_conv.Uint16(f32), uint16(123))
		t.AssertEQ(qn_conv.Uint32(f32), uint32(123))
		t.AssertEQ(qn_conv.Uint64(f32), uint64(123))
		t.AssertEQ(qn_conv.Float32(f32), float32(123.456))
		t.AssertEQ(qn_conv.Float64(i64), float64(i64))
		t.AssertEQ(qn_conv.Bool(f32), true)
		t.AssertEQ(qn_conv.String(f32), "123.456")
		t.AssertEQ(qn_conv.String(i64), "1552578474888")
	})

	qn_test.C(t, func(t *qn_test.T) {
		s := "-0xFF"
		t.Assert(qn_conv.Int(s), int64(-0xFF))
	})
}
