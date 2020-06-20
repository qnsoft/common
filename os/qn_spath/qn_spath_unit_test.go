// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_spath_test

import (
	"testing"

	"github.com/qnsoft/common/os/gspath"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func TestSPath_Api(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		pwd := qn_file.Pwd()
		root := pwd
		qn_file.Create(qn_file.Join(root, "gf_tmp", "gf.txt"))
		defer qn_file.Remove(qn_file.Join(root, "gf_tmp"))
		fp, isDir := gspath.Search(root, "gf_tmp")
		t.Assert(fp, qn_file.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gspath.Search(root, "gf_tmp", "gf.txt")
		t.Assert(fp, qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)

		fp, isDir = gspath.SearchWithCache(root, "gf_tmp")
		t.Assert(fp, qn_file.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gspath.SearchWithCache(root, "gf_tmp", "gf.txt")
		t.Assert(fp, qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
	})
}

func TestSPath_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		pwd := qn_file.Pwd()
		root := pwd

		qn_file.Create(qn_file.Join(root, "gf_tmp", "gf.txt"))
		defer qn_file.Remove(qn_file.Join(root, "gf_tmp"))
		gsp := gspath.New(root, false)
		realPath, err := gsp.Add(qn_file.Join(root, "gf_tmp"))
		t.Assert(err, nil)
		t.Assert(realPath, qn_file.Join(root, "gf_tmp"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.Add(qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		gsp.Remove("gf_tmp1")
		t.Assert(gsp.Size(), 2)
		t.Assert(len(gsp.Paths()), 2)
		t.Assert(len(gsp.AllPaths()), 0)
		realPath, err = gsp.Set(qn_file.Join(root, "gf_tmp1"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.Set(qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.AssertNE(err, nil)
		t.Assert(realPath, "")

		realPath, err = gsp.Set(root)
		t.Assert(err, nil)
		t.Assert(realPath, root)

		fp, isDir := gsp.Search("gf_tmp")
		t.Assert(fp, qn_file.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
		fp, isDir = gsp.Search("/", "gf.txt")
		t.Assert(fp, root)
		t.Assert(isDir, true)

		gsp = gspath.New(root, true)
		realPath, err = gsp.Add(qn_file.Join(root, "gf_tmp"))
		t.Assert(err, nil)
		t.Assert(realPath, qn_file.Join(root, "gf_tmp"))

		qn_file.Mkdir(qn_file.Join(root, "gf_tmp1"))
		qn_file.Rename(qn_file.Join(root, "gf_tmp1"), qn_file.Join(root, "gf_tmp2"))
		qn_file.Rename(qn_file.Join(root, "gf_tmp2"), qn_file.Join(root, "gf_tmp1"))
		defer qn_file.Remove(qn_file.Join(root, "gf_tmp1"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, false)
		t.Assert(realPath, qn_file.Join(root, "gf_tmp1"))
		realPath, err = gsp.Add("gf_tmp3")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		gsp.Remove(qn_file.Join(root, "gf_tmp"))
		gsp.Remove(qn_file.Join(root, "gf_tmp1"))
		gsp.Remove(qn_file.Join(root, "gf_tmp3"))
		t.Assert(gsp.Size(), 3)
		t.Assert(len(gsp.Paths()), 3)
		gsp.AllPaths()
		gsp.Set(root)
		fp, isDir = gsp.Search("gf_tmp")
		t.Assert(fp, qn_file.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, qn_file.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
		fp, isDir = gsp.Search("/", "gf.txt")
		t.Assert(fp, pwd)
		t.Assert(isDir, true)
	})
}
