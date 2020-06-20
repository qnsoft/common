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
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Router_Group_Hook1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	group := s.Group("/api")
	group.GET("/handler", func(r *qn_http.Request) {
		r.Response.Write("1")
	})
	group.ALL("/handler", func(r *qn_http.Request) {
		r.Response.Write("0")
	}, qn_http.HOOK_BEFORE_SERVE)
	group.ALL("/handler", func(r *qn_http.Request) {
		r.Response.Write("2")
	}, qn_http.HOOK_AFTER_SERVE)

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "02")
		t.Assert(client.GetContent("/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	g := s.Group("/api")
	g.GET("/handler", func(r *qn_http.Request) {
		r.Response.Write("1")
	})
	g.GET("/*", func(r *qn_http.Request) {
		r.Response.Write("0")
	}, qn_http.HOOK_BEFORE_SERVE)
	g.GET("/*", func(r *qn_http.Request) {
		r.Response.Write("2")
	}, qn_http.HOOK_AFTER_SERVE)

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "Not Found")
		t.Assert(client.GetContent("/api/ThisDoesNotExist"), "02")
		t.Assert(client.PostContent("/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Hook3(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/api").Bind([]qn.Slice{
		{"ALL", "handler", func(r *qn_http.Request) {
			r.Response.Write("1")
		}},
		{"ALL", "/*", func(r *qn_http.Request) {
			r.Response.Write("0")
		}, qn_http.HOOK_BEFORE_SERVE},
		{"ALL", "/*", func(r *qn_http.Request) {
			r.Response.Write("2")
		}, qn_http.HOOK_AFTER_SERVE},
	})

	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(client.GetContent("/api/handler"), "012")
		t.Assert(client.PostContent("/api/handler"), "012")
		t.Assert(client.DeleteContent("/api/ThisDoesNotExist"), "02")
	})
}
