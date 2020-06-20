// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import "github.com/qnsoft/common/os/qn_view"

// SetView sets template view engine object for this request.
func (r *Request) SetView(view *qn_view.View) {
	r.viewObject = view
}

// GetView returns the template view engine object for this request.
func (r *Request) GetView() *qn_view.View {
	view := r.viewObject
	if view == nil {
		view = r.Server.config.View
	}
	if view == nil {
		view = qn_view.Instance()
	}
	return view
}

// Assigns binds multiple template variables to current request.
func (r *Request) Assigns(data qn_view.Params) {
	if r.viewParams == nil {
		r.viewParams = make(qn_view.Params, len(data))
	}
	for k, v := range data {
		r.viewParams[k] = v
	}
}

// Assign binds a template variable to current request.
func (r *Request) Assign(key string, value interface{}) {
	if r.viewParams == nil {
		r.viewParams = make(qn_view.Params)
	}
	r.viewParams[key] = value
}
