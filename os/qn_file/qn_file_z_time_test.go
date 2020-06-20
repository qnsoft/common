// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"os"
	"testing"
	"time"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_MTime(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {

		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.Assert(err, nil)

		t.Assert(qn_file.MTime(testpath()+file1), fileobj.ModTime())
		t.Assert(qn_file.MTime(""), "")
	})
}

func Test_MTimeMillisecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.Assert(err, nil)

		time.Sleep(time.Millisecond * 100)
		t.AssertGE(
			qn_file.MTimestampMilli(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(qn_file.MTimestampMilli(""), -1)
	})
}
