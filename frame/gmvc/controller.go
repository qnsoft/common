// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// Package gmvc provides basic object classes for MVC.
package gmvc

import (
	"github.com/qnsoft/common/net/qn_http"
)

// Controller is used for controller register of qn_http.Server.
type Controller struct {
	Request  *qn_http.Request
	Response *qn_http.Response
	Server   *qn_http.Server
	Cookie   *qn_http.Cookie
	Session  *qn_http.Session
	View     *View
}

// Init is the callback function for each request initialization.
func (c *Controller) Init(r *qn_http.Request) {
	c.Request = r
	c.Response = r.Response
	c.Server = r.Server
	c.View = NewView(r.Response)
	c.Cookie = r.Cookie
	c.Session = r.Session
}

// Shut is the callback function for each request close.
func (c *Controller) Shut() {

}

// Exit equals to function Request.Exit().
func (c *Controller) Exit() {
	c.Request.Exit()
}
