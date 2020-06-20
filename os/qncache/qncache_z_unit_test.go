// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qncache_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/qnsoft/common/os/gfcache"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func TestGetContents(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {

		var f *os.File
		var err error
		fileName := "test"
		strTest := "123"

		if !qn_file.Exists(fileName) {
			f, err = ioutil.TempFile("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if qn_file.Exists(f.Name()) {

			f, err = qn_file.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = qn_file.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			cache := gfcache.GetContents(f.Name(), 1)
			t.Assert(cache, strTest)
		}
	})

	qn_test.C(t, func(t *qn_test.T) {

		var f *os.File
		var err error
		fileName := "test2"
		strTest := "123"

		if !qn_file.Exists(fileName) {
			f, err = ioutil.TempFile("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if qn_file.Exists(f.Name()) {
			cache := gfcache.GetContents(f.Name())

			f, err = qn_file.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = qn_file.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			t.Assert(cache, "")

			time.Sleep(100 * time.Millisecond)
		}
	})
}
