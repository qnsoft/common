// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Params_Xml_Request(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Time  *time.Time
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2" v:"required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致"`
	}
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/get", func(r *qn_http.Request) {
		r.Response.WriteExit(r.Get("id"), r.Get("name"))
	})
	s.BindHandler("/map", func(r *qn_http.Request) {
		if m := r.GetMap(); len(m) > 0 {
			r.Response.WriteExit(m["id"], m["name"], m["password1"], m["password2"])
		}
	})
	s.BindHandler("/parse", func(r *qn_http.Request) {
		if m := r.GetMap(); len(m) > 0 {
			var user *User
			if err := r.Parse(&user); err != nil {
				r.Response.WriteExit(err)
			}
			r.Response.WriteExit(user.Id, user.Name, user.Pass1, user.Pass2)
		}
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		content1 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123Abc!@#</password2></doc>`
		content2 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123</password2></doc>`
		t.Assert(client.GetContent("/get", content1), `1john`)
		t.Assert(client.PostContent("/get", content1), `1john`)
		t.Assert(client.GetContent("/map", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.PostContent("/map", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.PostContent("/parse", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.PostContent("/parse", content2), `密码强度不足; 两次密码不一致`)
	})
}
