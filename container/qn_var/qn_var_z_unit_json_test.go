// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_var_test

import (
	"math"
	"testing"

	"github.com/qnsoft/common/container/qn_var"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Json(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		s := "i love gf"
		v := qn_var.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	qn_test.C(t, func(t *qn_test.T) {
		s := int64(math.MaxInt64)
		v := qn_var.New(s)
		b1, err1 := json.Marshal(v)
		b2, err2 := json.Marshal(s)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})

	// Unmarshal
	qn_test.C(t, func(t *qn_test.T) {
		s := "i love gf"
		v := qn_var.New(nil)
		b, err := json.Marshal(s)
		t.Assert(err, nil)

		err = json.Unmarshal(b, v)
		t.Assert(err, nil)
		t.Assert(v.String(), s)
	})

	qn_test.C(t, func(t *qn_test.T) {
		var v qn_var.VarImp
		s := "i love gf"
		b, err := json.Marshal(s)
		t.Assert(err, nil)

		err = json.Unmarshal(b, &v)
		t.Assert(err, nil)
		t.Assert(v.String(), s)
	})
}
