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

type stringStruct1 struct {
	Name string
}

type stringStruct2 struct {
	Name string
}

func (s *stringStruct1) String() string {
	return s.Name
}

func Test_String(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.AssertEQ(qn_conv.String(int(123)), "123")
		t.AssertEQ(qn_conv.String(int(-123)), "-123")
		t.AssertEQ(qn_conv.String(int8(123)), "123")
		t.AssertEQ(qn_conv.String(int8(-123)), "-123")
		t.AssertEQ(qn_conv.String(int16(123)), "123")
		t.AssertEQ(qn_conv.String(int16(-123)), "-123")
		t.AssertEQ(qn_conv.String(int32(123)), "123")
		t.AssertEQ(qn_conv.String(int32(-123)), "-123")
		t.AssertEQ(qn_conv.String(int64(123)), "123")
		t.AssertEQ(qn_conv.String(int64(-123)), "-123")
		t.AssertEQ(qn_conv.String(int64(1552578474888)), "1552578474888")
		t.AssertEQ(qn_conv.String(int64(-1552578474888)), "-1552578474888")

		t.AssertEQ(qn_conv.String(uint(123)), "123")
		t.AssertEQ(qn_conv.String(uint8(123)), "123")
		t.AssertEQ(qn_conv.String(uint16(123)), "123")
		t.AssertEQ(qn_conv.String(uint32(123)), "123")
		t.AssertEQ(qn_conv.String(uint64(155257847488898765)), "155257847488898765")

		t.AssertEQ(qn_conv.String(float32(123.456)), "123.456")
		t.AssertEQ(qn_conv.String(float32(-123.456)), "-123.456")
		t.AssertEQ(qn_conv.String(float64(1552578474888.456)), "1552578474888.456")
		t.AssertEQ(qn_conv.String(float64(-1552578474888.456)), "-1552578474888.456")

		t.AssertEQ(qn_conv.String(true), "true")
		t.AssertEQ(qn_conv.String(false), "false")

		t.AssertEQ(qn_conv.String([]byte("bytes")), "bytes")

		t.AssertEQ(qn_conv.String(stringStruct1{"john"}), `{"Name":"john"}`)
		t.AssertEQ(qn_conv.String(&stringStruct1{"john"}), "john")

		t.AssertEQ(qn_conv.String(stringStruct2{"john"}), `{"Name":"john"}`)
		t.AssertEQ(qn_conv.String(&stringStruct2{"john"}), `{"Name":"john"}`)
	})
}
