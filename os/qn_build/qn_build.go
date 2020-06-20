// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// Package gbuild manages the build-in variables from "gf build".
package qn_build

import (
	"runtime"

	"github.com/gogf/gf"
	"github.com/qnsoft/common/container/qn_var"
	"github.com/qnsoft/common/encoding/qn_base64"
	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/internal/json"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

var (
	builtInVarStr = ""                       // Raw variable base64 string.
	builtInVarMap = map[string]interface{}{} // Binary custom variable map decoded.
)

func init() {
	if builtInVarStr != "" {
		err := json.Unmarshal(qn_base64.MustDecodeString(builtInVarStr), &builtInVarMap)
		if err != nil {
			intlog.Error(err)
		}
		builtInVarMap["gfVersion"] = gf.VERSION
		builtInVarMap["goVersion"] = runtime.Version()
		intlog.Printf("build variables: %+v", builtInVarMap)
	} else {
		intlog.Print("no build variables")
	}
}

// Info returns the basic built information of the binary as map.
// Note that it should be used with gf-cli tool "gf build",
// which injects necessary information into the binary.
func Info() map[string]string {
	return map[string]string{
		"gf":   GetString("gfVersion"),
		"go":   GetString("goVersion"),
		"git":  GetString("builtGit"),
		"time": GetString("builtTime"),
	}
}

// Get retrieves and returns the build-in binary variable with given name.
func Get(name string, def ...interface{}) interface{} {
	if v, ok := builtInVarMap[name]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

// Get retrieves and returns the build-in binary variable of given name as qn_var.Var.
func GetVar(name string, def ...interface{}) qn_var.Var {
	return qn_var.New(Get(name, def...))
}

// GetString retrieves and returns the build-in binary variable of given name as string.
func GetString(name string, def ...interface{}) string {
	return qn_conv.String(Get(name, def...))
}

// Map returns the custom build-in variable map.
func Map() map[string]interface{} {
	return builtInVarMap
}
