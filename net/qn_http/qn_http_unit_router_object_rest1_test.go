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

type ObjectRest struct{}

func (o *ObjectRest) Init(r *qn_http.Request) {
	r.Response.Write("1")
}

func (o *ObjectRest) Shut(r *qn_http.Request) {
	r.Response.Write("2")
}

func (o *ObjectRest) Get(r *qn_http.Request) {
	r.Response.Write("Object Get")
}

func (o *ObjectRest) Put(r *qn_http.Request) {
	r.Response.Write("Object Put")
}

func (o *ObjectRest) Post(r *qn_http.Request) {
	r.Response.Write("Object Post")
}

func (o *ObjectRest) Delete(r *qn_http.Request) {
	r.Response.Write("Object Delete")
}

func (o *ObjectRest) Patch(r *qn_http.Request) {
	r.Response.Write("Object Patch")
}

func (o *ObjectRest) Options(r *qn_http.Request) {
	r.Response.Write("Object Options")
}

func (o *ObjectRest) Head(r *qn_http.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_ObjectRest(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObjectRest("/", new(ObjectRest))
	s.BindObjectRest("/{.struct}/{.method}", new(ObjectRest))
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "1Object Get2")
		t.Assert(client.PutContent("/"), "1Object Put2")
		t.Assert(client.PostContent("/"), "1Object Post2")
		t.Assert(client.DeleteContent("/"), "1Object Delete2")
		t.Assert(client.PatchContent("/"), "1Object Patch2")
		t.Assert(client.OptionsContent("/"), "1Object Options2")
		resp1, err := client.Head("/")
		if err == nil {
			defer resp1.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp1.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/object-rest/get"), "1Object Get2")
		t.Assert(client.PutContent("/object-rest/put"), "1Object Put2")
		t.Assert(client.PostContent("/object-rest/post"), "1Object Post2")
		t.Assert(client.DeleteContent("/object-rest/delete"), "1Object Delete2")
		t.Assert(client.PatchContent("/object-rest/patch"), "1Object Patch2")
		t.Assert(client.OptionsContent("/object-rest/options"), "1Object Options2")
		resp2, err := client.Head("/object-rest/head")
		if err == nil {
			defer resp2.Close()
		}
		t.Assert(err, nil)
		t.Assert(resp2.Header.Get("head-ok"), "1")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
