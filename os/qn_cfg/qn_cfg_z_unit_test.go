// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qn_cfg_test

import (
	"os"
	"testing"

	"github.com/qnsoft/common/debug/gdebug"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/gcfg"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func init() {
	os.Setenv("GF_GCFG_ERRORPRINT", "false")
}

func Test_Basic1(t *testing.T) {
	config := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	qn_test.C(t, func(t *qn_test.T) {
		path := gcfg.DEFAULT_CONFIG_FILE
		err := qn_file.PutContents(path, config)
		t.Assert(err, nil)
		defer qn_file.Remove(path)

		c := gcfg.New()
		t.Assert(c.Get("v1"), 1)
		t.AssertEQ(c.GetInt("v1"), 1)
		t.AssertEQ(c.GetInt8("v1"), int8(1))
		t.AssertEQ(c.GetInt16("v1"), int16(1))
		t.AssertEQ(c.GetInt32("v1"), int32(1))
		t.AssertEQ(c.GetInt64("v1"), int64(1))
		t.AssertEQ(c.GetUint("v1"), uint(1))
		t.AssertEQ(c.GetUint8("v1"), uint8(1))
		t.AssertEQ(c.GetUint16("v1"), uint16(1))
		t.AssertEQ(c.GetUint32("v1"), uint32(1))
		t.AssertEQ(c.GetUint64("v1"), uint64(1))

		t.AssertEQ(c.GetVar("v1").String(), "1")
		t.AssertEQ(c.GetVar("v1").Bool(), true)
		t.AssertEQ(c.GetVar("v2").String(), "true")
		t.AssertEQ(c.GetVar("v2").Bool(), true)

		t.AssertEQ(c.GetString("v1"), "1")
		t.AssertEQ(c.GetFloat32("v4"), float32(1.23))
		t.AssertEQ(c.GetFloat64("v4"), float64(1.23))
		t.AssertEQ(c.GetString("v2"), "true")
		t.AssertEQ(c.GetBool("v2"), true)
		t.AssertEQ(c.GetBool("v3"), false)

		t.AssertEQ(c.Contains("v1"), true)
		t.AssertEQ(c.Contains("v2"), true)
		t.AssertEQ(c.Contains("v3"), true)
		t.AssertEQ(c.Contains("v4"), true)
		t.AssertEQ(c.Contains("v5"), false)

		t.AssertEQ(c.GetInts("array"), []int{1, 2, 3})
		t.AssertEQ(c.GetStrings("array"), []string{"1", "2", "3"})
		t.AssertEQ(c.GetArray("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetInterfaces("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetMap("redis"), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		t.AssertEQ(c.FilePath(), qn_file.Pwd()+qn_file.Separator+path)

	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	qn_test.C(t, func(t *qn_test.T) {
		path := gcfg.DEFAULT_CONFIG_FILE
		err := qn_file.PutContents(path, config)
		t.Assert(err, nil)
		defer func() {
			_ = qn_file.Remove(path)
		}()

		c := gcfg.New()
		t.Assert(c.Get("log-path"), "logs")
	})
}

func Test_Content(t *testing.T) {
	content := `
v1    = 1
v2    = "true"
v3    = "off"
v4    = "1.23"
array = [1,2,3]
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	gcfg.SetContent(content)
	defer gcfg.ClearContent()

	qn_test.C(t, func(t *qn_test.T) {
		c := gcfg.New()
		t.Assert(c.Get("v1"), 1)
		t.AssertEQ(c.GetInt("v1"), 1)
		t.AssertEQ(c.GetInt8("v1"), int8(1))
		t.AssertEQ(c.GetInt16("v1"), int16(1))
		t.AssertEQ(c.GetInt32("v1"), int32(1))
		t.AssertEQ(c.GetInt64("v1"), int64(1))
		t.AssertEQ(c.GetUint("v1"), uint(1))
		t.AssertEQ(c.GetUint8("v1"), uint8(1))
		t.AssertEQ(c.GetUint16("v1"), uint16(1))
		t.AssertEQ(c.GetUint32("v1"), uint32(1))
		t.AssertEQ(c.GetUint64("v1"), uint64(1))

		t.AssertEQ(c.GetVar("v1").String(), "1")
		t.AssertEQ(c.GetVar("v1").Bool(), true)
		t.AssertEQ(c.GetVar("v2").String(), "true")
		t.AssertEQ(c.GetVar("v2").Bool(), true)

		t.AssertEQ(c.GetString("v1"), "1")
		t.AssertEQ(c.GetFloat32("v4"), float32(1.23))
		t.AssertEQ(c.GetFloat64("v4"), float64(1.23))
		t.AssertEQ(c.GetString("v2"), "true")
		t.AssertEQ(c.GetBool("v2"), true)
		t.AssertEQ(c.GetBool("v3"), false)

		t.AssertEQ(c.Contains("v1"), true)
		t.AssertEQ(c.Contains("v2"), true)
		t.AssertEQ(c.Contains("v3"), true)
		t.AssertEQ(c.Contains("v4"), true)
		t.AssertEQ(c.Contains("v5"), false)

		t.AssertEQ(c.GetInts("array"), []int{1, 2, 3})
		t.AssertEQ(c.GetStrings("array"), []string{"1", "2", "3"})
		t.AssertEQ(c.GetArray("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetInterfaces("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetMap("redis"), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
	})
}

func Test_SetFileName(t *testing.T) {
	config := `
{
	"array": [
		1,
		2,
		3
	],
	"redis": {
		"cache": "127.0.0.1:6379,1",
		"disk": "127.0.0.1:6379,0"
	},
	"v1": 1,
	"v2": "true",
	"v3": "off",
	"v4": "1.234"
}
`
	qn_test.C(t, func(t *qn_test.T) {
		path := "config.json"
		err := qn_file.PutContents(path, config)
		t.Assert(err, nil)
		defer func() {
			_ = qn_file.Remove(path)
		}()

		c := gcfg.New()
		c.SetFileName(path)
		t.Assert(c.Get("v1"), 1)
		t.AssertEQ(c.GetInt("v1"), 1)
		t.AssertEQ(c.GetInt8("v1"), int8(1))
		t.AssertEQ(c.GetInt16("v1"), int16(1))
		t.AssertEQ(c.GetInt32("v1"), int32(1))
		t.AssertEQ(c.GetInt64("v1"), int64(1))
		t.AssertEQ(c.GetUint("v1"), uint(1))
		t.AssertEQ(c.GetUint8("v1"), uint8(1))
		t.AssertEQ(c.GetUint16("v1"), uint16(1))
		t.AssertEQ(c.GetUint32("v1"), uint32(1))
		t.AssertEQ(c.GetUint64("v1"), uint64(1))

		t.AssertEQ(c.GetVar("v1").String(), "1")
		t.AssertEQ(c.GetVar("v1").Bool(), true)
		t.AssertEQ(c.GetVar("v2").String(), "true")
		t.AssertEQ(c.GetVar("v2").Bool(), true)

		t.AssertEQ(c.GetString("v1"), "1")
		t.AssertEQ(c.GetFloat32("v4"), float32(1.234))
		t.AssertEQ(c.GetFloat64("v4"), float64(1.234))
		t.AssertEQ(c.GetString("v2"), "true")
		t.AssertEQ(c.GetBool("v2"), true)
		t.AssertEQ(c.GetBool("v3"), false)

		t.AssertEQ(c.Contains("v1"), true)
		t.AssertEQ(c.Contains("v2"), true)
		t.AssertEQ(c.Contains("v3"), true)
		t.AssertEQ(c.Contains("v4"), true)
		t.AssertEQ(c.Contains("v5"), false)

		t.AssertEQ(c.GetInts("array"), []int{1, 2, 3})
		t.AssertEQ(c.GetStrings("array"), []string{"1", "2", "3"})
		t.AssertEQ(c.GetArray("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetInterfaces("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetMap("redis"), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		t.AssertEQ(c.FilePath(), qn_file.Pwd()+qn_file.Separator+path)

	})
}

func Test_Instance(t *testing.T) {
	config := `
{
	"array": [
		1,
		2,
		3
	],
	"redis": {
		"cache": "127.0.0.1:6379,1",
		"disk": "127.0.0.1:6379,0"
	},
	"v1": 1,
	"v2": "true",
	"v3": "off",
	"v4": "1.234"
}
`
	qn_test.C(t, func(t *qn_test.T) {
		path := gcfg.DEFAULT_CONFIG_FILE
		err := qn_file.PutContents(path, config)
		t.Assert(err, nil)
		defer func() {
			t.Assert(qn_file.Remove(path), nil)
		}()

		c := gcfg.Instance()
		t.Assert(c.Get("v1"), 1)
		t.AssertEQ(c.GetInt("v1"), 1)
		t.AssertEQ(c.GetInt8("v1"), int8(1))
		t.AssertEQ(c.GetInt16("v1"), int16(1))
		t.AssertEQ(c.GetInt32("v1"), int32(1))
		t.AssertEQ(c.GetInt64("v1"), int64(1))
		t.AssertEQ(c.GetUint("v1"), uint(1))
		t.AssertEQ(c.GetUint8("v1"), uint8(1))
		t.AssertEQ(c.GetUint16("v1"), uint16(1))
		t.AssertEQ(c.GetUint32("v1"), uint32(1))
		t.AssertEQ(c.GetUint64("v1"), uint64(1))

		t.AssertEQ(c.GetVar("v1").String(), "1")
		t.AssertEQ(c.GetVar("v1").Bool(), true)
		t.AssertEQ(c.GetVar("v2").String(), "true")
		t.AssertEQ(c.GetVar("v2").Bool(), true)

		t.AssertEQ(c.GetString("v1"), "1")
		t.AssertEQ(c.GetFloat32("v4"), float32(1.234))
		t.AssertEQ(c.GetFloat64("v4"), float64(1.234))
		t.AssertEQ(c.GetString("v2"), "true")
		t.AssertEQ(c.GetBool("v2"), true)
		t.AssertEQ(c.GetBool("v3"), false)

		t.AssertEQ(c.Contains("v1"), true)
		t.AssertEQ(c.Contains("v2"), true)
		t.AssertEQ(c.Contains("v3"), true)
		t.AssertEQ(c.Contains("v4"), true)
		t.AssertEQ(c.Contains("v5"), false)

		t.AssertEQ(c.GetInts("array"), []int{1, 2, 3})
		t.AssertEQ(c.GetStrings("array"), []string{"1", "2", "3"})
		t.AssertEQ(c.GetArray("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetInterfaces("array"), []interface{}{1, 2, 3})
		t.AssertEQ(c.GetMap("redis"), map[string]interface{}{
			"disk":  "127.0.0.1:6379,0",
			"cache": "127.0.0.1:6379,1",
		})
		t.AssertEQ(c.FilePath(), qn_file.Pwd()+qn_file.Separator+path)

	})
}

func TestCfg_New(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		os.Setenv("GF_GCFG_PATH", "config")
		c := gcfg.New("config.yml")
		t.Assert(c.Get("name"), nil)
		t.Assert(c.GetFileName(), "config.yml")

		configPath := qn_file.Pwd() + qn_file.Separator + "config"
		_ = qn_file.Mkdir(configPath)
		defer qn_file.Remove(configPath)

		c = gcfg.New("config.yml")
		t.Assert(c.Get("name"), nil)

		_ = os.Unsetenv("GF_GCFG_PATH")
		c = gcfg.New("config.yml")
		t.Assert(c.Get("name"), nil)
	})
}

func TestCfg_SetPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := gcfg.New("config.yml")
		err := c.SetPath("tmp")
		t.AssertNE(err, nil)
		err = c.SetPath("gcfg.go")
		t.AssertNE(err, nil)
		t.Assert(c.Get("name"), nil)
	})
}

func TestCfg_SetViolenceCheck(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := gcfg.New("config.yml")
		c.SetViolenceCheck(true)
		t.Assert(c.Get("name"), nil)
	})
}

func TestCfg_AddPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := gcfg.New("config.yml")
		err := c.AddPath("tmp")
		t.AssertNE(err, nil)
		err = c.AddPath("gcfg.go")
		t.AssertNE(err, nil)
		t.Assert(c.Get("name"), nil)
	})
}

func TestCfg_FilePath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := gcfg.New("config.yml")
		path := c.FilePath("tmp")
		t.Assert(path, "")
		path = c.FilePath("tmp")
		t.Assert(path, "")
	})
}

func TestCfg_Get(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var err error
		configPath := qn_file.TempDir(qn_time.TimestampNanoStr())
		err = qn_file.Mkdir(configPath)
		t.Assert(err, nil)
		defer qn_file.Remove(configPath)

		defer qn_file.Chdir(qn_file.Pwd())
		err = qn_file.Chdir(configPath)
		t.Assert(err, nil)

		err = qn_file.PutContents(
			qn_file.Join(configPath, "config.yml"),
			"wrong config",
		)
		t.Assert(err, nil)
		c := gcfg.New("config.yml")
		t.Assert(c.Get("name"), nil)
		t.Assert(c.GetVar("name").Val(), nil)
		t.Assert(c.Contains("name"), false)
		t.Assert(c.GetMap("name"), nil)
		t.Assert(c.GetArray("name"), nil)
		t.Assert(c.GetString("name"), "")
		t.Assert(c.GetStrings("name"), nil)
		t.Assert(c.GetInterfaces("name"), nil)
		t.Assert(c.GetBool("name"), false)
		t.Assert(c.GetFloat32("name"), 0)
		t.Assert(c.GetFloat64("name"), 0)
		t.Assert(c.GetFloats("name"), nil)
		t.Assert(c.GetInt("name"), 0)
		t.Assert(c.GetInt8("name"), 0)
		t.Assert(c.GetInt16("name"), 0)
		t.Assert(c.GetInt32("name"), 0)
		t.Assert(c.GetInt64("name"), 0)
		t.Assert(c.GetInts("name"), nil)
		t.Assert(c.GetUint("name"), 0)
		t.Assert(c.GetUint8("name"), 0)
		t.Assert(c.GetUint16("name"), 0)
		t.Assert(c.GetUint32("name"), 0)
		t.Assert(c.GetUint64("name"), 0)
		t.Assert(c.GetTime("name").Format("2006-01-02"), "0001-01-01")
		t.Assert(c.Getqn_time("name"), nil)
		t.Assert(c.GetDuration("name").String(), "0s")
		name := struct {
			Name string
		}{}
		t.Assert(c.GetStruct("name", &name) == nil, false)

		c.Clear()

		arr, _ := qn_json.Encode(
			g.Map{
				"name":   "gf",
				"time":   "2019-06-12",
				"person": g.Map{"name": "gf"},
				"floats": g.Slice{1, 2, 3},
			},
		)
		err = qn_file.PutBytes(
			qn_file.Join(configPath, "config.yml"),
			arr,
		)
		t.Assert(err, nil)
		t.Assert(c.GetTime("time").Format("2006-01-02"), "2019-06-12")
		t.Assert(c.Getqn_time("time").Format("Y-m-d"), "2019-06-12")
		t.Assert(c.GetDuration("time").String(), "0s")

		err = c.GetStruct("person", &name)
		t.Assert(err, nil)
		t.Assert(name.Name, "gf")
		t.Assert(c.GetFloats("floats") == nil, false)
	})
}

func TestCfg_Instance(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(gcfg.Instance("gf") != nil, true)
	})
	qn_test.C(t, func(t *qn_test.T) {
		pwd := qn_file.Pwd()
		qn_file.Chdir(qn_file.Join(gdebug.TestDataPath()))
		defer qn_file.Chdir(pwd)
		t.Assert(gcfg.Instance("c1") != nil, true)
		t.Assert(gcfg.Instance("c1").Get("my-config"), "1")
		t.Assert(gcfg.Instance("folder1/c1").Get("my-config"), "2")
	})
}

func TestCfg_Config(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		gcfg.SetContent("gf", "config.yml")
		t.Assert(gcfg.GetContent("config.yml"), "gf")
		gcfg.SetContent("gf1", "config.yml")
		t.Assert(gcfg.GetContent("config.yml"), "gf1")
		gcfg.RemoveContent("config.yml")
		gcfg.ClearContent()
		t.Assert(gcfg.GetContent("name"), "")
	})
}
