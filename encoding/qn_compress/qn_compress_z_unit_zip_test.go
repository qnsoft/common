// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_compress_test

import (
	"bytes"
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/encoding/qn_compress"
	"github.com/qnsoft/common/os/gfile"
	"github.com/qnsoft/common/os/gtime"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_ZipPath(t *testing.T) {
	// file
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("zip", "path1", "1.txt")
		dstPath := qn_debug.TestDataPath("zip", "zip.zip")

		t.Assert(gfile.Exists(dstPath), false)
		t.Assert(qn_compress.ZipPath(srcPath, dstPath), nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// unzip to temporary dir.
		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		t.Assert(qn_compress.UnZipFile(dstPath, tempDirPath), nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "1.txt")),
			gfile.GetContents(srcPath),
		)
	})
	// multiple files
	qn_test.C(t, func(t *qn_test.T) {
		var (
			srcPath1 = qn_debug.TestDataPath("zip", "path1", "1.txt")
			srcPath2 = qn_debug.TestDataPath("zip", "path2", "2.txt")
			dstPath  = gfile.TempDir(gtime.TimestampNanoStr(), "zip.zip")
		)
		if p := gfile.Dir(dstPath); !gfile.Exists(p) {
			t.Assert(gfile.Mkdir(p), nil)
		}

		t.Assert(gfile.Exists(dstPath), false)
		err := qn_compress.ZipPath(srcPath1+","+srcPath2, dstPath)
		t.Assert(err, nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// unzip to another temporary dir.
		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		err = qn_compress.UnZipFile(dstPath, tempDirPath)
		t.Assert(err, nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "1.txt")),
			gfile.GetContents(srcPath1),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "2.txt")),
			gfile.GetContents(srcPath2),
		)
	})
	// one dir and one file.
	qn_test.C(t, func(t *qn_test.T) {
		var (
			srcPath1 = qn_debug.TestDataPath("zip", "path1")
			srcPath2 = qn_debug.TestDataPath("zip", "path2", "2.txt")
			dstPath  = gfile.TempDir(gtime.TimestampNanoStr(), "zip.zip")
		)
		if p := gfile.Dir(dstPath); !gfile.Exists(p) {
			t.Assert(gfile.Mkdir(p), nil)
		}

		t.Assert(gfile.Exists(dstPath), false)
		err := qn_compress.ZipPath(srcPath1+","+srcPath2, dstPath)
		t.Assert(err, nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// unzip to another temporary dir.
		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		err = qn_compress.UnZipFile(dstPath, tempDirPath)
		t.Assert(err, nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath1, "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "2.txt")),
			gfile.GetContents(srcPath2),
		)
	})
	// directory.
	qn_test.C(t, func(t *qn_test.T) {
		srcPath := qn_debug.TestDataPath("zip")
		dstPath := qn_debug.TestDataPath("zip", "zip.zip")

		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.Assert(err, nil)

		t.Assert(gfile.Exists(dstPath), false)
		err = qn_compress.ZipPath(srcPath, dstPath)
		t.Assert(err, nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.Assert(err, nil)

		err = qn_compress.UnZipFile(dstPath, tempDirPath)
		t.Assert(err, nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "zip", "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "zip", "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
		)
	})
	// multiple directory paths joined using char ','.
	qn_test.C(t, func(t *qn_test.T) {
		var (
			srcPath  = qn_debug.TestDataPath("zip")
			srcPath1 = qn_debug.TestDataPath("zip", "path1")
			srcPath2 = qn_debug.TestDataPath("zip", "path2")
			dstPath  = qn_debug.TestDataPath("zip", "zip.zip")
		)
		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.Assert(err, nil)

		t.Assert(gfile.Exists(dstPath), false)
		err = qn_compress.ZipPath(srcPath1+", "+srcPath2, dstPath)
		t.Assert(err, nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.Assert(err, nil)

		zipContent := gfile.GetBytes(dstPath)
		t.AssertGT(len(zipContent), 0)
		err = qn_compress.UnZipContent(zipContent, tempDirPath)
		t.Assert(err, nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathWriter(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			srcPath  = qn_debug.TestDataPath("zip")
			srcPath1 = qn_debug.TestDataPath("zip", "path1")
			srcPath2 = qn_debug.TestDataPath("zip", "path2")
		)
		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.Assert(err, nil)

		writer := bytes.NewBuffer(nil)
		t.Assert(writer.Len(), 0)
		err = qn_compress.ZipPathWriter(srcPath1+", "+srcPath2, writer)
		t.Assert(err, nil)
		t.AssertGT(writer.Len(), 0)

		tempDirPath := gfile.TempDir(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.Assert(err, nil)

		zipContent := writer.Bytes()
		t.AssertGT(len(zipContent), 0)
		err = qn_compress.UnZipContent(zipContent, tempDirPath)
		t.Assert(err, nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
		)
	})
}
