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

func Test_Session_Cookie(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *qn_http.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
	})
	s.BindHandler("/get", func(r *qn_http.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *qn_http.Request) {
		r.Session.Remove(r.GetString("k"))
	})
	s.BindHandler("/clear", func(r *qn_http.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r1, e1 := client.Get("/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}
		t.Assert(e1, nil)
		t.Assert(r1.ReadAllString(), "")

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/clear"), "")
		t.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_Header(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *qn_http.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
	})
	s.BindHandler("/get", func(r *qn_http.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.BindHandler("/remove", func(r *qn_http.Request) {
		r.Session.Remove(r.GetString("k"))
	})
	s.BindHandler("/clear", func(r *qn_http.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key1&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId := response.GetCookie(s.GetSessionIdName())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.ReadAllString(), "")

		client.SetHeader(s.GetSessionIdName(), sessionId)

		t.Assert(client.GetContent("/set?k=key2&v=200"), "")

		t.Assert(client.GetContent("/get?k=key1"), "100")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/get?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key1"), "")
		t.Assert(client.GetContent("/remove?k=key3"), "")
		t.Assert(client.GetContent("/remove?k=key4"), "")
		t.Assert(client.GetContent("/get?k=key1"), "")
		t.Assert(client.GetContent("/get?k=key2"), "200")
		t.Assert(client.GetContent("/clear"), "")
		t.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_StorageFile(t *testing.T) {
	sessionId := ""
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *qn_http.Request) {
		r.Session.Set(r.GetString("k"), r.GetString("v"))
		r.Response.Write(r.GetString("k"), "=", r.GetString("v"))
	})
	s.BindHandler("/get", func(r *qn_http.Request) {
		r.Response.Write(r.Session.Get(r.GetString("k")))
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId = response.GetCookie(s.GetSessionIdName())
		t.Assert(e1, nil)
		t.AssertNE(sessionId, nil)
		t.Assert(response.ReadAllString(), "key=100")
	})
	time.Sleep(time.Second)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		client.SetHeader(s.GetSessionIdName(), sessionId)
		t.Assert(client.GetContent("/get?k=key"), "100")
		t.Assert(client.GetContent("/get?k=key1"), "")
	})
}
