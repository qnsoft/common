// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/qn_ins"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Database(t *testing.T) {
	databaseContent := qn_file.GetContents(
		qn_debug.TestDataPath("database", "config.toml"),
	)
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		dirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(dirPath)
		t.Assert(err, nil)
		defer qn_file.Remove(dirPath)

		name := "config.toml"
		err = qn_file.PutContents(qn_file.Join(dirPath, name), databaseContent)
		t.Assert(err, nil)

		err = qn_ins.Config().AddPath(dirPath)
		t.Assert(err, nil)

		defer qn_ins.Config().Clear()

		// for gfsnotify callbacks to refresh cache of config file
		time.Sleep(500 * time.Millisecond)

		//fmt.Println("qn_ins Test_Database", Config().Get("test"))

		dbDefault := qn_ins.Database()
		dbTest := qn_ins.Database("test")
		t.AssertNE(dbDefault, nil)
		t.AssertNE(dbTest, nil)

		t.Assert(dbDefault.PingMaster(), nil)
		t.Assert(dbDefault.PingSlave(), nil)
		t.Assert(dbTest.PingMaster(), nil)
		t.Assert(dbTest.PingSlave(), nil)
	})
}
