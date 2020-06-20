// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package empty_test

import (
	"testing"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/internal/empty"
	"github.com/qnsoft/common/test/qn_test"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

type TestPerson interface {
	Say() string
}
type TestWoman struct {
}

func (woman TestWoman) Say() string {
	return "nice"
}

func TestIsEmpty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		tmpT1 := "0"
		tmpT2 := func() {}
		tmpT2 = nil
		tmpT3 := make(chan int, 0)
		var tmpT4 TestPerson = nil
		var tmpT5 *TestPerson = nil
		tmpF1 := "1"
		tmpF2 := func(a string) string { return "1" }
		tmpF3 := make(chan int, 1)
		tmpF3 <- 1
		var tmpF4 TestPerson = TestWoman{}
		tmpF5 := &tmpF4
		// true
		t.Assert(empty.IsEmpty(nil), true)
		t.Assert(empty.IsEmpty(qn_conv.Int(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Int8(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Int16(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Int32(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Int64(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Uint64(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Uint(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Uint16(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Uint32(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Uint64(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Float32(tmpT1)), true)
		t.Assert(empty.IsEmpty(qn_conv.Float64(tmpT1)), true)
		t.Assert(empty.IsEmpty(false), true)
		t.Assert(empty.IsEmpty([]byte("")), true)
		t.Assert(empty.IsEmpty(""), true)
		t.Assert(empty.IsEmpty(g.Map{}), true)
		t.Assert(empty.IsEmpty(g.Slice{}), true)
		t.Assert(empty.IsEmpty(g.Array{}), true)
		t.Assert(empty.IsEmpty(tmpT2), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT3), true)
		t.Assert(empty.IsEmpty(tmpT4), true)
		t.Assert(empty.IsEmpty(tmpT5), true)
		// false
		t.Assert(empty.IsEmpty(qn_conv.Int(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Int8(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Int16(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Int32(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Int64(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Uint(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Uint8(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Uint16(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Uint32(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Uint64(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Float32(tmpF1)), false)
		t.Assert(empty.IsEmpty(qn_conv.Float64(tmpF1)), false)
		t.Assert(empty.IsEmpty(true), false)
		t.Assert(empty.IsEmpty(tmpT1), false)
		t.Assert(empty.IsEmpty([]byte("1")), false)
		t.Assert(empty.IsEmpty(g.Map{"a": 1}), false)
		t.Assert(empty.IsEmpty(g.Slice{"1"}), false)
		t.Assert(empty.IsEmpty(g.Array{"1"}), false)
		t.Assert(empty.IsEmpty(tmpF2), false)
		t.Assert(empty.IsEmpty(tmpF3), false)
		t.Assert(empty.IsEmpty(tmpF4), false)
		t.Assert(empty.IsEmpty(tmpF5), false)
	})
}
