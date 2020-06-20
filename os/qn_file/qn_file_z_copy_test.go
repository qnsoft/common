// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"testing"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Copy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(qn_file.Copy(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(qn_file.IsFile(testpath()+topath), true)
		t.AssertNE(qn_file.Copy("", ""), nil)
	})
}

func Test_CopyFile(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(qn_file.CopyFile(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(qn_file.IsFile(testpath()+topath), true)
		t.AssertNE(qn_file.CopyFile("", ""), nil)
	})
	// Content replacement.
	qn_test.C(t, func(t *qn_test.T) {
		src := qn_file.TempDir(qn_time.TimestampNanoStr())
		dst := qn_file.TempDir(qn_time.TimestampNanoStr())
		srcContent := "1"
		dstContent := "1"
		t.Assert(qn_file.PutContents(src, srcContent), nil)
		t.Assert(qn_file.PutContents(dst, dstContent), nil)
		t.Assert(qn_file.GetContents(src), srcContent)
		t.Assert(qn_file.GetContents(dst), dstContent)

		t.Assert(qn_file.CopyFile(src, dst), nil)
		t.Assert(qn_file.GetContents(src), srcContent)
		t.Assert(qn_file.GetContents(dst), srcContent)
	})
}

func Test_CopyDir(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			dirPath1 = "/test-copy-dir1"
			dirPath2 = "/test-copy-dir2"
		)
		haveList := []string{
			"t1.txt",
			"t2.txt",
		}
		createDir(dirPath1)
		for _, v := range haveList {
			t.Assert(createTestFile(dirPath1+"/"+v, ""), nil)
		}
		defer delTestFiles(dirPath1)

		var (
			yfolder  = testpath() + dirPath1
			tofolder = testpath() + dirPath2
		)

		if qn_file.IsDir(tofolder) {
			t.Assert(qn_file.Remove(tofolder), nil)
			t.Assert(qn_file.Remove(""), nil)
		}

		t.Assert(qn_file.CopyDir(yfolder, tofolder), nil)
		defer delTestFiles(tofolder)

		t.Assert(qn_file.IsDir(yfolder), true)

		for _, v := range haveList {
			t.Assert(qn_file.IsFile(yfolder+"/"+v), true)
		}

		t.Assert(qn_file.IsDir(tofolder), true)

		for _, v := range haveList {
			t.Assert(qn_file.IsFile(tofolder+"/"+v), true)
		}

		t.Assert(qn_file.Remove(tofolder), nil)
		t.Assert(qn_file.Remove(""), nil)
	})
	// Content replacement.
	qn_test.C(t, func(t *qn_test.T) {
		src := qn_file.TempDir(qn_time.TimestampNanoStr(), qn_time.TimestampNanoStr())
		dst := qn_file.TempDir(qn_time.TimestampNanoStr(), qn_time.TimestampNanoStr())
		defer func() {
			qn_file.Remove(src)
			qn_file.Remove(dst)
		}()
		srcContent := "1"
		dstContent := "1"
		t.Assert(qn_file.PutContents(src, srcContent), nil)
		t.Assert(qn_file.PutContents(dst, dstContent), nil)
		t.Assert(qn_file.GetContents(src), srcContent)
		t.Assert(qn_file.GetContents(dst), dstContent)

		err := qn_file.CopyDir(qn_file.Dir(src), qn_file.Dir(dst))
		t.Assert(err, nil)
		t.Assert(qn_file.GetContents(src), srcContent)
		t.Assert(qn_file.GetContents(dst), srcContent)
	})
}
