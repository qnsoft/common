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

	"github.com/qnsoft/common/debug/gdebug"
	"github.com/qnsoft/common/os/gres"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_PackToGoFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		goFilePath := gdebug.TestDataPath("testdata.go")
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName)
		t.Assert(err, nil)
	})
}

func Test_Pack(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		data, err := gres.Pack(srcPath)
		t.Assert(err, nil)

		r := gres.New()

		err = r.Add(string(data))
		t.Assert(err, nil)
		t.Assert(r.Contains("files/"), true)
	})
}

func Test_PackToFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		dstPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err := gres.PackToFile(srcPath, dstPath)
		t.Assert(err, nil)

		defer qn_file.Remove(dstPath)

		r := gres.New()
		err = r.Load(dstPath)
		t.Assert(err, nil)
		t.Assert(r.Contains("files"), true)
	})
}

func Test_PackMulti(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		goFilePath := gdebug.TestDataPath("data/data.go")
		pkgName := "data"
		array, err := qn_file.ScanDir(srcPath, "*", false)
		t.Assert(err, nil)
		err = gres.PackToGoFile(strings.Join(array, ","), goFilePath, pkgName)
		t.Assert(err, nil)
	})
}

func Test_PackWithPrefix1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		goFilePath := qn_file.TempDir("testdata.go")
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName, "www/gf-site/test")
		t.Assert(err, nil)
	})
}

func Test_PackWithPrefix2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := gdebug.TestDataPath("files")
		goFilePath := qn_file.TempDir("testdata.go")
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName, "/var/www/gf-site/test")
		t.Assert(err, nil)
	})
}
