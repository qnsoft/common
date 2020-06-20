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

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/text/gstr"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_HTTPS_Basic(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.GET("/test", func(r *qn_http.Request) {
			r.Response.Write("test")
		})
	})
	s.EnableHTTPS(
		qn_debug.TestDataPath("https", "server.crt"),
		qn_debug.TestDataPath("https", "server.key"),
	)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.AssertIN(gstr.Trim(c.GetContent("/")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
		t.AssertIN(gstr.Trim(c.GetContent("/test")), g.Slice{"", "Client sent an HTTP request to an HTTPS server."})
	})
	// HTTPS
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("https://127.0.0.1:%d", p))
		t.Assert(c.GetContent("/"), "Not Found")
		t.Assert(c.GetContent("/test"), "test")
	})
}

func Test_HTTPS_HTTP_Basic(t *testing.T) {
	var (
		portHttp, _  = ports.PopRand()
		portHttps, _ = ports.PopRand()
	)
	s := g.Server(qn_time.TimestampNanoStr())
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.GET("/test", func(r *qn_http.Request) {
			r.Response.Write("test")
		})
	})
	s.EnableHTTPS(
		qn_debug.TestDataPath("https", "server.crt"),
		qn_debug.TestDataPath("https", "server.key"),
	)
	s.SetPort(portHttp)
	s.SetHTTPSPort(portHttps)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// HTTP
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", portHttp))
		t.Assert(c.GetContent("/"), "Not Found")
		t.Assert(c.GetContent("/test"), "test")
	})
	// HTTPS
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("https://127.0.0.1:%d", portHttps))
		t.Assert(c.GetContent("/"), "Not Found")
		t.Assert(c.GetContent("/test"), "test")
	})
}
