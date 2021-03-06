// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
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

func Test_Scan(t *testing.T) {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `qn_conv:"password1"`
		Pass2 string `qn_conv:"password2"`
	}
	qn_test.C(t, func(t *qn_test.T) {
		var (
			user   = new(User)
			params = qn.Map{
				"uid":   1,
				"name":  "john",
				"PASS1": "123",
				"PASS2": "456",
			}
		)
		err := qn_conv.Scan(params, user)
		t.Assert(err, nil)
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var (
			users  []qn.Sl
			params = g.Slice{
				qn.Map{
					"uid":   1,
					"name":  "john1",
					"PASS1": "111",
					"PASS2": "222",
				},
				qn.Map{
					"uid":   2,
					"name":  "john2",
					"PASS1": "333",
					"PASS2": "444",
				},
			}
		)
		err := qn_conv.Scan(params, &users)
		t.Assert(err, niqn.Sl
		t.Assert(users, g.Slice{
			&User{
				Uid:   1,
				Name:  "john1",
				Pass1: "111",
				Pass2: "222",
			},
			&User{
				Uid:   2,
				Name:  "john2",
				Pass1: "333",
				Pass2: "444",
			},
		})
	})
}

func Test_ScanStr(t *testing.T) {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `qn_conv:"password1"`
		Pass2 string `qn_conv:"password2"`
	}
	qn_test.C(t, func(t *qn_test.T) {
		var (
			user   = new(User)
			params = `{"uid":1,"name":"john", "pass1":"123","pass2":"456"}`
		)
		err := qn_conv.Scan(params, user)
		t.Assert(err, nil)
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var (
			users  []User
			params = `[
{"uid":1,"name":"john1", "pass1":"111","pass2":"222"},
{"uid":2,"name":"john2", "pass1":"333","pass2":"444"}
]`
		)
		err := qn_conv.Scan(params, &users)
		t.Assert(err, niqn.Sl
		t.Assert(users, g.Slice{
			&User{
				Uid:   1,
				Name:  "john1",
				Pass1: "111",
				Pass2: "222",
			},
			&User{
				Uid:   2,
				Name:  "john2",
				Pass1: "333",
				Pass2: "444",
			},
		})
	})
}
