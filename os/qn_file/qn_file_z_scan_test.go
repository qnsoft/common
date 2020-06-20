// Copyright 2017-2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"

	"github.com/qnsoft/common/os/qn_file"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_ScanDir(t *testing.T) {
	teatPath := qn_debug.TestDataPath()
	qn_test.C(t, func(t *qn_test.T) {
		files, err := qn_file.ScanDir(teatPath, "*", false)
		t.Assert(err, nil)
		t.AssertIN(teatPath+qn_file.Separator+"dir1", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir2", files)
		t.AssertNE(teatPath+qn_file.Separator+"dir1"+qn_file.Separator+"file1", files)
	})
	qn_test.C(t, func(t *qn_test.T) {
		files, err := qn_file.ScanDir(teatPath, "*", true)
		t.Assert(err, nil)
		t.AssertIN(teatPath+qn_file.Separator+"dir1", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir2", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir1"+qn_file.Separator+"file1", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir2"+qn_file.Separator+"file2", files)
	})
}

func Test_ScanDirFunc(t *testing.T) {
	teatPath := qn_debug.TestDataPath()
	qn_test.C(t, func(t *qn_test.T) {
		files, err := qn_file.ScanDirFunc(teatPath, "*", true, func(path string) string {
			if qn_file.Name(path) != "file1" {
				return ""
			}
			return path
		})
		t.Assert(err, nil)
		t.Assert(len(files), 1)
		t.Assert(qn_file.Name(files[0]), "file1")
	})
}

func Test_ScanDirFile(t *testing.T) {
	teatPath := qn_debug.TestDataPath()
	qn_test.C(t, func(t *qn_test.T) {
		files, err := qn_file.ScanDirFile(teatPath, "*", false)
		t.Assert(err, nil)
		t.Assert(len(files), 0)
	})
	qn_test.C(t, func(t *qn_test.T) {
		files, err := qn_file.ScanDirFile(teatPath, "*", true)
		t.Assert(err, nil)
		t.AssertNI(teatPath+qn_file.Separator+"dir1", files)
		t.AssertNI(teatPath+qn_file.Separator+"dir2", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir1"+qn_file.Separator+"file1", files)
		t.AssertIN(teatPath+qn_file.Separator+"dir2"+qn_file.Separator+"file2", files)
	})
}
