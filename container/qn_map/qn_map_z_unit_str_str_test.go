// Copyright 2017-2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_map_test

import (
	"testing"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/container/qn_map"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_StrStrMap_Var(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var m qn_map.StrStrMap
		m.Set("a", "a")

		t.Assert(m.Get("a"), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "b"), "b")
		t.Assert(m.SetIfNotExist("b", "b"), false)

		t.Assert(m.SetIfNotExist("c", "c"), true)

		t.Assert(m.Remove("b"), "b")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[string]string{"a": "a", "c": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_StrStrMap_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()
		m.Set("a", "a")

		t.Assert(m.Get("a"), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("b", "b"), "b")
		t.Assert(m.SetIfNotExist("b", "b"), false)

		t.Assert(m.SetIfNotExist("c", "c"), true)

		t.Assert(m.Remove("b"), "b")
		t.Assert(m.Contains("b"), false)

		t.AssertIN("c", m.Keys())
		t.AssertIN("a", m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m.Flip()

		t.Assert(m.Map(), map[string]string{"a": "a", "c": "c"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := qn_map.NewStrStrMapFrom(map[string]string{"a": "a", "b": "b"})
		t.Assert(m2.Map(), map[string]string{"a": "a", "b": "b"})
	})
}

func Test_StrStrMap_Set_Fun(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()

		m.GetOrSetFunc("a", getStr)
		m.GetOrSetFuncLock("b", getStr)
		t.Assert(m.Get("a"), "z")
		t.Assert(m.Get("b"), "z")
		t.Assert(m.SetIfNotExistFunc("a", getStr), false)
		t.Assert(m.SetIfNotExistFunc("c", getStr), true)

		t.Assert(m.SetIfNotExistFuncLock("b", getStr), false)
		t.Assert(m.SetIfNotExistFuncLock("d", getStr), true)
	})
}

func Test_StrStrMap_Batch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()

		m.Sets(map[string]string{"a": "a", "b": "b", "c": "c"})
		t.Assert(m.Map(), map[string]string{"a": "a", "b": "b", "c": "c"})
		m.Removes([]string{"a", "b"})
		t.Assert(m.Map(), map[string]string{"c": "c"})
	})
}
func Test_StrStrMap_Iterator(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := map[string]string{"a": "a", "b": "b"}
		m := qn_map.NewStrStrMapFrom(expect)
		m.Iterator(func(k string, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k string, v string) bool {
			i++
			return true
		})
		m.Iterator(func(k string, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_StrStrMap_Lock(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := map[string]string{"a": "a", "b": "b"}

		m := qn_map.NewStrStrMapFrom(expect)
		m.LockFunc(func(m map[string]string) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[string]string) {
			t.Assert(m, expect)
		})
	})
}
func Test_StrStrMap_Clone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		//clone 方法是深克隆
		m := qn_map.NewStrStrMapFrom(map[string]string{"a": "a", "b": "b", "c": "c"})

		m_clone := m.Clone()
		m.Remove("a")
		//修改原 map,clone 后的 map 不影响
		t.AssertIN("a", m_clone.Keys())

		m_clone.Remove("b")
		//修改clone map,原 map 不影响
		t.AssertIN("b", m.Keys())
	})
}
func Test_StrStrMap_Merge(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn_map.NewStrStrMap()
		m2 := qn_map.NewStrStrMap()
		m1.Set("a", "a")
		m2.Set("b", "b")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[string]string{"a": "a", "b": "b"})
	})
}

func Test_StrStrMap_Map(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()
		m.Set("1", "1")
		m.Set("2", "2")
		t.Assert(m.Get("1"), "1")
		t.Assert(m.Get("2"), "2")
		data := m.Map()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.Get("3"), "3")
		m.Set("4", "4")
		t.Assert(data["4"], "4")
	})
}

func Test_StrStrMap_MapCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()
		m.Set("1", "1")
		m.Set("2", "2")
		t.Assert(m.Get("1"), "1")
		t.Assert(m.Get("2"), "2")
		data := m.MapCopy()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.Get("3"), "")
		m.Set("4", "4")
		t.Assert(data["4"], "")
	})
}

func Test_StrStrMap_FilterEmpty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMap()
		m.Set("1", "")
		m.Set("2", "2")
		t.Assert(m.Size(), 2)
		t.Assert(m.Get("1"), "")
		t.Assert(m.Get("2"), "2")
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get("2"), "2")
	})
}

func Test_StrStrMap_Json(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		data := qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := qn_map.NewStrStrMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	qn_test.C(t, func(t *qn_test.T) {
		data := qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.Assert(err, nil)

		m := qn_map.NewStrStrMap()
		err = json.Unmarshal(b, m)
		t.Assert(err, nil)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
	qn_test.C(t, func(t *qn_test.T) {
		data := qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.Assert(err, nil)

		var m qn_map.StrStrMap
		err = json.Unmarshal(b, &m)
		t.Assert(err, nil)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func Test_StrStrMap_Pop(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMapFrom(qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.Size(), 2)

		k1, v1 := m.Pop()
		t.AssertIN(k1, qn.Slice{"k1", "k2"})
		t.AssertIN(v1, qn.Slice{"v1", "v2"})
		t.Assert(m.Size(), 1)
		k2, v2 := m.Pop()
		t.AssertIN(k2, qn.Slice{"k1", "k2"})
		t.AssertIN(v2, qn.Slice{"v1", "v2"})
		t.Assert(m.Size(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)
	})
}

func Test_StrStrMap_Pops(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewStrStrMapFrom(qn.MapStrStr{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
		t.Assert(m.Size(), 3)

		kArray := qn_array.New()
		vArray := qn_array.New()
		for k, v := range m.Pops(1) {
			t.AssertIN(k, qn.Slice{"k1", "k2", "k3"})
			t.AssertIN(v, qn.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 2)
		for k, v := range m.Pops(2) {
			t.AssertIN(k, qn.Slice{"k1", "k2", "k3"})
			t.AssertIN(v, qn.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 0)

		t.Assert(kArray.Unique().Len(), 3)
		t.Assert(vArray.Unique().Len(), 3)
	})
}

func TestStrStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *qn_map.StrStrMap
	}
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"k1":"v1","k2":"v2"}`),
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
	})
	// Map
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"map": qn.Map{
				"k1": "v1",
				"k2": "v2",
			},
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
	})
}
