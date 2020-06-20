// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/qnsoft/common/encoding/qn_parser"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/gfile"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Load_JSON(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.json"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_parser.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_XML(t *testing.T) {
	data := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n></doc>`)
	// XML
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("doc.n"), "123456789")
		t.Assert(j.Get("doc.m"), g.Map{"k": "v"})
		t.Assert(j.Get("doc.m.k"), "v")
		t.Assert(j.Get("doc.a"), g.Slice{"1", "2", "3"})
		t.Assert(j.Get("doc.a.1"), 2)
	})
	// XML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.xml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_parser.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("doc.n"), "123456789")
		t.Assert(j.Get("doc.m"), g.Map{"k": "v"})
		t.Assert(j.Get("doc.m.k"), "v")
		t.Assert(j.Get("doc.a"), g.Slice{"1", "2", "3"})
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
		j, err := qn_parser.LoadContent(xml)
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
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
	// YAML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.yaml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_parser.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_YAML2(t *testing.T) {
	data := []byte("i : 123456789")
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_TOML1(t *testing.T) {
	data := []byte(`
a = ["1", "2", "3"]
n = "123456789"

[m]
  k = "v"
`)
	// TOML
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{"1", "2", "3"})
		t.Assert(j.Get("a.1"), 2)
	})
	// TOML
	qn_test.C(t, func(t *qn_test.T) {
		path := "test.toml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := qn_parser.Load(path)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{"1", "2", "3"})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_TOML2(t *testing.T) {
	data := []byte("i=123456789")
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_parser.LoadContent(data)
		t.Assert(err, nil)
		t.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_Nil(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_parser.New(nil)
		t.Assert(p.Value(), nil)
		file := "test22222.json"
		filePath := gfile.Pwd() + gfile.Separator + file
		ioutil.WriteFile(filePath, []byte("{"), 0644)
		defer gfile.Remove(filePath)
		_, err := qn_parser.Load(file)
		t.AssertNE(err, nil)
		_, err = qn_parser.LoadContent("{")
		t.AssertNE(err, nil)
	})
}
