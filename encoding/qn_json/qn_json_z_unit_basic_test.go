// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/container/gmap"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_New(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(data)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
	})

	qn_test.C(t, func(t *qn_test.T) {
		m := gmap.NewAnyAnyMapFrom(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		j := qn_json.New(m)
		t.Assert(j.Get("k1"), "v1")
		t.Assert(j.Get("k2"), "v2")
		t.Assert(j.Get("k3"), nil)
	})
}

func Test_Valid(t *testing.T) {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_json.Valid(data1), true)
		t.Assert(qn_json.Valid(data2), false)
	})
}

func Test_Encode(t *testing.T) {
	value := g.Slice{1, 2, 3}
	qn_test.C(t, func(t *qn_test.T) {
		b, err := qn_json.Encode(value)
		t.Assert(err, nil)
		t.Assert(b, []byte(`[1,2,3]`))
	})
}

func Test_Decode(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		v, err := qn_json.Decode(data)
		t.Assert(err, nil)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var v interface{}
		err := qn_json.DecodeTo(data, &v)
		t.Assert(err, nil)
		t.Assert(v, g.Map{
			"n": 123456789,
			"a": g.Slice{1, 2, 3},
			"m": g.Map{
				"k": "v",
			},
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func Test_SplitChar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		j.SetSplitChar(byte('#'))
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m#k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a#1"), 2)
	})
}

