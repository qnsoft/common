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

func getStr() string {
	return "z"
}

func Test_IntStrMap_Var(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var m qn_map.IntStrMap
		m.Set(1, "a")

		t.Assert(m.Get(1), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(2, "b"), "b")
		t.Assert(m.SetIfNotExist(2, "b"), false)

		t.Assert(m.SetIfNotExist(3, "c"), true)

		t.Assert(m.Remove(2), "b")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		m_f := qn_map.NewIntStrMap()
		m_f.Set(1, "2")
		m_f.Flip()
		t.Assert(m_f.Map(), map[int]string{2: "1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_IntStrMap_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.Set(1, "a")

		t.Assert(m.Get(1), "a")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet(2, "b"), "b")
		t.Assert(m.SetIfNotExist(2, "b"), false)

		t.Assert(m.SetIfNotExist(3, "c"), true)

		t.Assert(m.Remove(2), "b")
		t.Assert(m.Contains(2), false)

		t.AssertIN(3, m.Keys())
		t.AssertIN(1, m.Keys())
		t.AssertIN("a", m.Values())
		t.AssertIN("c", m.Values())

		//反转之后不成为以下 map,flip 操作只是翻转原 map
		//t.Assert(m.Map(), map[string]int{"a": 1, "c": 3})
		m_f := qn_map.NewIntStrMap()
		m_f.Set(1, "2")
		m_f.Flip()
		t.Assert(m_f.Map(), map[int]string{2: "1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := qn_map.NewIntStrMapFrom(map[int]string{1: "a", 2: "b"})
		t.Assert(m2.Map(), map[int]string{1: "a", 2: "b"})
	})
}

func Test_IntStrMap_Set_Fun(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.GetOrSetFunc(1, getStr)
		m.GetOrSetFuncLock(2, getStr)
		t.Assert(m.Get(1), "z")
		t.Assert(m.Get(2), "z")
		t.Assert(m.SetIfNotExistFunc(1, getStr), false)
		t.Assert(m.SetIfNotExistFunc(3, getStr), true)

		t.Assert(m.SetIfNotExistFuncLock(2, getStr), false)
		t.Assert(m.SetIfNotExistFuncLock(4, getStr), true)
	})
}

func Test_IntStrMap_Batch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.Sets(map[int]string{1: "a", 2: "b", 3: "c"})
		t.Assert(m.Map(), map[int]string{1: "a", 2: "b", 3: "c"})
		m.Removes([]int{1, 2})
		t.Assert(m.Map(), map[int]interface{}{3: "c"})
	})
}
func Test_IntStrMap_Iterator(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := map[int]string{1: "a", 2: "b"}
		m := qn_map.NewIntStrMapFrom(expect)
		m.Iterator(func(k int, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.Iterator(func(k int, v string) bool {
			i++
			return true
		})
		m.Iterator(func(k int, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_IntStrMap_Lock(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := map[int]string{1: "a", 2: "b", 3: "c"}
		m := qn_map.NewIntStrMapFrom(expect)
		m.LockFunc(func(m map[int]string) {
			t.Assert(m, expect)
		})
		m.RLockFunc(func(m map[int]string) {
			t.Assert(m, expect)
		})
	})
}

func Test_IntStrMap_Clone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		//clone 方法是深克隆
		m := qn_map.NewIntStrMapFrom(map[int]string{1: "a", 2: "b", 3: "c"})

		m_clone := m.Clone()
		m.Remove(1)
		//修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove(2)
		//修改clone map,原 map 不影响
		t.AssertIN(2, m.Keys())
	})
}
func Test_IntStrMap_Merge(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := qn_map.NewIntStrMap()
		m2 := qn_map.NewIntStrMap()
		m1.Set(1, "a")
		m2.Set(2, "b")
		m1.Merge(m2)
		t.Assert(m1.Map(), map[int]string{1: "a", 2: "b"})
	})
}

func Test_IntStrMap_Map(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.Set(1, "0")
		m.Set(2, "2")
		t.Assert(m.Get(1), "0")
		t.Assert(m.Get(2), "2")
		data := m.Map()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.Get(3), "3")
		m.Set(4, "4")
		t.Assert(data[4], "4")
	})
}

func Test_IntStrMap_MapCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.Set(1, "0")
		m.Set(2, "2")
		t.Assert(m.Get(1), "0")
		t.Assert(m.Get(2), "2")
		data := m.MapCopy()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.Get(3), "")
		m.Set(4, "4")
		t.Assert(data[4], "")
	})
}

func Test_IntStrMap_FilterEmpty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMap()
		m.Set(1, "")
		m.Set(2, "2")
		t.Assert(m.Size(), 2)
		t.Assert(m.Get(2), "2")
		m.FilterEmpty()
		t.Assert(m.Size(), 1)
		t.Assert(m.Get(2), "2")
	})
}

func Test_IntStrMap_Json(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		data := qn.MapIntStr{
			1: "v1",
			2: "v2",
		}
		m1 := qn_map.NewIntStrMapFrom(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	qn_test.C(t, func(t *qn_test.T) {
		data := qn.MapIntStr{
			1: "v1",
			2: "v2",
		}
		b, err := json.Marshal(data)
		t.Assert(err, nil)

		m := qn_map.NewIntStrMap()
		err = json.Unmarshal(b, m)
		t.Assert(err, nil)
		t.Assert(m.Get(1), data[1])
		t.Assert(m.Get(2), data[2])
	})
}

func Test_IntStrMap_Pop(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMapFrom(qn.MapIntStr{
			1: "v1",
			2: "v2",
		})
		t.Assert(m.Size(), 2)

		k1, v1 := m.Pop()
		t.AssertIN(k1, qn.Slice{1, 2})
		t.AssertIN(v1, qn.Slice{"v1", "v2"})
		t.Assert(m.Size(), 1)
		k2, v2 := m.Pop()
		t.AssertIN(k2, qn.Slice{1, 2})
		t.AssertIN(v2, qn.Slice{"v1", "v2"})
		t.Assert(m.Size(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)
	})
}

func Test_IntStrMap_Pops(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_map.NewIntStrMapFrom(qn.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})
		t.Assert(m.Size(), 3)

		kArray := qn_array.New()
		vArray := qn_array.New()
		for k, v := range m.Pops(1) {
			t.AssertIN(k, qn.Slice{1, 2, 3})
			t.AssertIN(v, qn.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 2)
		for k, v := range m.Pops(2) {
			t.AssertIN(k, qn.Slice{1, 2, 3})
			t.AssertIN(v, qn.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.Size(), 0)

		t.Assert(kArray.Unique().Len(), 3)
		t.Assert(vArray.Unique().Len(), 3)
	})
}

func TestIntStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *qn_map.IntStrMap
	}
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"1":"v1","2":"v2"}`),
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get(1), "v1")
		t.Assert(v.Map.Get(2), "v2")
	})
	// Map
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"map": qn.MapIntAny{
				1: "v1",
				2: "v2",
			},
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get(1), "v1")
		t.Assert(v.Map.Get(2), "v2")
	})
}
