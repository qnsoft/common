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

func Test_Router_Exit(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHookHandlerByMap("/*", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE:  func(r *qn_http.Request) { r.Response.Write("1") },
		qn_http.HOOK_AFTER_SERVE:   func(r *qn_http.Request) { r.Response.Write("2") },
		qn_http.HOOK_BEFORE_OUTPUT: func(r *qn_http.Request) { r.Response.Write("3") },
		qn_http.HOOK_AFTER_OUTPUT:  func(r *qn_http.Request) { r.Response.Write("4") },
	})
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test-start")
		r.Exit()
		r.Response.Write("test-end")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "123")
		t.Assert(client.GetContent("/test/test"), "1test-start23")
	})
}

func Test_Router_ExitHook(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/priority/show", func(r *qn_http.Request) {
		r.Response.Write("show")
	})

	s.BindHookHandlerByMap("/priority/:name", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("1")
		},
	})
	s.BindHookHandlerByMap("/priority/*any", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("2")
		},
	})
	s.BindHookHandlerByMap("/priority/show", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("3")
			r.ExitHook()
		},
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
		t.Assert(client.GetContent("/priority/show"), "3show")
	})
}

func Test_Router_ExitAll(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/priority/show", func(r *qn_http.Request) {
		r.Response.Write("show")
	})

	s.BindHookHandlerByMap("/priority/:name", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("1")
		},
	})
	s.BindHookHandlerByMap("/priority/*any", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("2")
		},
	})
	s.BindHookHandlerByMap("/priority/show", map[string]qn_http.HandlerFunc{
		qn_http.HOOK_BEFORE_SERVE: func(r *qn_http.Request) {
			r.Response.Write("3")
			r.ExitAll()
		},
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
		t.Assert(client.GetContent("/priority/show"), "3")
	})
}