func Test_ViolenceCheck(t *testing.T) {
	data := []byte(`{"m":{"a":[1,2,3], "v1.v2":"4"}}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.Assert(j.Get("m.a.2"), 3)
		t.Assert(j.Get("m.v1.v2"), nil)
		j.SetViolenceCheck(true)
		t.Assert(j.Get("m.v1.v2"), 4)
	})
}

func Test_GetVar(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.Assert(j.GetVar("n").String(), "123456789")
		t.Assert(j.GetVar("m").Map(), g.Map{"k": "v"})
		t.Assert(j.GetVar("a").Interfaces(), g.Slice{1, 2, 3})
		t.Assert(j.GetVar("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.GetVar("a").Array(), g.Slice{1, 2, 3})
	})
}

func Test_GetMap(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.Assert(j.GetMap("n"), nil)
		t.Assert(j.GetMap("m"), g.Map{"k": "v"})
		t.Assert(j.GetMap("a"), g.Map{"1": "2", "3": nil})
	})
}

func Test_GetJson(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		j2 := j.GetJson("m")
		t.AssertNE(j2, nil)
		t.Assert(j2.Get("k"), "v")
		t.Assert(j2.Get("a"), nil)
		t.Assert(j2.Get("n"), nil)
	})
}

func Test_GetArray(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.Assert(j.GetArray("n"), g.Array{123456789})
		t.Assert(j.GetArray("m"), g.Array{g.Map{"k": "v"}})
		t.Assert(j.GetArray("a"), g.Array{1, 2, 3})
	})
}

func Test_GetString(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.AssertEQ(j.GetString("n"), "123456789")
		t.AssertEQ(j.GetString("m"), `{"k":"v"}`)
		t.AssertEQ(j.GetString("a"), `[1,2,3]`)
		t.AssertEQ(j.GetString("i"), "")
	})
}

func Test_GetStrings(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.AssertEQ(j.GetStrings("n"), g.SliceStr{"123456789"})
		t.AssertEQ(j.GetStrings("m"), g.SliceStr{`{"k":"v"}`})
		t.AssertEQ(j.GetStrings("a"), g.SliceStr{"1", "2", "3"})
		t.AssertEQ(j.GetStrings("i"), nil)
	})
}

func Test_GetInterfaces(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.DecodeToJson(data)
		t.Assert(err, nil)
		t.AssertEQ(j.GetInterfaces("n"), g.Array{123456789})
		t.AssertEQ(j.GetInterfaces("m"), g.Array{g.Map{"k": "v"}})
		t.AssertEQ(j.GetInterfaces("a"), g.Array{1, 2, 3})
	})
}

func Test_Len(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Append("a", 1)
		p.Append("a", 2)
		t.Assert(p.Len("a"), 2)
	})
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Append("a.b", 1)
		p.Append("a.c", 2)
		t.Assert(p.Len("a"), 2)
	})
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Set("a", 1)
		t.Assert(p.Len("a"), -1)
	})
}

func Test_Append(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Append("a", 1)
		p.Append("a", 2)
		t.Assert(p.Get("a"), g.Slice{1, 2})
	})
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Append("a.b", 1)
		p.Append("a.c", 2)
		t.Assert(p.Get("a"), g.Map{
			"b": g.Slice{1},
			"c": g.Slice{2},
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(nil)
		p.Set("a", 1)
		err := p.Append("a", 2)
		t.AssertNE(err, nil)
		t.Assert(p.Get("a"), 1)
	})
}

func TestJson_ToJson(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New(1)
		s, e := p.ToJsonString()
		t.Assert(e, nil)
		t.Assert(s, "1")
	})
	qn_test.C(t, func(t *qn_test.T) {
		p := qn_json.New("a")
		s, e := p.ToJsonString()
		t.Assert(e, nil)
		t.Assert(s, `"a"`)
	})
}

func TestJson_Default(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(nil)
		t.AssertEQ(j.Get("no", 100), 100)
		t.AssertEQ(j.GetString("no", 100), "100")
		t.AssertEQ(j.GetBool("no", "on"), true)
		t.AssertEQ(j.GetInt("no", 100), 100)
		t.AssertEQ(j.GetInt8("no", 100), int8(100))
		t.AssertEQ(j.GetInt16("no", 100), int16(100))
		t.AssertEQ(j.GetInt32("no", 100), int32(100))
		t.AssertEQ(j.GetInt64("no", 100), int64(100))
		t.AssertEQ(j.GetUint("no", 100), uint(100))
		t.AssertEQ(j.GetUint8("no", 100), uint8(100))
		t.AssertEQ(j.GetUint16("no", 100), uint16(100))
		t.AssertEQ(j.GetUint32("no", 100), uint32(100))
		t.AssertEQ(j.GetUint64("no", 100), uint64(100))
		t.AssertEQ(j.GetFloat32("no", 123.456), float32(123.456))
		t.AssertEQ(j.GetFloat64("no", 123.456), float64(123.456))
		t.AssertEQ(j.GetArray("no", g.Slice{1, 2, 3}), g.Slice{1, 2, 3})
		t.AssertEQ(j.GetInts("no", g.Slice{1, 2, 3}), g.SliceInt{1, 2, 3})
		t.AssertEQ(j.GetFloats("no", g.Slice{1, 2, 3}), []float64{1, 2, 3})
		t.AssertEQ(j.GetMap("no", g.Map{"k": "v"}), g.Map{"k": "v"})
		t.AssertEQ(j.GetVar("no", 123.456).Float64(), float64(123.456))
		t.AssertEQ(j.GetJson("no", g.Map{"k": "v"}).Get("k"), "v")
		t.AssertEQ(j.GetJsons("no", g.Slice{
			g.Map{"k1": "v1"},
			g.Map{"k2": "v2"},
			g.Map{"k3": "v3"},
		})[0].Get("k1"), "v1")
		t.AssertEQ(j.GetJsonMap("no", g.Map{
			"m1": g.Map{"k1": "v1"},
			"m2": g.Map{"k2": "v2"},
		})["m2"].Get("k2"), "v2")
	})
}

func Test_Convert(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(`{"name":"gf"}`)
		arr, err := j.ToXml()
		t.Assert(err, nil)
		t.Assert(string(arr), "<name>gf</name>")
		arr, err = j.ToXmlIndent()
		t.Assert(err, nil)
		t.Assert(string(arr), "<name>gf</name>")
		str, err := j.ToXmlString()
		t.Assert(err, nil)
		t.Assert(str, "<name>gf</name>")
		str, err = j.ToXmlIndentString()
		t.Assert(err, nil)
		t.Assert(str, "<name>gf</name>")

		arr, err = j.ToJsonIndent()
		t.Assert(err, nil)
		t.Assert(string(arr), "{\n\t\"name\": \"gf\"\n}")
		str, err = j.ToJsonIndentString()
		t.Assert(err, nil)
		t.Assert(string(arr), "{\n\t\"name\": \"gf\"\n}")

		arr, err = j.ToYaml()
		t.Assert(err, nil)
		t.Assert(string(arr), "name: gf\n")
		str, err = j.ToYamlString()
		t.Assert(err, nil)
		t.Assert(string(arr), "name: gf\n")

		arr, err = j.ToToml()
		t.Assert(err, nil)
		t.Assert(string(arr), "name = \"gf\"\n")
		str, err = j.ToTomlString()
		t.Assert(err, nil)
		t.Assert(string(arr), "name = \"gf\"\n")
	})
}

func Test_Convert2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		name := struct {
			Name string
		}{}
		j := qn_json.New(`{"name":"gf","time":"2019-06-12"}`)
		t.Assert(j.Value().(g.Map)["name"], "gf")
		t.Assert(j.GetMap("name1"), nil)
		t.AssertNE(j.GetJson("name1"), nil)
		t.Assert(j.GetJsons("name1"), nil)
		t.Assert(j.GetJsonMap("name1"), nil)
		t.Assert(j.Contains("name1"), false)
		t.Assert(j.GetVar("name1").IsNil(), true)
		t.Assert(j.GetVar("name").IsNil(), false)
		t.Assert(j.Len("name1"), -1)
		t.Assert(j.GetTime("time").Format("2006-01-02"), "2019-06-12")
		t.Assert(j.GetGTime("time").Format("Y-m-d"), "2019-06-12")
		t.Assert(j.GetDuration("time").String(), "0s")

		err := j.ToStruct(&name)
		t.Assert(err, nil)
		t.Assert(name.Name, "gf")
		//j.Dump()
		t.Assert(err, nil)

		j = qn_json.New(`{"person":{"name":"gf"}}`)
		err = j.GetStruct("person", &name)
		t.Assert(err, nil)
		t.Assert(name.Name, "gf")

		j = qn_json.New(`{"name":"gf""}`)
		//j.Dump()
		t.Assert(err, nil)

		j = qn_json.New(`[1,2,3]`)
		t.Assert(len(j.ToArray()), 3)
	})
}

