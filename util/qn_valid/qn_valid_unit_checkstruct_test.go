// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_valid_test

import (
	"testing"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_valid"
)

func Test_CheckStruct(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"@required|length:6,16",
			"@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := qn_valid.CheckStruct(obj, rules, msgs)
		t.Assert(err, nil)
	})

	qn_test.C(t, func(t *qn_test.T) {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"Name@required|length:6,16#名称不能为空",
			"Age@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := qn_valid.CheckStruct(obj, rules, msgs)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 2)
		t.Assert(err.Maps()["Name"]["required"], "")
		t.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		t.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"Name@required|length:6,16#名称不能为空|",
			"Age@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := qn_valid.CheckStruct(obj, rules, msgs)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 2)
		t.Assert(err.Maps()["Name"]["required"], "")
		t.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		t.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type Object struct {
			Name string
			Age  int
		}
		rules := map[string]string{
			"Name": "required|length:6,16",
			"Age":  "between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := qn_valid.CheckStruct(obj, rules, msgs)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 2)
		t.Assert(err.Maps()["Name"]["required"], "")
		t.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		t.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type LoginRequest struct {
			Username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := qn_valid.CheckStruct(login, nil)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 2)
		t.Assert(err.Maps()["username"]["required"], "用户名不能为空")
		t.Assert(err.Maps()["password"]["required"], "登录密码不能为空")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type LoginRequest struct {
			Username string `json:"username" qn_valid:"@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := qn_valid.CheckStruct(login, nil)
		t.Assert(err, nil)
	})

	qn_test.C(t, func(t *qn_test.T) {
		type LoginRequest struct {
			username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := qn_valid.CheckStruct(login, nil)
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["password"]["required"], "登录密码不能为空")
	})

	// qn_valid tag
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Id       int    `qn_valid:"uid@required|min:10#|ID不能为空"`
			Age      int    `qn_valid:"age@required#年龄不能为空"`
			Username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := qn_valid.CheckStruct(user, nil)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 1)
		t.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Id       int    `qn_valid:"uid@required|min:10#|ID不能为空"`
			Age      int    `qn_valid:"age@required#年龄不能为空"`
			Username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}

		rules := []string{
			"username@required#用户名不能为空",
		}

		err := qn_valid.CheckStruct(user, rules)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 1)
		t.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})

	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Id       int    `qn_valid:"uid@required|min:10#ID不能为空"`
			Age      int    `qn_valid:"age@required#年龄不能为空"`
			Username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := qn_valid.CheckStruct(user, nil)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 1)
	})

	// valid tag
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Id       int    `valid:"uid@required|min:10#|ID不能为空"`
			Age      int    `valid:"age@required#年龄不能为空"`
			Username string `json:"username" qn_valid:"username@required#用户名不能为空"`
			Password string `json:"password" qn_valid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := qn_valid.CheckStruct(user, nil)
		t.AssertNE(err, nil)
		t.Assert(len(err.Maps()), 1)
		t.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})
}

func Test_CheckStruct_With_Inherit(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type Pass struct {
			Pass1 string `valid:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
			Pass2 string `valid:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
		}
		type User struct {
			Id   int
			Name string `valid:"name@required#请输入您的姓名"`
			Pass Pass
		}
		user := &User{
			Name: "",
			Pass: Pass{
				Pass1: "1",
				Pass2: "2",
			},
		}
		err := qn_valid.CheckStruct(user, nil)
		t.AssertNE(err, nil)
		t.Assert(err.Maps()["name"], g.Map{"required": "请输入您的姓名"})
		t.Assert(err.Maps()["password1"], g.Map{"same": "您两次输入的密码不一致"})
		t.Assert(err.Maps()["password2"], g.Map{"same": "您两次输入的密码不一致"})
	})
}
