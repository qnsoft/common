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

	"github.com/qnsoft/common/container/gtype"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Bool(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		i := gtype.NewBool(true)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(false), true)
		t.AssertEQ(iClone.Val(), false)

		i1 := gtype.NewBool(false)
		iClone1 := i1.Clone()
		t.AssertEQ(iClone1.Set(true), false)
		t.AssertEQ(iClone1.Val(), true)

		//空参测试
		i2 := gtype.NewBool()
		t.AssertEQ(i2.Val(), false)
	})
}

func Test_Bool_JSON(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		i := gtype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	qn_test.C(t, func(t *qn_test.T) {
		i := gtype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)
	})
	// Unmarshal
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		i := gtype.NewBool()
		err = json.Unmarshal([]byte("true"), &i)
		t.Assert(err, nil)
		t.Assert(i.Val(), true)
		err = json.Unmarshal([]byte("false"), &i)
		t.Assert(err, nil)
		t.Assert(i.Val(), false)
		err = json.Unmarshal([]byte("1"), &i)
		t.Assert(err, nil)
		t.Assert(i.Val(), true)
		err = json.Unmarshal([]byte("0"), &i)
		t.Assert(err, nil)
		t.Assert(i.Val(), false)
	})

	qn_test.C(t, func(t *qn_test.T) {
		i := gtype.NewBool(true)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBool()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), i.Val())
	})
	qn_test.C(t, func(t *qn_test.T) {
		i := gtype.NewBool(false)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := gtype.NewBool()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), i.Val())
	})
}

func Test_Bool_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *gtype.Bool
	}
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "true",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), true)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "false",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), false)
	})
}
