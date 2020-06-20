// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.
package qn_toml_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_parser"
	"github.com/qnsoft/common/encoding/qn_toml"
	"github.com/qnsoft/common/test/qn_test"
)

var tomlStr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

var tomlErr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

func TestEncode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := qn_toml.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		p, err := qn_parser.LoadContent(res)
		if err != nil {
			t.Errorf("parser failed. %v", err)
			return
		}

		t.Assert(p.GetString("toml"), tomlStr)
	})

	qn_test.C(t, func(t *qn_test.T) {
		_, err := qn_toml.Encode(tomlErr)
		if err == nil {
			t.Errorf("encode should be failed. %v", err)
			return
		}
	})
}

func TestDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := qn_toml.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		decodeStr, err := qn_toml.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}

		t.Assert(decodeStr.(map[string]interface{})["toml"], tomlStr)

		decodeStr1 := make(map[string]interface{})
		err = qn_toml.DecodeTo(res, &decodeStr1)
		if err != nil {
			t.Errorf("decodeTo failed. %v", err)
			return
		}
		t.Assert(decodeStr1["toml"], tomlStr)
	})

	qn_test.C(t, func(t *qn_test.T) {
		_, err := qn_toml.Decode([]byte(tomlErr))
		if err == nil {
			t.Errorf("decode failed. %v", err)
			return
		}

		decodeStr1 := make(map[string]interface{})
		err = qn_toml.DecodeTo([]byte(tomlErr), &decodeStr1)
		if err == nil {
			t.Errorf("decodeTo failed. %v", err)
			return
		}
	})
}

func TestToJson(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := qn_toml.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		jsonToml, err := qn_toml.ToJson(res)
		if err != nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}

		p, err := qn_parser.LoadContent(res)
		if err != nil {
			t.Errorf("parser failed. %v", err)
			return
		}
		expectJson, err := p.ToJson()
		if err != nil {
			t.Errorf("parser ToJson failed. %v", err)
			return
		}
		t.Assert(jsonToml, expectJson)
	})

	qn_test.C(t, func(t *qn_test.T) {
		_, err := qn_toml.ToJson([]byte(tomlErr))
		if err == nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}
	})
}
