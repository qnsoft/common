// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv_test

import (
	"testing"

	"github.com/gogf/gf/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_MapToMap1(t *testing.T) {
	// map[int]int -> map[string]string
	// empty original map.
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn.MapIntInt{}
		m2 := qn.MapStrStr{}
		t.Assert(qn_conv.MapToMap(m1, &m2), nil)
		t.Assert(len(m1), len(m2))
	})
	// map[int]int -> map[string]string
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn.MapIntInt{
			1: 100,
			2: 200,
		}
		m2 := qn.MapStrStr{}
		t.Assert(qn_conv.MapToMap(m1, &m2), nil)
		t.Assert(m2["1"], m1[1])
		t.Assert(m2["2"], m1[2])
	})
	// map[string]interface{} -> map[string]string
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn.Map{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := qn.MapStrStr{}
		t.Assert(qn_conv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]string -> map[string]interface{}
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := qn.Map{}
		t.Assert(qn_conv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// map[string]interface{} -> map[interface{}]interface{}
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := qn.MapAnyAny{}
		t.Assert(qn_conv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
}

func Test_MapToMap2(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := qn.Map{
		"key": qn.Map{
			"id":   1,
			"name": "john",
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]User)
		err := qn_conv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string]User)(nil)
		err := qn_conv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string]*User)
		err := qn_conv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string]*User)(nil)
		err := qn_conv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMapDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := qn.Map{
		"key": qn.Map{
			"id":   1,
			"name": "john",
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string]*User)(nil)
		err := qn_conv.MapToMap(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 0)
		t.Assert(m["key"].Name, "john")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string]*User)(nil)
		err := qn_conv.MapToMapDeep(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMaps1(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}
	params :=qn.SlMap{
		"key1": g.Slice{
			qn.Map{"id": 1, "name": "john"},
			qn.Map{"id": 2, "name": "smith"},
		},qn.Sl
		"key2": g.Slice{
			qn.Map{"id": 3, "name": "green"},
			qn.Map{"id": 4, "name": "jim"},
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["key1"][0].Id, 1)
		t.Assert(m["key1"][1].Id, 2)
		t.Assert(m["key2"][0].Id, 3)
		t.Assert(m["key2"][1].Id, 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string][]User)(nil)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["key1"][0].Id, 1)
		t.Assert(m["key1"][1].Id, 2)
		t.Assert(m["key2"][0].Id, 3)
		t.Assert(m["key2"][1].Id, 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["key1"][0].Id, 1)
		t.Assert(m["key1"][1].Id, 2)
		t.Assert(m["key2"][0].Id, 3)
		t.Assert(m["key2"][1].Id, 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := (map[string][]*User)(nil)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["key1"][0].Id, 1)
		t.Assert(m["key1"][1].Id, 2)
		t.Assert(m["key2"][0].Id, 3)
		t.Assert(m["key2"][1].Id, 4)
	})
}

func Test_MapToMaps2(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}
	paramsqn.Slqn.MapIntAny{
		100: g.Slice{
			qn.Map{"id": 1, "name": "john"},
			qn.Map{"id": 2, "name": "smith"},
		},qn.Sl
		200: g.Slice{
			qn.Map{"id": 3, "name": "green"},
			qn.Map{"id": 4, "name": "jim"},
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[int][]User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m[100][0].Id, 1)
		t.Assert(m[100][1].Id, 2)
		t.Assert(m[200][0].Id, 3)
		t.Assert(m[200][1].Id, 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[int][]*User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m[100][0].Id, 1)
		t.Assert(m[100][1].Id, 2)
		t.Assert(m[200][0].Id, 3)
		t.Assert(m[200][1].Id, 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["100"][0].Id, 1)
		t.Assert(m["100"][1].Id, 2)
		t.Assert(m["200"][0].Id, 3)
		t.Assert(m["200"][1].Id, 4)
	})
}

func Test_MapToMapsDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	paramsqn.Slqn.MapIntAny{
		100: g.Slice{
			qn.Map{"id": 1, "name": "john"},
			qn.Map{"id": 2, "name": "smith"},
		},qn.Sl
		200: g.Slice{
			qn.Map{"id": 3, "name": "green"},
			qn.Map{"id": 4, "name": "jim"},
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["100"][0].Id, 0)
		t.Assert(m["100"][1].Id, 0)
		t.Assert(m["100"][0].Name, "john")
		t.Assert(m["100"][1].Name, "smith")
		t.Assert(m["200"][0].Id, 0)
		t.Assert(m["200"][1].Id, 0)
		t.Assert(m["200"][0].Name, "green")
		t.Assert(m["200"][1].Name, "jim")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMapsDeep(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["100"][0].Id, 1)
		t.Assert(m["100"][1].Id, 2)
		t.Assert(m["100"][0].Name, "john")
		t.Assert(m["100"][1].Name, "smith")
		t.Assert(m["200"][0].Id, 3)
		t.Assert(m["200"][1].Id, 4)
		t.Assert(m["200"][0].Name, "green")
		t.Assert(m["200"][1].Name, "jim")
	})
}

func Test_MapToMapsDeepWithTag(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids  `json:"ids"`
		Time string
	}
	type User struct {
		Base `json:"base"`
		Name string
	}
	paramsqn.Slqn.MapIntAny{
		100: g.Slice{
			qn.Map{"id": 1, "name": "john"},
			qn.Map{"id": 2, "name": "smith"},
		},qn.Sl
		200: g.Slice{
			qn.Map{"id": 3, "name": "green"},
			qn.Map{"id": 4, "name": "jim"},
		},
	}
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMaps(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["100"][0].Id, 0)
		t.Assert(m["100"][1].Id, 0)
		t.Assert(m["100"][0].Name, "john")
		t.Assert(m["100"][1].Name, "smith")
		t.Assert(m["200"][0].Id, 0)
		t.Assert(m["200"][1].Id, 0)
		t.Assert(m["200"][0].Name, "green")
		t.Assert(m["200"][1].Name, "jim")
	})
	qn_test.C(t, func(t *qn_test.T) {
		m := make(map[string][]*User)
		err := qn_conv.MapToMapsDeep(params, &m)
		t.Assert(err, nil)
		t.Assert(len(m), 2)
		t.Assert(m["100"][0].Id, 1)
		t.Assert(m["100"][1].Id, 2)
		t.Assert(m["100"][0].Name, "john")
		t.Assert(m["100"][1].Name, "smith")
		t.Assert(m["200"][0].Id, 3)
		t.Assert(m["200"][1].Id, 4)
		t.Assert(m["200"][0].Name, "green")
		t.Assert(m["200"][1].Name, "jim")
	})
}
