// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Load_NewWithTag(t *testing.T) {
	type User struct {
		Age  int    `xml:"age-xml"  json:"age-json"`
		Name string `xml:"name-xml" json:"name-json"`
		Addr string `xml:"addr-xml" json:"addr-json"`
	}
	data := User{
		Age:  18,
		Name: "john",
		Addr: "chengdu",
	}
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(data)
		t.AssertNE(j, nil)
		t.Assert(j.Get("age-xml"), nil)
		t.Assert(j.Get("age-json"), data.Age)
		t.Assert(j.Get("name-xml"), nil)
		t.Assert(j.Get("name-json"), data.Name)
		t.Assert(j.Get("addr-xml"), nil)
		t.Assert(j.Get("addr-json"), data.Addr)
	})
	// XML
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.NewWithTag(data, "xml")
		t.AssertNE(j, nil)
		t.Assert(j.Get("age-xml"), data.Age)
		t.Assert(j.Get("age-json"), nil)
		t.Assert(j.Get("name-xml"), data.Name)
		t.Assert(j.Get("name-json"), nil)
		t.Assert(j.Get("addr-xml"), data.Addr)
		t.Assert(j.Get("addr-json"), nil)
	})
}
