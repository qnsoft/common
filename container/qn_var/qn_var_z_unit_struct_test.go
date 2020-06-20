// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_var_test

import (
	"testing"

	"github.com/qnsoft/common/container/qn_var"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Struct(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type StTest struct {
			Test int
		}

		Kv := make(map[string]int, 1)
		Kv["Test"] = 100

		testObj := &StTest{}

		objOne := qn_var.New(Kv, true)

		objOne.Struct(testObj)

		t.Assert(testObj.Test, Kv["Test"])
	})
	qn_test.C(t, func(t *qn_test.T) {
		type StTest struct {
			Test int8
		}
		o := &StTest{}
		v := qn_var.New(g.Slice{"Test", "-25"})
		v.Struct(o)
		t.Assert(o.Test, -25)
	})
}
