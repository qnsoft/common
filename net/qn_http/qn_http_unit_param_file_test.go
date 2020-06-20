// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/text/gstr"
	"github.com/qnsoft/common/debug/gdebug"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/ghttp"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Params_File_Single(t *testing.T) {
	dstDirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/single", func(r *ghttp.Request) {
		file := r.GetUploadFile("file")
		if file == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}

		if name, err := file.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(name)
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	// normal name
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file1.txt")
		dstPath := qn_file.Join(dstDirPath, "file1.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt")
		t.Assert(qn_file.GetContents(dstPath), qn_file.GetContents(srcPath))
	})
	// randomly rename.
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file2.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file":           "@file:" + srcPath,
			"randomlyRename": true,
		})
		dstPath := qn_file.Join(dstDirPath, content)
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(qn_file.GetContents(dstPath), qn_file.GetContents(srcPath))
	})
}

func Test_Params_File_CustomName(t *testing.T) {
	dstDirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/single", func(r *ghttp.Request) {
		file := r.GetUploadFile("file")
		if file == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}
		file.Filename = "my.txt"
		if name, err := file.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(name)
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath := gdebug.TestDataPath("upload", "file1.txt")
		dstPath := qn_file.Join(dstDirPath, "my.txt")
		content := client.PostContent("/upload/single", g.Map{
			"file": "@file:" + srcPath,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "my.txt")
		t.Assert(qn_file.GetContents(dstPath), qn_file.GetContents(srcPath))
	})
}

func Test_Params_File_Batch(t *testing.T) {
	dstDirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/upload/batch", func(r *ghttp.Request) {
		files := r.GetUploadFiles("file")
		if files == nil {
			r.Response.WriteExit("upload file cannot be empty")
		}
		if names, err := files.Save(dstDirPath, r.GetBool("randomlyRename")); err == nil {
			r.Response.WriteExit(gstr.Join(names, ","))
		}
		r.Response.WriteExit("upload failed")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	// normal name
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath1 := gdebug.TestDataPath("upload", "file1.txt")
		srcPath2 := gdebug.TestDataPath("upload", "file2.txt")
		dstPath1 := qn_file.Join(dstDirPath, "file1.txt")
		dstPath2 := qn_file.Join(dstDirPath, "file2.txt")
		content := client.PostContent("/upload/batch", g.Map{
			"file[0]": "@file:" + srcPath1,
			"file[1]": "@file:" + srcPath2,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")
		t.Assert(content, "file1.txt,file2.txt")
		t.Assert(qn_file.GetContents(dstPath1), qn_file.GetContents(srcPath1))
		t.Assert(qn_file.GetContents(dstPath2), qn_file.GetContents(srcPath2))
	})
	// randomly rename.
	qn_test.C(t, func(t *qn_test.T) {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		srcPath1 := gdebug.TestDataPath("upload", "file1.txt")
		srcPath2 := gdebug.TestDataPath("upload", "file2.txt")
		content := client.PostContent("/upload/batch", g.Map{
			"file[0]":        "@file:" + srcPath1,
			"file[1]":        "@file:" + srcPath2,
			"randomlyRename": true,
		})
		t.AssertNE(content, "")
		t.AssertNE(content, "upload file cannot be empty")
		t.AssertNE(content, "upload failed")

		array := gstr.SplitAndTrim(content, ",")
		t.Assert(len(array), 2)
		dstPath1 := qn_file.Join(dstDirPath, array[0])
		dstPath2 := qn_file.Join(dstDirPath, array[1])
		t.Assert(qn_file.GetContents(dstPath1), qn_file.GetContents(srcPath1))
		t.Assert(qn_file.GetContents(dstPath2), qn_file.GetContents(srcPath2))
	})
}
