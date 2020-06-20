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

type Object struct{}

func (o *Object) Init(r *qn_http.Request) {
	r.Response.Write("1")
}

func (o *Object) Shut(r *qn_http.Request) {
	r.Response.Write("2")
}

func (o *Object) Index(r *qn_http.Request) {
	r.Response.Write("Object Index")
}

func (o *Object) Show(r *qn_http.Request) {
	r.Response.Write("Object Show")
}

func (o *Object) Info(r *qn_http.Request) {
	r.Response.Write("Object Info")
}

func Test_Router_Object1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObject("/", new(Object))
	s.BindObject("/{.struct}/{.method}", new(Object))
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "1Object Index2")
		t.Assert(client.GetContent("/init"), "Not Found")
		t.Assert(client.GetContent("/shut"), "Not Found")
		t.Assert(client.GetContent("/index"), "1Object Index2")
		t.Assert(client.GetContent("/show"), "1Object Show2")

		t.Assert(client.GetContent("/object"), "Not Found")
		t.Assert(client.GetContent("/object/init"), "Not Found")
		t.Assert(client.GetContent("/object/shut"), "Not Found")
		t.Assert(client.GetContent("/object/index"), "1Object Index2")
		t.Assert(client.GetContent("/object/show"), "1Object Show2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_Object2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObject("/object", new(Object), "Show, Info")
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/object"), "Not Found")
		t.Assert(client.GetContent("/object/init"), "Not Found")
		t.Assert(client.GetContent("/object/shut"), "Not Found")
		t.Assert(client.GetContent("/object/index"), "Not Found")
		t.Assert(client.GetContent("/object/show"), "1Object Show2")
		t.Assert(client.GetContent("/object/info"), "1Object Info2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_ObjectMethod(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObjectMethod("/object-info", new(Object), "Info")
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/object"), "Not Found")
		t.Assert(client.GetContent("/object/init"), "Not Found")
		t.Assert(client.GetContent("/object/shut"), "Not Found")
		t.Assert(client.GetContent("/object/index"), "Not Found")
		t.Assert(client.GetContent("/object/show"), "Not Found")
		t.Assert(client.GetContent("/object/info"), "Not Found")
		t.Assert(client.GetContent("/object-info"), "1Object Info2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
