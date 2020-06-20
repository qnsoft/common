// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// static service testing.

package qn_http_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/encoding/qn_html"
	"github.com/qnsoft/common/os/gview"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Template_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New(qn_debug.TestDataPath("template", "basic"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("index.html", g.Map{
				"name": "john",
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Name:john")
		t.Assert(client.GetContent("/"), "Name:john")
	})
}

func Test_Template_Encode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New(qn_debug.TestDataPath("template", "basic"))
		v.SetAutoEncode(true)
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("index.html", g.Map{
				"name": "john",
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Name:john")
		t.Assert(client.GetContent("/"), "Name:john")
	})
}

func Test_Template_Layout1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New(qn_debug.TestDataPath("template", "layout1"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/layout", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/nil", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("layout.html", nil)
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/layout"), "123")
		t.Assert(client.GetContent("/nil"), "123")
	})
}

func Test_Template_Layout2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New(qn_debug.TestDataPath("template", "layout2"))
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/main1", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main1.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/main2", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("layout.html", g.Map{
				"mainTpl": "main/main2.html",
			})
			t.Assert(err, nil)
		})
		s.BindHandler("/nil", func(r *qn_http.Request) {
			err := r.Response.WriteTpl("layout.html", nil)
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/main1"), "a1b")
		t.Assert(client.GetContent("/main2"), "a2b")
		t.Assert(client.GetContent("/nil"), "ab")
	})
}

func Test_Template_XSS(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New()
		v.SetAutoEncode(true)
		c := "<br>"
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetView(v)
		s.BindHandler("/", func(r *qn_http.Request) {
			err := r.Response.WriteTplContent("{{if eq 1 1}}{{.v}}{{end}}", g.Map{
				"v": c,
			})
			t.Assert(err, nil)
		})
		s.SetDumpRouterMap(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), qn_html.Entities(c))
	})
}
