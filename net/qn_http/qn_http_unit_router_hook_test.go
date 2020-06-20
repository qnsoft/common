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
	"github.com/qnsoft/common/net/ghttp"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Router_Hook_Basic(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE:  func(r *ghttp.Request) { r.Response.Write("1") },
		ghttp.HOOK_AFTER_SERVE:   func(r *ghttp.Request) { r.Response.Write("2") },
		ghttp.HOOK_BEFORE_OUTPUT: func(r *ghttp.Request) { r.Response.Write("3") },
		ghttp.HOOK_AFTER_OUTPUT:  func(r *ghttp.Request) { r.Response.Write("4") },
	})
	s.BindHandler("/test/test", func(r *ghttp.Request) {
		r.Response.Write("test")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "123")
		t.Assert(client.GetContent("/test/test"), "1test23")
	})
}

func Test_Router_Hook_Fuzzy_Router(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	i := 1000
	pattern1 := "/:name/info"
	s.BindHookHandlerByMap(pattern1, map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.SetParam("uid", i)
			i++
		},
	})
	s.BindHandler(pattern1, func(r *ghttp.Request) {
		r.Response.Write(r.Get("uid"))
	})

	pattern2 := "/{object}/list/{page}.java"
	s.BindHookHandlerByMap(pattern2, map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_OUTPUT: func(r *ghttp.Request) {
			r.Response.SetBuffer([]byte(
				fmt.Sprint(r.Get("object"), "&", r.Get("page"), "&", i),
			))
		},
	})
	s.BindHandler(pattern2, func(r *ghttp.Request) {
		r.Response.Write(r.Router.Uri)
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/john"), "Not Found")
		t.Assert(client.GetContent("/john/info"), "1000")
		t.Assert(client.GetContent("/john/info"), "1001")
		t.Assert(client.GetContent("/john/list/1.java"), "john&1&1002")
		t.Assert(client.GetContent("/john/list/2.java"), "john&2&1002")
	})
}

func Test_Router_Hook_Priority(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/priority/show", func(r *ghttp.Request) {
		r.Response.Write("show")
	})

	s.BindHookHandlerByMap("/priority/:name", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.Response.Write("1")
		},
	})
	s.BindHookHandlerByMap("/priority/*any", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.Response.Write("2")
		},
	})
	s.BindHookHandlerByMap("/priority/show", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.Response.Write("3")
		},
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/priority/show"), "312show")
		t.Assert(client.GetContent("/priority/any/any"), "2")
		t.Assert(client.GetContent("/priority/name"), "12")
	})
}

func Test_Router_Hook_Multi(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/multi-hook", func(r *ghttp.Request) {
		r.Response.Write("show")
	})

	s.BindHookHandlerByMap("/multi-hook", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.Response.Write("1")
		},
	})
	s.BindHookHandlerByMap("/multi-hook", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: func(r *ghttp.Request) {
			r.Response.Write("2")
		},
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/multi-hook"), "12show")
	})
}
