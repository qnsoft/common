// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv_test

import (
	"testing"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_Slice(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := 123.456
		t.AssertEQ(qn_conv.Bytes("123"), []byte("123"))
		t.AssertEQ(qn_conv.Strings(value), []string{"123.456"})
		t.AssertEQ(qn_conv.Ints(value), []int{123})
		t.AssertEQ(qn_conv.Floats(value), []float64{123.456})
		t.AssertEQ(qn_conv.Interfaces(value), []interface{}{123.456})
	})
}

func Test_Slice_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	qn_test.C(t, func(t *qn_test.T) {
		user := &User{1, "john"}qn.Sl
		t.Assert(qn_conv.Interfaces(user), g.Slice{1})
	})
}

func Test_Slice_Structs(t *testing.T) {
	type Base struct {
		Age int
	}
	type User struct {
		Id   int
		Name string
		Base
	}

	qn_test.C(t, func(t *qn_test.T) {
		users := make([]User, 0)
		params := []qn.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := qn_conv.Structs(params, &users)
		t.Assert(err, nil)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, 0)

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, 0)
	})

	qn_test.C(t, func(t *qn_test.T) {
		users := make([]User, 0)
		params := []qn.Map{
			{"id": 1, "name": "john", "age": 18},
			{"id": 2, "name": "smith", "age": 20},
		}
		err := qn_conv.StructsDeep(params, &users)
		t.Assert(err, nil)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, params[0]["id"])
		t.Assert(users[0].Name, params[0]["name"])
		t.Assert(users[0].Age, params[0]["age"])

		t.Assert(users[1].Id, params[1]["id"])
		t.Assert(users[1].Name, params[1]["name"])
		t.Assert(users[1].Age, params[1]["age"])
	})
}
