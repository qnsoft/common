// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/os/gfile"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Load_JSON1(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.json"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_json.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_JSON2(t *testing.T) {
	data := []byte(`{"n":123456789000000000000, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789000000000000")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_XML(t *testing.T) {
	data := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n></doc>`)
	// XML
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("doc.n"), "123456789")
		t.Assert(j.Get("doc.m"), qn.Map{"k": "v"})
		t.Assert(j.Get("doc.m.k"), "v")
		t.Assert(j.Get("doc.a"), qn.Slice{"1", "2", "3"})
		t.Assert(j.Get("doc.a.1"), 2)
	})
	// XML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.xml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_json.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("doc.n"), "123456789")
		t.Assert(j.Get("doc.m"), qn.Map{"k": "v"})
		t.Assert(j.Get("doc.m.k"), "v")
		t.Assert(j.Get("doc.a"), qn.Slice{"1", "2", "3"})
		t.Assert(j.Get("doc.a.1"), 2)
	})

	// XML
	qn_test.C(t, func(t *qn_test.T) {
		xml := `<?xml version="1.0"?>

	<Output type="o">
	<itotalSize>0</itotalSize>
	<ipageSize>1</ipageSize>
	<ipageIndex>2</ipageIndex>
	<itotalRecords>GF框架</itotalRecords>
	<nworkOrderDtos/>
	<nworkOrderFrontXML/>
	</Output>`
		j, err := qn_json.LoadContent(xml)
		t.Assert(err, nil)
		t.Assert(j.Get("Output.ipageIndex"), "2")
		t.Assert(j.Get("Output.itotalRecords"), "GF框架")
	})
}

func Test_Load_YAML1(t *testing.T) {
	data := []byte(`
a:
- 1
- 2
- 3
m:
 k: v
"n": 123456789
    `)
	// YAML
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
	// YAML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.yaml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_json.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_YAML2(t *testing.T) {
	data := []byte("i : 123456789")
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_TOML1(t *testing.T) {
	data := []byte(`
a = ["1", "2", "3"]
n = 123456789

[m]
  k = "v"
`)
	// TOML
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{"1", "2", "3"})
		t.Assert(j.Get("a.1"), 2)
	})
	// TOML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.toml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_json.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{"1", "2", "3"})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_TOML2(t *testing.T) {
	data := []byte("i=123456789")
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(nil)
		t.Assert(j.Value(), nil)
		_, err := qn_json.Decode(nil)
		t.AssertNE(err, nil)
		_, err = qn_json.DecodeToJson(nil)
		t.AssertNE(err, nil)
		j, err = qn_json.LoadContent(nil)
		t.Assert(err, nil)
		t.Assert(j.Value(), nil)

		j, err = qn_json.LoadContent(`{"name": "gf"}`)
		t.Assert(err, nil)

		j, err = qn_json.LoadContent(`{"name": "gf"""}`)
		t.AssertNE(err, nil)

		j = qn_json.New(&qn.Map{"name": "gf"})
		t.Assert(j.GetString("name"), "gf")

	})
}

func Test_Load_Ini(t *testing.T) {
	var data = `

;注释

[addr] 
#注释
ip = 127.0.0.1
port=9001
enable=true

	[DBINFO]
	type=mysql
	user=root
	password=password

`

	qn_test.C(t, func(t *qn_test.T) {
		json, err := qn_json.LoadContent(data)
		if err != nil {
			qn_test.Fatal(err)
		}

		t.Assert(json.GetString("addr.ip"), "127.0.0.1")
		t.Assert(json.GetString("addr.port"), "9001")
		t.Assert(json.GetString("addr.enable"), "true")
		t.Assert(json.GetString("DBINFO.type"), "mysql")
		t.Assert(json.GetString("DBINFO.user"), "root")
		t.Assert(json.GetString("DBINFO.password"), "password")

		_, err = json.ToIni()
		if err != nil {
			qn_test.Fatal(err)
		}
	})
}
