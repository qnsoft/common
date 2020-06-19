// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import "github.com/qnsoft/common/container/qn_var"

// SetParam sets custom parameter with key-value pair.
func (r *Request) SetParam(key string, value interface{}) {
	if r.paramsMap == nil {
		r.paramsMap = make(map[string]interface{})
	}
	r.paramsMap[key] = value
}

// GetParam returns custom parameter with given name <key>.
// It returns <def> if <key> does not exist.
// It returns nil if <def> is not passed.
func (r *Request) GetParam(key string, def ...interface{}) interface{} {
	if r.paramsMap != nil {
		return r.paramsMap[key]
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

// GetParamVar returns custom parameter with given name <key> as qn_var.Var.
// It returns <def> if <key> does not exist.
// It returns nil if <def> is not passed.
func (r *Request) GetParamVar(key string, def ...interface{}) qn_var.Var {
	return qn_var.New(r.GetParam(key, def...))
}
