// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// package qn_view implements a template engine based on text/template.
//
// Reserved template variable names:
//     I18nLanguage: Assign this variable to define i18n language for each page.
package qn_view

import (
	"github.com/gogf/gf"
	"github.com/qnsoft/common/container/gmap"
	"github.com/qnsoft/common/internal/intlog"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/internal/cmdenv"
	"github.com/qnsoft/common/os/glog"
	"github.com/qnsoft/common/os/qn_file"
)

// View object for template engine.
type View struct {
	paths        *qn_array.StrArray     // Searching array for path, NOT concurrent-safe for performance purpose.
	data         map[string]interface{} // Global template variables.
	funcMap      map[string]interface{} // Global template function map.
	fileCacheMap *gmap.StrAnyMap        // File cache map.
	config       Config                 // Extra configuration for the view.
}

type (
	Params  = map[string]interface{} // Params is type for template params.
	FuncMap = map[string]interface{} // FuncMap is type for custom template functions.
)

var (
	// Default view object.
	defaultViewObj *View
)

// checkAndInitDefaultView checks and initializes the default view object.
// The default view object will be initialized just once.
func checkAndInitDefaultView() {
	if defaultViewObj == nil {
		defaultViewObj = New()
	}
}

// ParseContent parses the template content directly using the default view object
// and returns the parsed content.
func ParseContent(content string, params ...Params) (string, error) {
	checkAndInitDefaultView()
	return defaultViewObj.ParseContent(content, params...)
}

// New returns a new view object.
// The parameter <path> specifies the template directory path to load template files.
func New(path ...string) *View {
	view := &View{
		paths:        qn_array.NewStrArray(),
		data:         make(map[string]interface{}),
		funcMap:      make(map[string]interface{}),
		fileCacheMap: gmap.NewStrAnyMap(true),
		config:       DefaultConfig(),
	}
	if len(path) > 0 && len(path[0]) > 0 {
		if err := view.SetPath(path[0]); err != nil {
			intlog.Error(err)
		}
	} else {
		// Customized dir path from env/cmd.
		if envPath := cmdenv.Get("gf.gview.path").String(); envPath != "" {
			if qn_file.Exists(envPath) {
				if err := view.SetPath(envPath); err != nil {
					intlog.Error(err)
				}
			} else {
				if errorPrint() {
					glog.Errorf("Template directory path does not exist: %s", envPath)
				}
			}
		} else {
			// Dir path of working dir.
			if err := view.SetPath(qn_file.Pwd()); err != nil {
				intlog.Error(err)
			}
			// Dir path of binary.
			if selfPath := qn_file.SelfDir(); selfPath != "" && qn_file.Exists(selfPath) {
				if err := view.AddPath(selfPath); err != nil {
					intlog.Error(err)
				}
			}
			// Dir path of main package.
			if mainPath := qn_file.MainPkgPath(); mainPath != "" && qn_file.Exists(mainPath) {
				if err := view.AddPath(mainPath); err != nil {
					intlog.Error(err)
				}
			}
		}
	}
	view.SetDelimiters("{{", "}}")
	// default build-in variables.
	view.data["GF"] = map[string]interface{}{
		"version": gf.VERSION,
	}
	// default build-in functions.
	view.BindFuncMap(FuncMap{
		"eq":         view.funcEq,
		"ne":         view.funcNe,
		"lt":         view.funcLt,
		"le":         view.funcLe,
		"gt":         view.funcGt,
		"ge":         view.funcGe,
		"text":       view.funcText,
		"html":       view.funcHtmlEncode,
		"htmlencode": view.funcHtmlEncode,
		"htmldecode": view.funcHtmlDecode,
		"encode":     view.funcHtmlEncode,
		"decode":     view.funcHtmlDecode,
		"url":        view.funcUrlEncode,
		"urlencode":  view.funcUrlEncode,
		"urldecode":  view.funcUrlDecode,
		"date":       view.funcDate,
		"substr":     view.funcSubStr,
		"strlimit":   view.funcStrLimit,
		"concat":     view.funcConcat,
		"replace":    view.funcReplace,
		"compare":    view.funcCompare,
		"hidestr":    view.funcHideStr,
		"highlight":  view.funcHighlight,
		"toupper":    view.funcToUpper,
		"tolower":    view.funcToLower,
		"nl2br":      view.funcNl2Br,
		"include":    view.funcInclude,
		"dump":       view.funcDump,
	})

	return view
}
