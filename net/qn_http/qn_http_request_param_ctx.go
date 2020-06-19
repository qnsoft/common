// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import (
	"context"
	"github.com/qnsoft/common/container/qn_var"
)

// Context is alias for function GetCtx.
// This function overwrites the http.Request.Context function.
// See GetCtx.
func (r *Request) Context() context.Context {
	if r.context == nil {
		r.context = r.Request.Context()
	}
	return r.context
}

// GetCtx retrieves and returns the request's context.
func (r *Request) GetCtx() context.Context {
	return r.Context()
}

// GetCtxVar retrieves and returns a Var with given key name.
// The optional parameter <def> specifies the default value of the Var if given <key>
// does not exist in the context.
func (r *Request) GetCtxVar(key interface{}, def ...interface{}) qn_var.Var {
	value := r.Context().Value(key)
	if value == nil && len(def) > 0 {
		value = def[0]
	}
	return qn_var.New(value)
}

// SetCtxVar sets custom parameter to context with key-value pair.
func (r *Request) SetCtxVar(key interface{}, value interface{}) {
	r.context = context.WithValue(r.Context(), key, value)
}
