// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
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

func Test_Struct_Slice(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []int
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []int32
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []int64
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []uint
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []uint32
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []uint64
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []float32
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Scores []float64
		}
		user := nqn.Slser)
		array := g.Slice{1, 2, 3}
		err := qn_conv.Struct(qn.Map{"scores": array}, user)
		t.Assert(err, nil)
		t.Assert(user.Scores, array)
	})
}

func Test_Struct_SliceWithTag(t *testing.T) {
	type User struct {
		Uid      int    `json:"id"`
		NickName string `json:"name"`
	}
	qn_test.C(t, func(t *qn_test.T) {
		var users qn.Sler
		params := g.Slice{
			qn.Map{
				"id":   1,
				"name": "name1",
			},
			qn.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := qn_conv.Structs(params, &users)
		t.Assert(err, nil)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
	qn_test.C(t, func(t *qn_test.T) {
		var users qn.Slser
		params := g.Slice{
			qn.Map{
				"id":   1,
				"name": "name1",
			},
			qn.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := qn_conv.Structs(params, &users)
		t.Assert(err, nil)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
}
