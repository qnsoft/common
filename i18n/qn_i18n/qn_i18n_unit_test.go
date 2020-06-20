// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_i18n_test

import (
	"testing"

	"github.com/qnsoft/common/os/qn_res"
	"github.com/qnsoft/common/os/qn_time"
	qn_conv "github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/i18n/qn_i18n"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/os/qn_file"

	"github.com/qnsoft/common/test/qn_test"

	_ "github.com/qnsoft/common/os/qn_res/testdata/data"
)

func Test_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		i18n := qn_i18n.New(qn_i18n.Options{
			Path: qn_debug.TestDataPath("i18n"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T("{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T("{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T("{#hello}{#world}"), "你好世界")
		i18n.SetDelimiters("{$", "}")
		t.Assert(i18n.T("{#hello}{#world}"), "{#hello}{#world}")
		t.Assert(i18n.T("{$hello}{$world}"), "你好世界")
	})

	qn_test.C(t, func(t *qn_test.T) {
		i18n := qn_i18n.New(qn_i18n.Options{
			Path: qn_debug.TestDataPath("i18n-file"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T("{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T("{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T("{#hello}{#world}"), "你好世界")
	})

	qn_test.C(t, func(t *qn_test.T) {
		i18n := qn_i18n.New(qn_i18n.Options{
			Path: qn_debug.CallerDirectory() + qn_file.Separator + "testdata" + qn_file.Separator + "i18n-dir",
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T("{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T("{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T("{#hello}{#world}"), "你好世界")
	})
}

func Test_DefaultManager(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		err := qn_i18n.SetPath(qn_debug.TestDataPath("i18n"))
		t.Assert(err, nil)

		qn_i18n.SetLanguage("none")
		t.Assert(qn_i18n.T("{#hello}{#world}"), "{#hello}{#world}")

		qn_i18n.SetLanguage("ja")
		t.Assert(qn_i18n.T("{#hello}{#world}"), "こんにちは世界")

		qn_i18n.SetLanguage("zh-CN")
		t.Assert(qn_i18n.T("{#hello}{#world}"), "你好世界")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_i18n.SetPath(qn_debug.CallerDirectory() + qn_file.Separator + "testdata" + qn_file.Separator + "i18n-dir")
		t.Assert(err, nil)

		qn_i18n.SetLanguage("none")
		t.Assert(qn_i18n.Translate("{#hello}{#world}"), "{#hello}{#world}")

		qn_i18n.SetLanguage("ja")
		t.Assert(qn_i18n.Translate("{#hello}{#world}"), "こんにちは世界")

		qn_i18n.SetLanguage("zh-CN")
		t.Assert(qn_i18n.Translate("{#hello}{#world}"), "你好世界")
	})
}

func Test_Instance(t *testing.T) {
	qn_res.Dump()
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_i18n.Instance()
		err := m.SetPath("i18n-dir")
		t.Assert(err, nil)
		m.SetLanguage("zh-CN")
		t.Assert(m.T("{#hello}{#world}"), "你好世界")
	})

	qn_test.C(t, func(t *qn_test.T) {
		m := qn_i18n.Instance()
		t.Assert(m.T("{#hello}{#world}"), "你好世界")
	})

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(g.I18n().T("{#hello}{#world}"), "你好世界")
	})
	// Default language is: en
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_i18n.Instance(qn_conv.String(qn_time.TimestampNano()))
		t.Assert(m.T("{#hello}{#world}"), "HelloWorld")
	})
}

func Test_Resource(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := g.I18n("resource")
		err := m.SetPath("i18n-dir")
		t.Assert(err, nil)

		m.SetLanguage("none")
		t.Assert(m.T("{#hello}{#world}"), "{#hello}{#world}")

		m.SetLanguage("ja")
		t.Assert(m.T("{#hello}{#world}"), "こんにちは世界")

		m.SetLanguage("zh-CN")
		t.Assert(m.T("{#hello}{#world}"), "你好世界")
	})
}
