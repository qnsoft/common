// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
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
	"github.com/qnsoft/common/text/qn.str"
)

func Test_Client_Request_13_Dump(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/hello", func(r *qn_http.Request) {
		r.Response.WriteHeader(200)
		r.Response.WriteJson(qn.Map{"field": "test_for_response_body"})
	})
	s.BindHandler("/hello2", func(r *qn_http.Request) {
		r.Response.WriteHeader(200)
		r.Response.Writeln(qn.Map{"field": "test_for_response_body"})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := qn_http.NewClient().SetPrefix(url).ContentJson()
		r, err := client.Post("/hello", qn.Map{"field": "test_for_request_body"})
		t.Assert(err, nil)
		dumpedText := r.RawRequest()
		t.Assert(qn.str.Contains(dumpedText, "test_for_request_body"), true)
		dumpedText2 := r.RawResponse()
		t.Assert(qn.str.Contains(dumpedText2, "test_for_response_body"), true)

		client2 := qn_http.NewClient().SetPrefix(url).ContentType("text/html")
		r2, err := client2.Post("/hello2", qn.Map{"field": "test_for_request_body"})
		t.Assert(err, nil)
		dumpedText3 := r2.RawRequest()
		t.Assert(qn.str.Contains(dumpedText3, "test_for_request_body"), true)
		dumpedText4 := r2.RawResponse()
		t.Assert(qn.str.Contains(dumpedText4, "test_for_request_body"), false)

	})

}
