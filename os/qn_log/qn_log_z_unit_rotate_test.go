// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_log_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_log"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/gstr"
)

func Test_Rotate_Size(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		l := qn_log.New()
		p := qn_file.TempDir(qn_time.TimestampNanoStr())
		err := l.SetConfigWithMap(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateSize":           10,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // For unit testing only.
		})
		t.Assert(err, nil)
		defer qn_file.Remove(p)

		s := "1234567890abcdefg"
		for i := 0; i < 10; i++ {
			l.Print(s)
		}

		time.Sleep(time.Second * 3)

		files, err := qn_file.ScanDirFile(p, "*.gz")
		t.Assert(err, nil)
		t.Assert(len(files), 2)

		content := qn_file.GetContents(qn_file.Join(p, "access.log"))
		t.Assert(gstr.Count(content, s), 1)

		time.Sleep(time.Second * 5)
		files, err = qn_file.ScanDirFile(p, "*.gz")
		t.Assert(err, nil)
		t.Assert(len(files), 0)
	})
}

func Test_Rotate_Expire(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		l := qn_log.New()
		p := qn_file.TempDir(qn_time.TimestampNanoStr())
		err := l.SetConfigWithMap(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateExpire":         time.Second,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // For unit testing only.
		})
		t.Assert(err, nil)
		defer qn_file.Remove(p)

		s := "1234567890abcdefg"
		for i := 0; i < 10; i++ {
			l.Print(s)
		}

		files, err := qn_file.ScanDirFile(p, "*.gz")
		t.Assert(err, nil)
		t.Assert(len(files), 0)

		t.Assert(gstr.Count(qn_file.GetContents(qn_file.Join(p, "access.log")), s), 10)

		time.Sleep(time.Second * 3)

		files, err = qn_file.ScanDirFile(p, "*.gz")
		t.Assert(err, nil)
		t.Assert(len(files), 1)

		t.Assert(gstr.Count(qn_file.GetContents(qn_file.Join(p, "access.log")), s), 0)

		time.Sleep(time.Second * 5)
		files, err = qn_file.ScanDirFile(p, "*.gz")
		t.Assert(err, nil)
		t.Assert(len(files), 0)
	})
}
