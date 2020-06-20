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

func Test_Unsafe(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := "I love 小泽玛利亚"
		t.AssertEQ(qn_conv.UnsafeStrToBytes(s), []byte(s))
	})

	qn_test.C(t, func(t *qn_test.T) {
		b := []byte("I love 小泽玛利亚")
		t.AssertEQ(qn_conv.UnsafeBytesToStr(b), string(b))
	})
}
