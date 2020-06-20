// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins_test

import (
	"testing"

	"github.com/qnsoft/common/frame/qn_ins"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_SetGet(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_ins.Set("test-user", 1)
		t.Assert(qn_ins.Get("test-user"), 1)
		t.Assert(qn_ins.Get("none-exists"), nil)
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_ins.GetOrSet("test-1", 1), 1)
		t.Assert(qn_ins.Get("test-1"), 1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_ins.GetOrSetFunc("test-2", func() interface{} {
			return 2
		}), 2)
		t.Assert(qn_ins.Get("test-2"), 2)
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_ins.GetOrSetFuncLock("test-3", func() interface{} {
			return 3
		}), 3)
		t.Assert(qn_ins.Get("test-3"), 3)
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_ins.SetIfNotExist("test-4", 4), true)
		t.Assert(qn_ins.Get("test-4"), 4)
		t.Assert(qn_ins.SetIfNotExist("test-4", 5), false)
		t.Assert(qn_ins.Get("test-4"), 4)
	})
}
