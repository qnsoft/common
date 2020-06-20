// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qn_cache_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_set"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/os/grpool"
	"github.com/qnsoft/common/os/qn_cache"
	"github.com/qnsoft/common/test/qn_test"
)

func TestCache_qn_cache_Set(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_cache.Set(1, 11, 0)
		defer qn_cache.Removes(qn.Slice{1, 2, 3})
		t.Assert(qn_cache.Get(1), 11)
		t.Assert(qn_cache.Contains(1), true)
	})
}

func TestCache_Set(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_cache.New()
		defer c.Close()
		c.Set(1, 11, 0)
		t.Assert(c.Get(1), 11)
		t.Assert(c.Contains(1), true)
	})
}

func TestCache_GetVar(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := qn_cache.New()
		defer c.Close()
		c.Set(1, 11, 0)
		t.Assert(c.Get(1), 11)
		t.Assert(c.Contains(1), true)
		t.Assert(c.GetVar(1).Int(), 11)
		t.Assert(c.GetVar(2).Int(), 0)
		t.Assert(c.GetVar(2).IsNil(), true)
		t.Assert(c.GetVar(2).IsEmpty(), true)
	})
}

func TestCache_Set_Expire(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.Set(2, 22, 100*time.Millisecond)
		t.Assert(cache.Get(2), 22)
		time.Sleep(200 * time.Millisecond)
		t.Assert(cache.Get(2), nil)
		time.Sleep(3 * time.Second)
		t.Assert(cache.Size(), 0)
		cache.Close()
	})

	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.Set(1, 11, 100*time.Millisecond)
		t.Assert(cache.Get(1), 11)
		time.Sleep(200 * time.Millisecond)
		t.Assert(cache.Get(1), nil)
	})
}

func TestCache_Keys_Values(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		for i := 0; i < 10; i++ {
			cache.Set(i, i*10, 0)
		}
		t.Assert(len(cache.Keys()), 10)
		t.Assert(len(cache.Values()), 10)
		t.AssertIN(0, cache.Keys())
		t.AssertIN(90, cache.Values())
	})
}

func TestCache_LRU(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New(2)
		for i := 0; i < 10; i++ {
			cache.Set(i, i, 0)
		}
		t.Assert(cache.Size(), 10)
		t.Assert(cache.Get(6), 6)
		time.Sleep(4 * time.Second)
		t.Assert(cache.Size(), 2)
		t.Assert(cache.Get(6), 6)
		t.Assert(cache.Get(1), nil)
		cache.Close()
	})
}

func TestCache_LRU_expire(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New(2)
		cache.Set(1, nil, 1000)
		t.Assert(cache.Size(), 1)
		t.Assert(cache.Get(1), nil)
	})
}

func TestCache_SetIfNotExist(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.SetIfNotExist(1, 11, 0)
		t.Assert(cache.Get(1), 11)
		cache.SetIfNotExist(1, 22, 0)
		t.Assert(cache.Get(1), 11)
		cache.SetIfNotExist(2, 22, 0)
		t.Assert(cache.Get(2), 22)

		qn_cache.Removes(qn.Slice{1, 2, 3})
		qn_cache.SetIfNotExist(1, 11, 0)
		t.Assert(qn_cache.Get(1), 11)
		qn_cache.SetIfNotExist(1, 22, 0)
		t.Assert(qn_cache.Get(1), 11)
	})
}

func TestCache_Sets(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.Sets(qn.MapAnyAny{1: 11, 2: 22}, 0)
		t.Assert(cache.Get(1), 11)

		qn_cache.Removes(qn.Slice{1, 2, 3})
		qn_cache.Sets(qn.MapAnyAny{1: 11, 2: 22}, 0)
		t.Assert(qn_cache.Get(1), 11)
	})
}

func TestCache_GetOrSet(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.GetOrSet(1, 11, 0)
		t.Assert(cache.Get(1), 11)
		cache.GetOrSet(1, 111, 0)
		t.Assert(cache.Get(1), 11)

		qn_cache.Removes(qn.Slice{1, 2, 3})
		qn_cache.GetOrSet(1, 11, 0)
		t.Assert(qn_cache.Get(1), 11)
		qn_cache.GetOrSet(1, 111, 0)
		t.Assert(qn_cache.Get(1), 11)
	})
}

