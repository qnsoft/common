// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_view_test

import (
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/gview"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Encode_Parse(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New()
		v.SetPath(qn_debug.TestDataPath("tpl"))
		v.SetAutoEncode(true)
		result, err := v.Parse("encode.tpl", g.Map{
			"title": "<b>my title</b>",
		})
		t.Assert(err, nil)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}

func Test_Encode_ParseContent(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		v := gview.New()
		tplContent := qn_file.GetContents(qn_debug.TestDataPath("tpl", "encode.tpl"))
		v.SetAutoEncode(true)
		result, err := v.ParseContent(tplContent, g.Map{
			"title": "<b>my title</b>",
		})
		t.Assert(err, nil)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}
