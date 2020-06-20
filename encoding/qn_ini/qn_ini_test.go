// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_int_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_int"
	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/test/qn_test"
)

var iniContent = `

;注释
aa=bb
[addr] 
#注释
ip = 127.0.0.1
port=9001
enable=true

	[DBINFO]
	type=mysql
	user=root
	password=password
[键]
呵呵=值

`

func TestDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		res, err := qn_int.Decode([]byte(iniContent))
		if err != nil {
			qn_test.Fatal(err)
		}
		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")
		t.Assert(res["键"].(map[string]interface{})["呵呵"], "值")
	})

	qn_test.C(t, func(t *qn_test.T) {
		errContent := `
		a = b
`
		_, err := qn_int.Decode([]byte(errContent))
		if err == nil {
			qn_test.Fatal(err)
		}
	})
}

func TestEncode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		iniMap, err := qn_int.Decode([]byte(iniContent))
		if err != nil {
			qn_test.Fatal(err)
		}

		iniStr, err := qn_int.Encode(iniMap)
		if err != nil {
			qn_test.Fatal(err)
		}

		res, err := qn_int.Decode(iniStr)
		if err != nil {
			qn_test.Fatal(err)
		}

		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")

	})
}

func TestToJson(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		jsonStr, err := qn_int.ToJson([]byte(iniContent))
		if err != nil {
			qn_test.Fatal(err)
		}

		json, err := qn_json.LoadContent(jsonStr)
		if err != nil {
			qn_test.Fatal(err)
		}

		iniMap, err := qn_int.Decode([]byte(iniContent))
		t.Assert(err, nil)

		t.Assert(iniMap["addr"].(map[string]interface{})["ip"], json.GetString("addr.ip"))
		t.Assert(iniMap["addr"].(map[string]interface{})["port"], json.GetString("addr.port"))
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["user"], json.GetString("DBINFO.user"))
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["type"], json.GetString("DBINFO.type"))
	})
}
