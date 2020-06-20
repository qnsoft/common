// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_file_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func createTestFile(filename, content string) error {
	TempDir := testpath()
	err := ioutil.WriteFile(TempDir+filename, []byte(content), 0666)
	return err
}

func delTestFiles(filenames string) {
	os.RemoveAll(testpath() + filenames)
}

func createDir(paths string) {
	TempDir := testpath()
	os.Mkdir(TempDir+paths, 0777)
}

func formatpaths(paths []string) []string {
	for k, v := range paths {
		paths[k] = filepath.ToSlash(v)
		paths[k] = strings.Replace(paths[k], "./", "/", 1)
	}

	return paths
}

func formatpath(paths string) string {
	paths = filepath.ToSlash(paths)
	paths = strings.Replace(paths, "./", "/", 1)
	return paths
}

func testpath() string {
	return qn.str.TrimRight(os.TempDir(), "\\/")
}

func Test_GetContents(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {

		var (
			filepaths string = "/testfile_t1.txt"
		)
		createTestFile(filepaths, "my name is jroam")
		defer delTestFiles(filepaths)

		t.Assert(qn_file.GetContents(testpath()+filepaths), "my name is jroam")
		t.Assert(qn_file.GetContents(""), "")

	})
}

func Test_GetBinContents(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths1  = "/testfile_t1.txt"
			filepaths2  = testpath() + "/testfile_t1_no.txt"
			readcontent []byte
			str1        = "my name is jroam"
		)
		createTestFile(filepaths1, str1)
		defer delTestFiles(filepaths1)
		readcontent = qn_file.GetBytes(testpath() + filepaths1)
		t.Assert(readcontent, []byte(str1))

		readcontent = qn_file.GetBytes(filepaths2)
		t.Assert(string(readcontent), "")

		t.Assert(string(qn_file.GetBytes(filepaths2)), "")

	})
}

func Test_Truncate(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths1 = "/testfile_GetContentsyyui.txt"
			err        error
			files      *os.File
		)
		createTestFile(filepaths1, "abcdefghijkmln")
		defer delTestFiles(filepaths1)
		err = qn_file.Truncate(testpath()+filepaths1, 10)
		t.Assert(err, nil)

		files, err = os.Open(testpath() + filepaths1)
		defer files.Close()
		t.Assert(err, nil)
		fileinfo, err2 := files.Stat()
		t.Assert(err2, nil)
		t.Assert(fileinfo.Size(), 10)

		err = qn_file.Truncate("", 10)
		t.AssertNE(err, nil)

	})
}

func Test_PutContents(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = qn_file.PutContents(testpath()+filepaths, "test!")
		t.Assert(err, nil)

		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		t.Assert(err, nil)
		t.Assert(string(readcontent), "test!")

		err = qn_file.PutContents("", "test!")
		t.AssertNE(err, nil)

	})
}

func Test_PutContentsAppend(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)

		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)
		err = qn_file.PutContentsAppend(testpath()+filepaths, "hello")
		t.Assert(err, nil)

		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		t.Assert(err, nil)
		t.Assert(string(readcontent), "ahello")

		err = qn_file.PutContentsAppend("", "hello")
		t.AssertNE(err, nil)

	})

}

func Test_PutBinContents(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = qn_file.PutBytes(testpath()+filepaths, []byte("test!!"))
		t.Assert(err, nil)

		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		t.Assert(err, nil)
		t.Assert(string(readcontent), "test!!")

		err = qn_file.PutBytes("", []byte("test!!"))
		t.AssertNE(err, nil)

	})
}

func Test_PutBinContentsAppend(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "test!!")
		defer delTestFiles(filepaths)
		err = qn_file.PutBytesAppend(testpath()+filepaths, []byte("word"))
		t.Assert(err, nil)

		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		t.Assert(err, nil)
		t.Assert(string(readcontent), "test!!word")

		err = qn_file.PutBytesAppend("", []byte("word"))
		t.AssertNE(err, nil)

	})
}

func Test_GetBinContentsByTwoOffsetsByPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths   = "/testfile_GetContents.txt"
			readcontent []byte
		)

		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		readcontent = qn_file.GetBytesByTwoOffsetsByPath(testpath()+filepaths, 2, 5)

		t.Assert(string(readcontent), "cde")

		readcontent = qn_file.GetBytesByTwoOffsetsByPath("", 2, 5)
		t.Assert(len(readcontent), 0)

	})

}

func Test_GetNextCharOffsetByPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			filepaths  = "/testfile_GetContents.txt"
			localindex int64
		)
		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		localindex = qn_file.GetNextCharOffsetByPath(testpath()+filepaths, 'd', 1)
		t.Assert(localindex, 3)

		localindex = qn_file.GetNextCharOffsetByPath("", 'd', 1)
		t.Assert(localindex, -1)

	})
}

func Test_GetNextCharOffset(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			localindex int64
		)
		reader := strings.NewReader("helloword")

		localindex = qn_file.GetNextCharOffset(reader, 'w', 1)
		t.Assert(localindex, 5)

		localindex = qn_file.GetNextCharOffset(reader, 'j', 1)
		t.Assert(localindex, -1)

	})
}

func Test_GetBinContentsByTwoOffsets(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			reads []byte
		)
		reader := strings.NewReader("helloword")

		reads = qn_file.GetBytesByTwoOffsets(reader, 1, 3)
		t.Assert(string(reads), "el")

		reads = qn_file.GetBytesByTwoOffsets(reader, 10, 30)
		t.Assert(string(reads), "")

	})
}

func Test_GetBinContentsTilChar(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			reads  []byte
			indexs int64
		)
		reader := strings.NewReader("helloword")

		reads, _ = qn_file.GetBytesTilChar(reader, 'w', 2)
		t.Assert(string(reads), "llow")

		_, indexs = qn_file.GetBytesTilChar(reader, 'w', 20)
		t.Assert(indexs, -1)

	})
}

func Test_GetBinContentsTilCharByPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			reads     []byte
			indexs    int64
			filepaths = "/testfile_GetContents.txt"
		)

		createTestFile(filepaths, "abcdefghijklmn")
		defer delTestFiles(filepaths)

		reads, _ = qn_file.GetBytesTilCharByPath(testpath()+filepaths, 'c', 2)
		t.Assert(string(reads), "c")

		reads, _ = qn_file.GetBytesTilCharByPath(testpath()+filepaths, 'y', 1)
		t.Assert(string(reads), "")

		_, indexs = qn_file.GetBytesTilCharByPath(testpath()+filepaths, 'x', 1)
		t.Assert(indexs, -1)

	})
}

func Test_Home(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			reads string
			err   error
		)

		reads, err = qn_file.Home()
		t.Assert(err, nil)
		t.AssertNE(reads, "")
	})
}