func TestCache_GetOrSetFunc(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.GetOrSetFunc(1, func() interface{} {
			return 11
		}, 0)
		t.Assert(cache.Get(1), 11)
		cache.GetOrSetFunc(1, func() interface{} {
			return 111
		}, 0)
		t.Assert(cache.Get(1), 11)

		qn_cache.Removes(qn.Slice{1, 2, 3})
		qn_cache.GetOrSetFunc(1, func() interface{} {
			return 11
		}, 0)
		t.Assert(qn_cache.Get(1), 11)
		qn_cache.GetOrSetFunc(1, func() interface{} {
			return 111
		}, 0)
		t.Assert(qn_cache.Get(1), 11)
	})
}

func TestCache_GetOrSetFuncLock(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.GetOrSetFuncLock(1, func() interface{} {
			return 11
		}, 0)
		t.Assert(cache.Get(1), 11)
		cache.GetOrSetFuncLock(1, func() interface{} {
			return 111
		}, 0)
		t.Assert(cache.Get(1), 11)

		qn_cache.Removes(qn.Slice{1, 2, 3})
		qn_cache.GetOrSetFuncLock(1, func() interface{} {
			return 11
		}, 0)
		t.Assert(qn_cache.Get(1), 11)
		qn_cache.GetOrSetFuncLock(1, func() interface{} {
			return 111
		}, 0)
		t.Assert(qn_cache.Get(1), 11)
	})
}

func TestCache_Clear(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		cache.Sets(qn.MapAnyAny{1: 11, 2: 22}, 0)
		cache.Clear()
		t.Assert(cache.Size(), 0)
	})
}

func TestCache_SetConcurrency(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cache := qn_cache.New()
		pool := grpool.New(4)
		go func() {
			for {
				pool.Add(func() {
					cache.SetIfNotExist(1, 11, 10)
				})
			}
		}()
		select {
		case <-time.After(2 * time.Second):
			//t.Log("first part end")
		}

		go func() {
			for {
				pool.Add(func() {
					cache.SetIfNotExist(1, nil, 10)
				})
			}
		}()
		select {
		case <-time.After(2 * time.Second):
			//t.Log("second part end")
		}
	})
}

func TestCache_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		{
			cache := qn_cache.New()
			cache.Sets(qn.MapAnyAny{1: 11, 2: 22}, 0)
			t.Assert(cache.Contains(1), true)
			t.Assert(cache.Get(1), 11)
			data := cache.Data()
			t.Assert(data[1], 11)
			t.Assert(data[2], 22)
			t.Assert(data[3], nil)
			t.Assert(cache.Size(), 2)
			keys := cache.Keys()
			t.Assert(qn_set.NewFrom(qn.Slice{1, 2}).Equal(qn_set.NewFrom(keys)), true)
			keyStrs := cache.KeyStrings()
			t.Assert(qn_set.NewFrom(qn.Slice{"1", "2"}).Equal(qn_set.NewFrom(keyStrs)), true)
			values := cache.Values()
			t.Assert(qn_set.NewFrom(qn.Slice{11, 22}).Equal(qn_set.NewFrom(values)), true)
			removeData1 := cache.Remove(1)
			t.Assert(removeData1, 11)
			t.Assert(cache.Size(), 1)
			cache.Removes(qn.Slice{2})
			t.Assert(cache.Size(), 0)
		}

		qn_cache.Removes(qn.Slice{1, 2, 3})
		{
			qn_cache.Sets(qn.MapAnyAny{1: 11, 2: 22}, 0)
			t.Assert(qn_cache.Contains(1), true)
			t.Assert(qn_cache.Get(1), 11)
			data := qn_cache.Data()
			t.Assert(data[1], 11)
			t.Assert(data[2], 22)
			t.Assert(data[3], nil)
			t.Assert(qn_cache.Size(), 2)
			keys := qn_cache.Keys()
			t.Assert(qn_set.NewFrom(qn.Slice{1, 2}).Equal(qn_set.NewFrom(keys)), true)
			keyStrs := qn_cache.KeyStrings()
			t.Assert(qn_set.NewFrom(qn.Slice{"1", "2"}).Equal(qn_set.NewFrom(keyStrs)), true)
			values := qn_cache.Values()
			t.Assert(qn_set.NewFrom(qn.Slice{11, 22}).Equal(qn_set.NewFrom(values)), true)
			removeData1 := qn_cache.Remove(1)
			t.Assert(removeData1, 11)
			t.Assert(qn_cache.Size(), 1)
			qn_cache.Removes(qn.Slice{2})
			t.Assert(qn_cache.Size(), 0)
		}
	})
}
