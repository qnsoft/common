// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package structs_test

import (
	"testing"

	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/internal/structs"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Id   int
			Name string `params:"name"`
			Pass string `my-tag1:"pass1" my-tag2:"pass2" params:"pass"`
		}
		var user User
		t.Assert(structs.Taqn_mapName(user, []string{"params"}, true), qn.Map{"name": "Name", "pass": "Pass"})
		t.Assert(structs.Taqn_mapName(&user, []string{"params"}, true), qn.Map{"name": "Name", "pass": "Pass"})

		t.Assert(structs.Taqn_mapName(&user, []string{"params", "my-tag1"}, true), qn.Map{"name": "Name", "pass": "Pass"})
		t.Assert(structs.Taqn_mapName(&user, []string{"my-tag1", "params"}, true), qn.Map{"name": "Name", "pass1": "Pass"})
		t.Assert(structs.Taqn_mapName(&user, []string{"my-tag2", "params"}, true), qn.Map{"name": "Name", "pass2": "Pass"})
	})

	qn_test.C(t, func(t *qn_test.T) {
		type Base struct {
			Pass1 string `params:"password1"`
			Pass2 string `params:"password2"`
		}
		type UserWithBase struct {
			Id   int
			Name string
			Base `params:"base"`
		}
		user := new(UserWithBase)
		t.Assert(structs.Taqn_mapName(user, []string{"params"}, true), qn.Map{
			"base":      "Base",
			"password1": "Pass1",
			"password2": "Pass2",
		})
	})

	qn_test.C(t, func(t *qn_test.T) {
		type Base struct {
			Pass1 string `params:"password1"`
			Pass2 string `params:"password2"`
		}
		type UserWithBase1 struct {
			Id   int
			Name string
			Base
		}
		type UserWithBase2 struct {
			Id   int
			Name string
			Pass Base
		}
		user1 := new(UserWithBase1)
		user2 := new(UserWithBase2)
		t.Assert(structs.Taqn_mapName(user1, []string{"params"}, true), qn.Map{"password1": "Pass1", "password2": "Pass2"})
		t.Assert(structs.Taqn_mapName(user2, []string{"params"}, true), qn.Map{"password1": "Pass1", "password2": "Pass2"})
	})
}
