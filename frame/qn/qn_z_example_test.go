// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_test

import (
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/net/qn_http"
)

func ExampleServer() {
	// A hello world example.
	s := qn.Server()
	s.BindHandler("/", func(r *qn_http.Request) {
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}
