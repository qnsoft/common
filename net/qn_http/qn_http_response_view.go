// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.
//

package qn_http

import (
	"github.com/qnsoft/common/os/qn_cfg"
	"github.com/qnsoft/common/os/qn_view"
	"github.com/qnsoft/common/util/gmode"
	qn_util "github.com/qnsoft/common/util/qn_util"
)

// WriteTpl parses and responses given template file.
// The parameter <params> specifies the template variables for parsing.
func (r *Response) WriteTpl(tpl string, params ...qn_view.Params) error {
	if b, err := r.ParseTpl(tpl, params...); err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	} else {
		r.Write(b)
	}
	return nil
}

// WriteTplDefault parses and responses the default template file.
// The parameter <params> specifies the template variables for parsing.
func (r *Response) WriteTplDefault(params ...qn_view.Params) error {
	if b, err := r.ParseTplDefault(params...); err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	} else {
		r.Write(b)
	}
	return nil
}

// WriteTplContent parses and responses the template content.
// The parameter <params> specifies the template variables for parsing.
func (r *Response) WriteTplContent(content string, params ...qn_view.Params) error {
	if b, err := r.ParseTplContent(content, params...); err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	} else {
		r.Write(b)
	}
	return nil
}

// ParseTpl parses given template file <tpl> with given template variables <params>
// and returns the parsed template content.
func (r *Response) ParseTpl(tpl string, params ...qn_view.Params) (string, error) {
	return r.Request.GetView().Parse(tpl, r.buildInVars(params...))
}

// ParseDefault parses the default template file with params.
func (r *Response) ParseTplDefault(params ...qn_view.Params) (string, error) {
	return r.Request.GetView().ParseDefault(r.buildInVars(params...))
}

// ParseTplContent parses given template file <file> with given template parameters <params>
// and returns the parsed template content.
func (r *Response) ParseTplContent(content string, params ...qn_view.Params) (string, error) {
	return r.Request.GetView().ParseContent(content, r.buildInVars(params...))
}

// buildInVars merges build-in variables into <params> and returns the new template variables.
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	m := qn_util.MapMergeCopy(params...)
	// Retrieve custom template variables from request object.
	qn_util.MapMerge(m, r.Request.viewParams, map[string]interface{}{
		"Form":    r.Request.GetFormMap(),
		"Query":   r.Request.GetQueryMap(),
		"Request": r.Request.GetMap(),
		"Cookie":  r.Request.Cookie.Map(),
		"Session": r.Request.Session.Map(),
	})
	// Note that it should assign no Config variable to template
	// if there's no configuration file.
	if c := qn_cfg.Instance(); c.Available() {
		m["Config"] = c.GetMap(".")
	}
	return m
}
