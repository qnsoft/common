// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"fmt"
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/os/qn_cfg"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_View(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.AssertNE(View(), nil)
		b, e := View().ParseContent(`{{"我是中国人" | substr 2 -1}}`, nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
	qn_test.C(t, func(t *qn_test.T) {
		tpl := "t.tpl"
		err := qn_file.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.Assert(err, nil)
		defer qn_file.Remove(tpl)

		b, e := View().Parse("t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := fmt.Sprintf(`%s/%d`, qn_file.TempDir(), qn_time.TimestampNano())
		tpl := fmt.Sprintf(`%s/%s`, path, "t.tpl")
		err := qn_file.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.Assert(err, nil)
		defer qn_file.Remove(tpl)
		err = View().AddPath(path)
		t.Assert(err, nil)

		b, e := View().Parse("t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国人")
	})
}

func Test_View_Config(t *testing.T) {
	// view1 test1
	qn_test.C(t, func(t *qn_test.T) {
		dirPath := qn_debug.TestDataPath("view1")
		qn_cfg.SetContent(qn_file.GetContents(qn_file.Join(dirPath, "config.toml")))
		defer qn_cfg.ClearContent()
		defer instances.Clear()

		view := View("test1")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello ${.name},version:${.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test1,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test1:test1")
	})
	// view1 test2
	qn_test.C(t, func(t *qn_test.T) {
		dirPath := qn_debug.TestDataPath("view1")
		qn_cfg.SetContent(qn_file.GetContents(qn_file.Join(dirPath, "config.toml")))
		defer qn_cfg.ClearContent()
		defer instances.Clear()

		view := View("test2")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello #{.name},version:#{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test2,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test2:test2")
	})
	// view2
	qn_test.C(t, func(t *qn_test.T) {
		dirPath := qn_debug.TestDataPath("view2")
		qn_cfg.SetContent(qn_file.GetContents(qn_file.Join(dirPath, "config.toml")))
		defer qn_cfg.ClearContent()
		defer instances.Clear()

		view := View()
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test:test")
	})
	// view2
	qn_test.C(t, func(t *qn_test.T) {
		dirPath := qn_debug.TestDataPath("view2")
		qn_cfg.SetContent(qn_file.GetContents(qn_file.Join(dirPath, "config.toml")))
		defer qn_cfg.ClearContent()
		defer instances.Clear()

		view := View("test100")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.Assert(err, nil)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(str, nil)
		t.Assert(err, nil)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault()
		t.Assert(err, nil)
		t.Assert(result, "test:test")
	})
}
