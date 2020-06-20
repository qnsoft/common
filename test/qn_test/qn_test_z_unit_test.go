// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package qn_test_test

import (
	"testing"

	"github.com/gogf/gf/test/qn_test"
)

func TestC(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float64(123.456), float64(123.456))
	})
}

func TestCase(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float64(123.456), float64(123.456))
	})
}
