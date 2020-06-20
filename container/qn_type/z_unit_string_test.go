// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_type_test

import (
	"testing"

	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_String(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		i := qn_type.NewString("abc")
		iClone := i.Clone()
		t.AssertEQ(iClone.Set("123"), "abc")
		t.AssertEQ(iClone.Val(), "123")

		//空参测试
		i1 := qn_type.NewString()
		t.AssertEQ(i1.Val(), "")
	})
}

func Test_String_JSON(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := "i love gf"
		i1 := qn_type.NewString(s)
		b1, err1 := json.Marshal(i1)
		b2, err2 := json.Marshal(i1.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := qn_type.NewString()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), s)
	})
}

func Test_String_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *qn_type.String
	}
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
