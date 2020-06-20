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

func Test_NotFound(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		teatFile := qn_file.Dir(qn_debug.CallerFilePath()) + qn_file.Separator + "testdata/readline/error.log"
		callback := func(line string) {
		}
		err := qn_file.ReadLines(teatFile, callback)
		t.AssertNE(err, nil)
	})
}

func Test_ReadLines(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expectList := []string{"a", "b", "c", "d", "e"}

		getList := make([]string, 0)
		callback := func(line string) {
			getList = append(getList, line)
		}

		teatFile := qn_file.Dir(qn_debug.CallerFilePath()) + qn_file.Separator + "testdata/readline/file.log"
		err := qn_file.ReadLines(teatFile, callback)

		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}

func Test_ReadByteLines(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expectList := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e")}

		getList := make([][]byte, 0)
		callback := func(line []byte) {
			getList = append(getList, line)
		}

		teatFile := qn_file.Dir(qn_debug.CallerFilePath()) + qn_file.Separator + "testdata/readline/file.log"
		err := qn_file.ReadByteLines(teatFile, callback)

		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}
