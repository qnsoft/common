// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_compress_test

import (
	"testing"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/os/gfile"
	"github.com/qnsoft/common/os/gtime"

	"github.com/qnsoft/common/encoding/qn_compress"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Gzip_UnGzip(t *testing.T) {
	src := "Hello World!!"

	gzip := []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xff,
		0xf2, 0x48, 0xcd, 0xc9, 0xc9,
		0x57, 0x08, 0xcf, 0x2f, 0xca,
		0x49, 0x51, 0x54, 0x04, 0x04,
		0x00, 0x00, 0xff, 0xff, 0x9d,
		0x24, 0xa8, 0xd1, 0x0d, 0x00,
		0x00, 0x00,
	}
	qn_test.C(t, func(t *qn_test.T) {
		arr := []byte(src)
		data, _ := qn_compress.Gzip(arr)
		t.Assert(data, gzip)

		data, _ = qn_compress.UnGzip(gzip)
		t.Assert(data, arr)

		data, _ = qn_compress.UnGzip(gzip[1:])
		t.Assert(data, nil)
	})
}

func Test_Gzip_UnGzip_File(t *testing.T) {
	srcPath := qn_debug.TestDataPath("gzip", "file.txt")
	dstPath1 := gfile.TempDir(gtime.TimestampNanoStr(), "gzip.zip")
	dstPath2 := gfile.TempDir(gtime.TimestampNanoStr(), "file.txt")

	// Compress.
	qn_test.C(t, func(t *qn_test.T) {
		err := qn_compress.GzipFile(srcPath, dstPath1, 9)
		t.Assert(err, nil)
		defer gfile.Remove(dstPath1)
		t.Assert(gfile.Exists(dstPath1), true)

		// Decompress.
		err = qn_compress.UnGzipFile(dstPath1, dstPath2)
		t.Assert(err, nil)
		defer gfile.Remove(dstPath2)
		t.Assert(gfile.Exists(dstPath2), true)

		t.Assert(gfile.GetContents(srcPath), gfile.GetContents(dstPath2))
	})
}
