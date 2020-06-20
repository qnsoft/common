// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_type_test

import (
	"testing"

	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Bytes(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		i := qn_type.NewBytes([]byte("abc"))
		iClone := i.Clone()
		t.AssertEQ(iClone.Set([]byte("123")), []byte("abc"))
		t.AssertEQ(iClone.Val(), []byte("123"))

		//空参测试
		i1 := qn_type.NewBytes()
		t.AssertEQ(i1.Val(), nil)
	})
}

func Test_Bytes_JSON(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		b := []byte("i love gf")
		i := qn_type.NewBytes(b)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := qn_type.NewBytes()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), b)
	})
}

func Test_Bytes_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *qn_type.Bytes
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
