// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Client_Basic(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/hello", func(r *qn_http.Request) {
		r.Response.Write("hello")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := qn_http.NewClient()
		client.SetPrefix(url)

		t.Assert(qn_http.GetContent(""), ``)
		t.Assert(client.GetContent("/hello"), `hello`)

		_, err := qn_http.Post("")
		t.AssertNE(err, nil)
	})
}

func Test_Client_BasicAuth(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/auth", func(r *qn_http.Request) {
		if r.BasicAuth("admin", "admin") {
			r.Response.Write("1")
		} else {
			r.Response.Write("0")
		}
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(c.BasicAuth("admin", "123").GetContent("/auth"), "0")
		t.Assert(c.BasicAuth("admin", "admin").GetContent("/auth"), "1")
	})
}

func Test_Client_Cookie(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/cookie", func(r *qn_http.Request) {
		r.Response.Write(r.Cookie.Get("test"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		c.SetCookie("test", "0123456789")
		t.Assert(c.PostContent("/cookie"), "0123456789")
	})
}

func Test_Client_MapParam(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/map", func(r *qn_http.Request) {
		r.Response.Write(r.Get("test"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.GetContent("/map", qn.Map{"test": "1234567890"}), "1234567890")
	})
}

func Test_Client_Cookies(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/cookie", func(r *qn_http.Request) {
		r.Cookie.Set("test1", "1")
		r.Cookie.Set("test2", "2")
		r.Response.Write("ok")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		resp, err := c.Get("/cookie")
		t.Assert(err, nil)
		defer resp.Close()

		t.AssertNE(resp.Header.Get("Set-Cookie"), "")

		m := resp.GetCookieMap()
		t.Assert(len(m), 2)
		t.Assert(m["test1"], 1)
		t.Assert(m["test2"], 2)
		t.Assert(resp.GetCookie("test1"), 1)
		t.Assert(resp.GetCookie("test2"), 2)
	})
}

func Test_Client_Chain_Header(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/header1", func(r *qn_http.Request) {
		r.Response.Write(r.Header.Get("test1"))
	})
	s.BindHandler("/header2", func(r *qn_http.Request) {
		r.Response.Write(r.Header.Get("test2"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(c.Header(qn.MapStrStr{"test1": "1234567890"}).GetContent("/header1"), "1234567890")
		t.Assert(c.HeaderRaw("test1: 1234567890\ntest2: abcdefg").GetContent("/header1"), "1234567890")
		t.Assert(c.HeaderRaw("test1: 1234567890\ntest2: abcdefg").GetContent("/header2"), "abcdefg")
	})
}

func Test_Client_Chain_Context(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/context", func(r *qn_http.Request) {
		time.Sleep(1 * time.Second)
		r.Response.Write("ok")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
		t.Assert(c.Ctx(ctx).GetContent("/context"), "")

		ctx, _ = context.WithTimeout(context.Background(), 2000*time.Millisecond)
		t.Assert(c.Ctx(ctx).GetContent("/context"), "ok")
	})
}

func Test_Client_Chain_Timeout(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/timeout", func(r *qn_http.Request) {
		time.Sleep(1 * time.Second)
		r.Response.Write("ok")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(c.Timeout(100*time.Millisecond).GetContent("/timeout"), "")
		t.Assert(c.Timeout(2000*time.Millisecond).GetContent("/timeout"), "ok")
	})
}

func Test_Client_Chain_ContentJson(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/json", func(r *qn_http.Request) {
		r.Response.Write(r.Get("name"), r.Get("score"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(c.ContentJson().PostContent("/json", qn.Map{
			"name":  "john",
			"score": 100,
		}), "john100")
		t.Assert(c.ContentJson().PostContent("/json", `{"name":"john", "score":100}`), "john100")

		type User struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}
		t.Assert(c.ContentJson().PostContent("/json", User{"john", 100}), "john100")
	})
}

func Test_Client_Chain_ContentXml(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/xml", func(r *qn_http.Request) {
		r.Response.Write(r.Get("name"), r.Get("score"))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_http.NewClient()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		t.Assert(c.ContentXml().PostContent("/xml", qn.Map{
			"name":  "john",
			"score": 100,
		}), "john100")
		t.Assert(c.ContentXml().PostContent("/xml", `{"name":"john", "score":100}`), "john100")

		type User struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}
		t.Assert(c.ContentXml().PostContent("/xml", User{"john", 100}), "john100")
	})
}
