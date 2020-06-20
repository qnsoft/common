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

func Test_Router_Group_Group(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/api.v2", func(group *qn_http.RouterGroup) {
		group.Middleware(func(r *qn_http.Request) {
			r.Response.Write("1")
			r.Middleware.Next()
			r.Response.Write("2")
		})
		group.GET("/test", func(r *qn_http.Request) {
			r.Response.Write("test")
		})
		group.Group("/order", func(group *qn_http.RouterGroup) {
			group.GET("/list", func(r *qn_http.Request) {
				r.Response.Write("list")
			})
			group.PUT("/update", func(r *qn_http.Request) {
				r.Response.Write("update")
			})
		})
		group.Group("/user", func(group *qn_http.RouterGroup) {
			group.GET("/info", func(r *qn_http.Request) {
				r.Response.Write("info")
			})
			group.POST("/edit", func(r *qn_http.Request) {
				r.Response.Write("edit")
			})
			group.DELETE("/drop", func(r *qn_http.Request) {
				r.Response.Write("drop")
			})
		})
		group.Group("/hook", func(group *qn_http.RouterGroup) {
			group.Hook("/*", qn_http.HOOK_BEFORE_SERVE, func(r *qn_http.Request) {
				r.Response.Write("hook any")
			})
			group.Hook("/:name", qn_http.HOOK_BEFORE_SERVE, func(r *qn_http.Request) {
				r.Response.Write("hook name")
			})
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/api.v2"), "Not Found")
		t.Assert(client.GetContent("/api.v2/test"), "1test2")
		t.Assert(client.GetContent("/api.v2/hook"), "hook any")
		t.Assert(client.GetContent("/api.v2/hook/name"), "hook namehook any")
		t.Assert(client.GetContent("/api.v2/hook/name/any"), "hook any")
		t.Assert(client.GetContent("/api.v2/order/list"), "1list2")
		t.Assert(client.GetContent("/api.v2/order/update"), "Not Found")
		t.Assert(client.PutContent("/api.v2/order/update"), "1update2")
		t.Assert(client.GetContent("/api.v2/user/drop"), "Not Found")
		t.Assert(client.DeleteContent("/api.v2/user/drop"), "1drop2")
		t.Assert(client.GetContent("/api.v2/user/edit"), "Not Found")
		t.Assert(client.PostContent("/api.v2/user/edit"), "1edit2")
		t.Assert(client.GetContent("/api.v2/user/info"), "1info2")
	})
}
