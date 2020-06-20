// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_res_test

import (
	"strings"
	"testing"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/os/qn_res"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_PackToGoFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		goFilePath := qn_debug.TestDataPath("testdata.go")
		pkgName := "testdata"
		err := qn_res.PackToGoFile(srcPath, goFilePath, pkgName)
		t.Assert(err, nil)
	})
}

func Test_Pack(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		data, err := qn_res.Pack(srcPath)
		t.Assert(err, nil)

		r := qn_res.New()

		err = r.Add(string(data))
		t.Assert(err, nil)
		t.Assert(r.Contains("files/"), true)
	})
}

func Test_PackToFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		dstPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err := qn_res.PackToFile(srcPath, dstPath)
		t.Assert(err, nil)

		defer qn_file.Remove(dstPath)

		r := qn_res.New()
		err = r.Load(dstPath)
		t.Assert(err, nil)
		t.Assert(r.Contains("files"), true)
	})
}

func Test_PackMulti(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		goFilePath := qn_debug.TestDataPath("data/data.go")
		pkgName := "data"
		array, err := qn_file.ScanDir(srcPath, "*", false)
		t.Assert(err, nil)
		err = qn_res.PackToGoFile(strings.Join(array, ","), goFilePath, pkgName)
		t.Assert(err, nil)
	})
}

func Test_PackWithPrefix1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		goFilePath := qn_file.TempDir("testdata.go")
		pkgName := "testdata"
		err := qn_res.PackToGoFile(srcPath, goFilePath, pkgName, "www/gf-site/test")
		t.Assert(err, nil)
	})
}

func Test_PackWithPrefix2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("files")
		goFilePath := qn_file.TempDir("testdata.go")
		pkgName := "testdata"
		err := qn_res.PackToGoFile(srcPath, goFilePath, pkgName, "/var/www/gf-site/test")
		t.Assert(err, nil)
	})
}
