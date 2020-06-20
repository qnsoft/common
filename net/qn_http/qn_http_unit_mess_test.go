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
	"github.com/qnsoft/common/net/ghttp"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_GetUrl(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/url", func(r *ghttp.Request) {
		r.Response.Write(r.GetUrl())
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", p)
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(prefix)

		t.Assert(client.GetContent("/url"), prefix+"/url")
	})
}