func Test_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(`{"name":"gf","time":"2019-06-12"}`)
		j.SetViolenceCheck(true)
		t.Assert(j.Get(""), nil)
		t.Assert(j.Get(".").(g.Map)["name"], "gf")
		t.Assert(j.Get(".").(g.Map)["name1"], nil)
		j.SetViolenceCheck(false)
		t.Assert(j.Get(".").(g.Map)["name"], "gf")

		err := j.Set("name", "gf1")
		t.Assert(err, nil)
		t.Assert(j.Get("name"), "gf1")

		j = qn_json.New(`[1,2,3]`)
		err = j.Set("\"0\".1", 11)
		t.Assert(err, nil)
		t.Assert(j.Get("1"), 11)

		j = qn_json.New(`[1,2,3]`)
		err = j.Set("11111111111111111111111", 11)
		t.AssertNE(err, nil)

		j = qn_json.New(`[1,2,3]`)
		err = j.Remove("1")
		t.Assert(err, nil)
		t.Assert(j.Get("0"), 1)
		t.Assert(len(j.ToArray()), 2)

		j = qn_json.New(`[1,2,3]`)
		// If index 0 is delete, its next item will be at index 0.
		t.Assert(j.Remove("0"), nil)
		t.Assert(j.Remove("0"), nil)
		t.Assert(j.Remove("0"), nil)
		t.Assert(j.Get("0"), nil)
		t.Assert(len(j.ToArray()), 0)

		j = qn_json.New(`[1,2,3]`)
		err = j.Remove("3")
		t.Assert(err, nil)
		t.Assert(j.Get("0"), 1)
		t.Assert(len(j.ToArray()), 3)

		j = qn_json.New(`[1,2,3]`)
		err = j.Remove("0.3")
		t.Assert(err, nil)
		t.Assert(len(j.Get("0").([]interface{})), 3)

		j = qn_json.New(`[1,2,3]`)
		err = j.Remove("0.a")
		t.Assert(err, nil)
		t.Assert(j.Get("0"), 1)

		name := struct {
			Name string
		}{Name: "gf"}
		j = qn_json.New(name)
		t.Assert(j.Get("Name"), "gf")
		err = j.Remove("Name")
		t.Assert(err, nil)
		t.Assert(j.Get("Name"), nil)

		err = j.Set("Name", "gf1")
		t.Assert(err, nil)
		t.Assert(j.Get("Name"), "gf1")

		j = qn_json.New(nil)
		err = j.Remove("Name")
		t.Assert(err, nil)
		t.Assert(j.Get("Name"), nil)

		j = qn_json.New(name)
		t.Assert(j.Get("Name"), "gf")
		err = j.Set("Name1", g.Map{"Name": "gf1"})
		t.Assert(err, nil)
		t.Assert(j.Get("Name1").(g.Map)["Name"], "gf1")
		err = j.Set("Name2", g.Slice{1, 2, 3})
		t.Assert(err, nil)
		t.Assert(j.Get("Name2").(g.Slice)[0], 1)
		err = j.Set("Name3", name)
		t.Assert(err, nil)
		t.Assert(j.Get("Name3").(g.Map)["Name"], "gf")
		err = j.Set("Name4", &name)
		t.Assert(err, nil)
		t.Assert(j.Get("Name4").(g.Map)["Name"], "gf")
		arr := [3]int{1, 2, 3}
		err = j.Set("Name5", arr)
		t.Assert(err, nil)
		t.Assert(j.Get("Name5").(g.Array)[0], 1)

	})
}

func Test_IsNil(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_json.New(nil)
		t.Assert(j.IsNil(), true)
	})
}
