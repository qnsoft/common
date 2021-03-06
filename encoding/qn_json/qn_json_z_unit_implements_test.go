// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/util/gconv"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/test/qn_test"
)

func TestJson_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(nil)
		err := json.Unmarshal(data, j)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), qn.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func TestJson_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Json *qn_json.Json
	}
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := gconv.Struct(qn.Map{
			"name": "john",
			"json": []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`),
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n"), "123456789")
		t.Assert(v.Json.Get("m"), qn.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k"), "v")
		t.Assert(v.Json.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1"), 2)
	})
	// Map
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := gconv.Struct(qn.Map{
			"name": "john",
			"json": qn.Map{
				"n": 123456789,
				"m": qn.Map{"k": "v"},
				"a": qn.Slice{1, 2, 3},
			},
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n"), "123456789")
		t.Assert(v.Json.Get("m"), qn.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k"), "v")
		t.Assert(v.Json.Get("a"), qn.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1"), 2)
	})
}
