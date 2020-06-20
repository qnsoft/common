// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/qn_ins"

	"github.com/qnsoft/common/os/qn_cfg"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

var (
	configContent = qn_file.GetContents(
		qn_debug.TestDataPath("config", "config.toml"),
	)
)

func Test_Config1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.AssertNE(configContent, "")
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.AssertNE(qn_ins.Config(), nil)
	})
}

func Test_Config2(t *testing.T) {
	// relative path
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		dirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(dirPath)
		t.Assert(err, nil)
		defer qn_file.Remove(dirPath)

		name := "config.toml"
		err = qn_file.PutContents(qn_file.Join(dirPath, name), configContent)
		t.Assert(err, nil)

		err = qn_ins.Config().AddPath(dirPath)
		t.Assert(err, nil)

		defer qn_ins.Config().Clear()

		t.Assert(qn_ins.Config().Get("test"), "v=1")
		t.Assert(qn_ins.Config().Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for qn_snotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	// relative path, config folder
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		dirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(dirPath)
		t.Assert(err, nil)
		defer qn_file.Remove(dirPath)

		name := "config/config.toml"
		err = qn_file.PutContents(qn_file.Join(dirPath, name), configContent)
		t.Assert(err, nil)

		err = qn_ins.Config().AddPath(dirPath)
		t.Assert(err, nil)

		defer qn_ins.Config().Clear()

		t.Assert(qn_ins.Config().Get("test"), "v=1")
		t.Assert(qn_ins.Config().Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config().Get("redis.disk"), "127.0.0.1:6379,0")

		// for qn_snotify callbacks to refresh cache of config file
		time.Sleep(500 * time.Millisecond)
	})
}

func Test_Config3(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		dirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(dirPath)
		t.Assert(err, nil)
		defer qn_file.Remove(dirPath)

		name := "test.toml"
		err = qn_file.PutContents(qn_file.Join(dirPath, name), configContent)
		t.Assert(err, nil)

		err = qn_ins.Config("test").AddPath(dirPath)
		t.Assert(err, nil)

		defer qn_ins.Config("test").Clear()
		qn_ins.Config("test").SetFileName("test.toml")

		t.Assert(qn_ins.Config("test").Get("test"), "v=1")
		t.Assert(qn_ins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for qn_snotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		var err error
		dirPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(dirPath)
		t.Assert(err, nil)
		defer qn_file.Remove(dirPath)

		name := "config/test.toml"
		err = qn_file.PutContents(qn_file.Join(dirPath, name), configContent)
		t.Assert(err, nil)

		err = qn_ins.Config("test").AddPath(dirPath)
		t.Assert(err, nil)

		defer qn_ins.Config("test").Clear()
		qn_ins.Config("test").SetFileName("test.toml")

		t.Assert(qn_ins.Config("test").Get("test"), "v=1")
		t.Assert(qn_ins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for qn_snotify callbacks to refresh cache of config file for next unit testing case.
	time.Sleep(500 * time.Millisecond)
}

func Test_Config4(t *testing.T) {
	// absolute path
	qn_test.C(t, func(t *qn_test.T) {
		path := fmt.Sprintf(`%s/%d`, qn_file.TempDir(), qn_time.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := qn_file.PutContents(file, configContent)
		t.Assert(err, nil)
		defer qn_file.Remove(file)
		defer qn_ins.Config().Clear()

		t.Assert(qn_ins.Config().AddPath(path), nil)
		t.Assert(qn_ins.Config().Get("test"), "v=1")
		t.Assert(qn_ins.Config().Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		path := fmt.Sprintf(`%s/%d/config`, qn_file.TempDir(), qn_time.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := qn_file.PutContents(file, configContent)
		t.Assert(err, nil)
		defer qn_file.Remove(file)
		defer qn_ins.Config().Clear()
		t.Assert(qn_ins.Config().AddPath(path), nil)
		t.Assert(qn_ins.Config().Get("test"), "v=1")
		t.Assert(qn_ins.Config().Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		path := fmt.Sprintf(`%s/%d`, qn_file.TempDir(), qn_time.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := qn_file.PutContents(file, configContent)
		t.Assert(err, nil)
		defer qn_file.Remove(file)
		defer qn_ins.Config("test").Clear()
		qn_ins.Config("test").SetFileName("test.toml")
		t.Assert(qn_ins.Config("test").AddPath(path), nil)
		t.Assert(qn_ins.Config("test").Get("test"), "v=1")
		t.Assert(qn_ins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		path := fmt.Sprintf(`%s/%d/config`, qn_file.TempDir(), qn_time.TimestampNano())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := qn_file.PutContents(file, configContent)
		t.Assert(err, nil)
		defer qn_file.Remove(file)
		defer qn_ins.Config().Clear()
		qn_ins.Config("test").SetFileName("test.toml")
		t.Assert(qn_ins.Config("test").AddPath(path), nil)
		t.Assert(qn_ins.Config("test").Get("test"), "v=1")
		t.Assert(qn_ins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		t.Assert(qn_ins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
}
func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	qn_test.C(t, func(t *qn_test.T) {
		path := qn_cfg.DEFAULT_CONFIG_FILE
		err := qn_file.PutContents(path, config)
		t.Assert(err, nil)
		defer func() {
			_ = qn_file.Remove(path)
		}()

		t.Assert(qn_ins.Config().Get("log-path"), "logs")
	})
}
