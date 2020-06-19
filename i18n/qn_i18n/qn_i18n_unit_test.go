// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_i18n_test

import (
	"testing"

	"github.com/qnsoft/common/os/gres"

	"github.com/qnsoft/common/os/gtime"
	"github.com/qnsoft/common/util/gconv"

	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/i18n/gi18n"

	"github.com/qnsoft/common/debug/gdebug"
	"github.com/qnsoft/common/os/gfile"

	"github.com/qnsoft/common/test/gtest"

	_ "github.com/qnsoft/common/os/gres/testdata/data"
)

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.TestDataPath("i18n"),
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

	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.TestDataPath("i18n-file"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T("{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T("{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T("{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n-dir",
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
	gtest.C(t, func(t *gtest.T) {
		err := gi18n.SetPath(gdebug.TestDataPath("i18n"))
		t.Assert(err, nil)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.T("{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.T("{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.T("{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gi18n.SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n-dir")
		t.Assert(err, nil)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.Translate("{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.Translate("{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.Translate("{#hello}{#world}"), "你好世界")
	})
}

func Test_Instance(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance()
		err := m.SetPath("i18n-dir")
		t.Assert(err, nil)
		m.SetLanguage("zh-CN")
		t.Assert(m.T("{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance()
		t.Assert(m.T("{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.I18n().T("{#hello}{#world}"), "你好世界")
	})
	// Default language is: en
	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance(gconv.String(gtime.TimestampNano()))
		t.Assert(m.T("{#hello}{#world}"), "HelloWorld")
	})
}

func Test_Resource(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
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
