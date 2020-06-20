// Copyright 2017-2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_tree_test

import (
	"fmt"
	"testing"

	"github.com/qnsoft/common/container/qn_tree"
	"github.com/qnsoft/common/container/qn_var"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_util"
)

func Test_AVLTree_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		m.Set("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)

		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Sets(map[interface{}]interface{}{"key3": "val3", "key1": "val1"})

		m.Flip()
		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Flip(qn_util.ComparatorString)
		t.Assert(m.Map(), map[interface{}]interface{}{"key3": "val3", "key1": "val1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := qn_tree.NewAVLTreeFrom(qn_util.ComparatorString, map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}
func Test_AVLTree_Set_Fun(t *testing.T) {
	//GetOrSetFunc lock or unlock
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		t.Assert(m.GetOrSetFunc("fun", getValue), 3)
		t.Assert(m.GetOrSetFunc("fun", getValue), 3)
		t.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		t.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
	})
	//SetIfNotExistFunc lock or unlock
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), true)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), true)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
	})

}

func Test_AVLTree_Get_Set_Var(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		t.AssertEQ(m.SetIfNotExist("key1", "val1"), true)
		t.AssertEQ(m.SetIfNotExist("key1", "val1"), false)
		t.AssertEQ(m.GetVarOrSet("key1", "val1"), qn_var.New("val1", true))
		t.AssertEQ(m.GetVar("key1"), qn_var.New("val1", true))
	})

	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		t.AssertEQ(m.GetVarOrSetFunc("fun", getValue), qn_var.New(3, true))
		t.AssertEQ(m.GetVarOrSetFunc("fun", getValue), qn_var.New(3, true))
		t.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), qn_var.New(3, true))
		t.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), qn_var.New(3, true))
	})
}

func Test_AVLTree_Batch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTree(qn_util.ComparatorString)
		m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.Removes([]interface{}{"key1", 1})
		t.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_AVLTree_Iterator(t *testing.T) {

	keys := []string{"1", "key1", "key2", "key3", "key4"}
	keyLen := len(keys)
	index := 0

	expect := map[interface{}]interface{}{"key4": "val4", 1: 1, "key1": "val1", "key2": "val2", "key3": "val3"}

	m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorString, expect)

	qn_test.C(t, func(t *qn_test.T) {
		m.Iterator(func(k interface{}, v interface{}) bool {
			t.Assert(k, keys[index])
			index++
			t.Assert(expect[k], v)
			return true
		})

		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			index--
			t.Assert(k, keys[index])
			t.Assert(expect[k], v)
			return true
		})
	})

	m.Print()
	// 断言返回值对遍历控制
	qn_test.C(t, func(t *qn_test.T) {
		i := 0
		j := 0
		m.Iterator(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.Iterator(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, keyLen)
		t.Assert(j, 1)
	})

	qn_test.C(t, func(t *qn_test.T) {
		i := 0
		j := 0
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, keyLen)
		t.Assert(j, 1)
	})

}

func Test_AVLTree_IteratorFrom(t *testing.T) {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 10; i++ {
		m[i] = i * 10
	}
	tree := qn_tree.NewAVLTreeFrom(qn_util.ComparatorInt, m)

	qn_test.C(t, func(t *qn_test.T) {
		n := 5
		tree.IteratorFrom(5, true, func(key, value interface{}) bool {
			t.Assert(n, key)
			t.Assert(n*10, value)
			n++
			return true
		})

		i := 5
		tree.IteratorAscFrom(5, true, func(key, value interface{}) bool {
			t.Assert(i, key)
			t.Assert(i*10, value)
			i++
			return true
		})

		j := 5
		tree.IteratorDescFrom(5, true, func(key, value interface{}) bool {
			t.Assert(j, key)
			t.Assert(j*10, value)
			j--
			return true
		})
	})
}

func Test_AVLTree_Clone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		//clone 方法是深克隆
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorString, map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.Clone()
		m.Remove(1)
		//修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove("key1")
		//修改clone map,原 map 不影响
		t.AssertIN("key1", m.Keys())
	})
}

func Test_AVLTree_LRNode(t *testing.T) {
	expect := map[interface{}]interface{}{"key4": "val4", "key1": "val1", "key2": "val2", "key3": "val3"}
	//safe
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorString, expect)
		t.Assert(m.Left().Key, "key1")
		t.Assert(m.Right().Key, "key4")
	})
	//unsafe
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorString, expect, true)
		t.Assert(m.Left().Key, "key1")
		t.Assert(m.Right().Key, "key4")
	})
}

func Test_AVLTree_CeilingFloor(t *testing.T) {
	expect := map[interface{}]interface{}{
		20: "val20",
		6:  "val6",
		10: "val10",
		12: "val12",
		1:  "val1",
		15: "val15",
		19: "val19",
		8:  "val8",
		4:  "val4"}
	//found and eq
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorInt, expect)
		c, cf := m.Ceiling(8)
		t.Assert(cf, true)
		t.Assert(c.Value, "val8")
		f, ff := m.Floor(20)
		t.Assert(ff, true)
		t.Assert(f.Value, "val20")
	})
	//found and neq
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorInt, expect)
		c, cf := m.Ceiling(9)
		t.Assert(cf, true)
		t.Assert(c.Value, "val10")
		f, ff := m.Floor(5)
		t.Assert(ff, true)
		t.Assert(f.Value, "val4")
	})
	//nofound
	qn_test.C(t, func(t *qn_test.T) {
		m := qn_tree.NewAVLTreeFrom(qn_util.ComparatorInt, expect)
		c, cf := m.Ceiling(21)
		t.Assert(cf, false)
		t.Assert(c, nil)
		f, ff := m.Floor(-1)
		t.Assert(ff, false)
		t.Assert(f, nil)
	})
}

func Test_AVLTree_Remove(t *testing.T) {
	m := qn_tree.NewAVLTree(qn_util.ComparatorInt)
	for i := 1; i <= 50; i++ {
		m.Set(i, fmt.Sprintf("val%d", i))
	}
	expect := m.Map()
	qn_test.C(t, func(t *qn_test.T) {
		for k, v := range expect {
			m1 := m.Clone()
			t.Assert(m1.Remove(k), v)
			t.Assert(m1.Remove(k), nil)
		}
	})
}
