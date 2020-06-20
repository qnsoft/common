// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_util_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_util"
)

func Test_Dump(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_util.Dump(map[int]int{
			100: 100,
		})
	})
}

func Test_TryCatch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_util.TryCatch(func() {
			panic("qn_util TryCatch test")
		})
	})

	qn_test.C(t, func(t *qn_test.T) {
		qn_util.TryCatch(func() {
			panic("qn_util TryCatch test")

		}, func(err interface{}) {
			t.Assert(err, "qn_util TryCatch test")
		})
	})
}

func Test_IsEmpty(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.IsEmpty(1), false)
	})
}

func Test_Throw(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		defer func() {
			t.Assert(recover(), "qn_util Throw test")
		}()

		qn_util.Throw("qn_util Throw test")
	})
}
