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

type boolStruct struct {
}

func Test_Bool(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var i interface{} = nil
		t.AssertEQ(qn_conv.Bool(i), false)
		t.AssertEQ(qn_conv.Bool(false), false)
		t.AssertEQ(qn_conv.Bool(nil), false)
		t.AssertEQ(qn_conv.Bool(0), false)
		t.AssertEQ(qn_conv.Bool("0"), false)
		t.AssertEQ(qn_conv.Bool(""), false)
		t.AssertEQ(qn_conv.Bool("false"), false)
		t.AssertEQ(qn_conv.Bool("off"), false)
		t.AssertEQ(qn_conv.Bool([]byte{}), false)
		t.AssertEQ(qn_conv.Bool([]string{}), false)
		t.AssertEQ(qn_conv.Bool([]interface{}{}), false)
		t.AssertEQ(qn_conv.Bool([]map[int]int{}), false)

		t.AssertEQ(qn_conv.Bool("1"), true)
		t.AssertEQ(qn_conv.Bool("on"), true)
		t.AssertEQ(qn_conv.Bool(1), true)
		t.AssertEQ(qn_conv.Bool(123.456), true)
		t.AssertEQ(qn_conv.Bool(boolStruct{}), true)
		t.AssertEQ(qn_conv.Bool(&boolStruct{}), true)
	})
}
