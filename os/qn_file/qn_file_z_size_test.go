// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"testing"

	qn_conv "github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Size(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			paths1 string = "/testfile_t1.txt"
			sizes  int64
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = qn_file.Size(testpath() + paths1)
		t.Assert(sizes, 14)

		sizes = qn_file.Size("")
		t.Assert(sizes, 0)

	})
}

func Test_StrToSize(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_file.StrToSize("0.00B"), 0)
		t.Assert(qn_file.StrToSize("16.00B"), 16)
		t.Assert(qn_file.StrToSize("1.00K"), 1024)
		t.Assert(qn_file.StrToSize("1.00KB"), 1024)
		t.Assert(qn_file.StrToSize("1.00KiloByte"), 1024)
		t.Assert(qn_file.StrToSize("15.26M"), qn_conv.Int64(15.26*1024*1024))
		t.Assert(qn_file.StrToSize("15.26MB"), qn_conv.Int64(15.26*1024*1024))
		t.Assert(qn_file.StrToSize("1.49G"), qn_conv.Int64(1.49*1024*1024*1024))
		t.Assert(qn_file.StrToSize("1.49GB"), qn_conv.Int64(1.49*1024*1024*1024))
		t.Assert(qn_file.StrToSize("8.73T"), qn_conv.Int64(8.73*1024*1024*1024*1024))
		t.Assert(qn_file.StrToSize("8.73TB"), qn_conv.Int64(8.73*1024*1024*1024*1024))
		t.Assert(qn_file.StrToSize("8.53P"), qn_conv.Int64(8.53*1024*1024*1024*1024*1024))
		t.Assert(qn_file.StrToSize("8.53PB"), qn_conv.Int64(8.53*1024*1024*1024*1024*1024))
		t.Assert(qn_file.StrToSize("8.01EB"), qn_conv.Int64(8.01*1024*1024*1024*1024*1024*1024))
	})
}

func Test_FormatSize(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_file.FormatSize(0), "0.00B")
		t.Assert(qn_file.FormatSize(16), "16.00B")

		t.Assert(qn_file.FormatSize(1024), "1.00K")

		t.Assert(qn_file.FormatSize(16000000), "15.26M")

		t.Assert(qn_file.FormatSize(1600000000), "1.49G")

		t.Assert(qn_file.FormatSize(9600000000000), "8.73T")
		t.Assert(qn_file.FormatSize(9600000000000000), "8.53P")
	})
}

func Test_ReadableSize(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {

		var (
			paths1 string = "/testfile_t1.txt"
		)
		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)
		t.Assert(qn_file.ReadableSize(testpath()+paths1), "14.00B")
		t.Assert(qn_file.ReadableSize(""), "0.00B")

	})
}
