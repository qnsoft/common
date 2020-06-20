// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_var_test

import (
	"testing"

	"github.com/qnsoft/common/container/qn_var"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Map(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn.Map{
			"k1": "v1",
			"k2": "v2",
		}
		objOne := qn_var.New(m, true)
		t.Assert(objOne.Map()["k1"], m["k1"])
		t.Assert(objOne.Map()["k2"], m["k2"])
	})
}
