// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"
)

func ExampleGetServer() {
	s := g.Server()
	s.BindHandler("/", func(r *qn_http.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}

func ExampleClientResponse_RawDump() {
	response, err := g.Client().Get("https://goframe.org")
	if err != nil {
		panic(err)
	}
	response.RawDump()
}
