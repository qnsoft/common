// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_view_test

import (
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/qn_view"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Config(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		config := qn_view.Config{
			Paths: []string{qn_debug.TestDataPath("config")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}
		view := qn_view.New()
		err := view.SetConfig(config)
		t.Assert(err, nil)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "name:gf")
	})
}

func Test_ConfigWithMap(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		view := qn_view.New()
		err := view.SetConfigWithMap(g.Map{
			"Paths":       []string{qn_debug.TestDataPath("config")},
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "gf",
			},
		})
		t.Assert(err, nil)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "name:gf")
	})
}
