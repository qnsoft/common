// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_yaml_test

import (
	"testing"

	"github.com/qnsoft/common/internal/json"

	"github.com/qnsoft/common/encoding/qn_parser"

	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/encoding/qn_yaml"
	"github.com/qnsoft/common/test/qn_test"
)

var yamlStr string = `
#即表示url属性值；
url: https://goframe.org

#数组，即表示server为[a,b,c]
server:
    - 120.168.117.21
    - 120.168.117.22
#常量
pi: 3.14   #定义一个数值3.14
hasChild: true  #定义一个boolean值
name: '你好YAML'   #定义一个字符串
`

var yamlErr string = `
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

func Test_Decode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		result, err := qn_yaml.Decode([]byte(yamlStr))
		t.Assert(err, nil)

		m, ok := result.(map[string]interface{})
		t.Assert(ok, true)
		t.Assert(m, map[string]interface{}{
			"url":      "https://goframe.org",
			"server":   g.Slice{"120.168.117.21", "120.168.117.22"},
			"pi":       3.14,
			"hasChild": true,
			"name":     "你好YAML",
		})
	})
}

func Test_DecodeTo(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		result := make(map[string]interface{})
		err := qn_yaml.DecodeTo([]byte(yamlStr), &result)
		t.Assert(err, nil)
		t.Assert(result, map[string]interface{}{
			"url":      "https://goframe.org",
			"server":   g.Slice{"120.168.117.21", "120.168.117.22"},
			"pi":       3.14,
			"hasChild": true,
			"name":     "你好YAML",
		})
	})
}

func Test_DecodeError(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		_, err := qn_yaml.Decode([]byte(yamlErr))
		t.AssertNE(err, nil)

		result := make(map[string]interface{})
		err = qn_yaml.DecodeTo([]byte(yamlErr), &result)
		t.AssertNE(err, nil)
	})
}

func Test_DecodeMapToJson(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		data := []byte(`
m:
 k: v
    `)
		v, err := qn_yaml.Decode(data)
		t.Assert(err, nil)
		b, err := json.Marshal(v)
		t.Assert(err, nil)
		t.Assert(b, `{"m":{"k":"v"}}`)
	})
}

func Test_ToJson(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]string)
		m["yaml"] = yamlStr
		res, err := qn_yaml.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		jsonyaml, err := qn_yaml.ToJson(res)
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
		t.Assert(jsonyaml, expectJson)
	})

	qn_test.C(t, func(t *qn_test.T) {
		_, err := qn_yaml.ToJson([]byte(yamlErr))
		if err == nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}
	})
}
