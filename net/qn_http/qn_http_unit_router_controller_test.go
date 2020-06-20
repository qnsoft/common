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
	"github.com/qnsoft/common/frame/gmvc"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

// 控制器
type Controller struct {
	gmvc.Controller
}

func (c *Controller) Init(r *qn_http.Request) {
	c.Controller.Init(r)
	c.Response.Write("1")
}

func (c *Controller) Shut() {
	c.Response.Write("2")
}

func (c *Controller) Index() {
	c.Response.Write("Controller Index")
}

func (c *Controller) Show() {
	c.Response.Write("Controller Show")
}

func (c *Controller) Info() {
	c.Response.Write("Controller Info")
}

func Test_Router_Controller1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindController("/", new(Controller))
	s.BindController("/{.struct}/{.method}", new(Controller))
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "1Controller Index2")
		t.Assert(client.GetContent("/init"), "Not Found")
		t.Assert(client.GetContent("/shut"), "Not Found")
		t.Assert(client.GetContent("/index"), "1Controller Index2")
		t.Assert(client.GetContent("/show"), "1Controller Show2")

		t.Assert(client.GetContent("/controller"), "Not Found")
		t.Assert(client.GetContent("/controller/init"), "Not Found")
		t.Assert(client.GetContent("/controller/shut"), "Not Found")
		t.Assert(client.GetContent("/controller/index"), "1Controller Index2")
		t.Assert(client.GetContent("/controller/show"), "1Controller Show2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_Controller2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindController("/controller", new(Controller), "Show, Info")
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/controller"), "Not Found")
		t.Assert(client.GetContent("/controller/init"), "Not Found")
		t.Assert(client.GetContent("/controller/shut"), "Not Found")
		t.Assert(client.GetContent("/controller/index"), "Not Found")
		t.Assert(client.GetContent("/controller/show"), "1Controller Show2")
		t.Assert(client.GetContent("/controller/info"), "1Controller Info2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_ControllerMethod(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindControllerMethod("/controller-info", new(Controller), "Info")
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/controller"), "Not Found")
		t.Assert(client.GetContent("/controller/init"), "Not Found")
		t.Assert(client.GetContent("/controller/shut"), "Not Found")
		t.Assert(client.GetContent("/controller/index"), "Not Found")
		t.Assert(client.GetContent("/controller/show"), "Not Found")
		t.Assert(client.GetContent("/controller/info"), "Not Found")
		t.Assert(client.GetContent("/controller-info"), "1Controller Info2")

		t.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
