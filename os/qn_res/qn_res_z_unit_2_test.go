// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_res_test

import (
	_ "github.com/qnsoft/common/os/qn_res/testdata/data"

	"testing"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/test/qn_test"

	"github.com/qnsoft/common/os/qn_res"
)

func Test_Basic(t *testing.T) {
	qn_res.Dump()
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_res.Get("none"), nil)
		t.Assert(qn_res.Contains("none"), false)
		t.Assert(qn_res.Contains("dir1"), true)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := "dir1/test1"
		file := qn_res.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)

		info := file.FileInfo()
		t.AssertNE(info, nil)
		t.Assert(info.IsDir(), false)
		t.Assert(info.Name(), "test1")

		rc, err := file.Open()
		t.Assert(err, nil)
		defer rc.Close()

		b := make([]byte, 5)
		n, err := rc.Read(b)
		t.Assert(n, 5)
		t.Assert(err, nil)
		t.Assert(string(b), "test1")

		t.Assert(file.Content(), "test1 content")
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := "dir2"
		file := qn_res.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)

		info := file.FileInfo()
		t.AssertNE(info, nil)
		t.Assert(info.IsDir(), true)
		t.Assert(info.Name(), "dir2")

		rc, err := file.Open()
		t.Assert(err, nil)
		defer rc.Close()

		t.Assert(file.Content(), nil)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := "dir2/test2"
		file := qn_res.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)
		t.Assert(file.Content(), "test2 content")
	})
}

func Test_Get(t *testing.T) {
	qn_res.Dump()
	qn_test.C(t, func(t *qn_test.T) {
		t.AssertNE(qn_res.Get("dir1/test1"), nil)
	})
	qn_test.C(t, func(t *qn_test.T) {
		file := qn_res.GetWithIndex("dir1", g.SliceStr{"test1"})
		t.AssertNE(file, nil)
		t.Assert(file.Name(), "dir1/test1")
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_res.GetContent("dir1"), "")
		t.Assert(qn_res.GetContent("dir1/test1"), "test1 content")
	})
}

func Test_ScanDir(t *testing.T) {
	qn_res.Dump()
	qn_test.C(t, func(t *qn_test.T) {
		path := "dir1"
		files := qn_res.ScanDir(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := "dir1"
		files := qn_res.ScanDir(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 3)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := "dir1"
		files := qn_res.ScanDir(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir1/sub/sub-test1.txt")
		t.Assert(files[0].Content(), "sub-test1 content")
	})
}

func Test_ScanDirFile(t *testing.T) {
	qn_res.Dump()
	qn_test.C(t, func(t *qn_test.T) {
		path := "dir2"
		files := qn_res.ScanDirFile(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := "dir2"
		files := qn_res.ScanDirFile(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := "dir2"
		files := qn_res.ScanDirFile(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir2/sub/sub-test2.txt")
		t.Assert(files[0].Content(), "sub-test2 content")
	})
}
