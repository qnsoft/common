// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go

package qn_array_test

import (
	"strings"
	"testing"
	"time"

	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/util/qn_conv"
	"github.com/qnsoft/common/util/qn_util"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/test/qn_test"
)

func TestSortedArray_NewSortedArrayFrom(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "f", "c"}
		a2 := []interface{}{"h", "j", "i", "k"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		func2 := func(v1, v2 interface{}) int {
			return -1
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array2 := qn_array.NewSortedArrayFrom(a2, func2)

		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})

		t.Assert(array2.Len(), 4)
		t.Assert(array2, []interface{}{"k", "i", "j", "h"})
	})
}

func TestNewSortedArrayFromCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "f", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		func2 := func(v1, v2 interface{}) int {
			return -1
		}
		array1 := qn_array.NewSortedArrayFromCopy(a1, func1)
		array2 := qn_array.NewSortedArrayFromCopy(a1, func2)
		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})
		t.Assert(array1.Len(), 3)
		t.Assert(array2, []interface{}{"c", "f", "a"})
	})
}

func TestSortedArray_SetArray(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "f", "c"}
		a2 := []interface{}{"e", "h", "g", "k"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}

		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array1.SetArray(a2)
		t.Assert(array1.Len(), 4)
		t.Assert(array1, []interface{}{"e", "g", "h", "k"})
	})

}

func TestSortedArray_Sort(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "f", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array1.Sort()
		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})
	})

}

func TestSortedArray_Get(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "f", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		v, ok := array1.Get(2)
		t.Assert(v, "f")
		t.Assert(ok, true)

		v, ok = array1.Get(1)
		t.Assert(v, "c")
		t.Assert(ok, true)
	})

}

func TestSortedArray_Remove(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1, ok := array1.Remove(1)
		t.Assert(ok, true)
		t.Assert(qn_conv.String(i1), "b")
		t.Assert(array1.Len(), 3)
		t.Assert(array1.Contains("b"), false)

		v, ok := array1.Remove(-1)
		t.Assert(v, nil)
		t.Assert(ok, false)

		v, ok = array1.Remove(100000)
		t.Assert(v, nil)
		t.Assert(ok, false)

		i2, ok := array1.Remove(0)
		t.Assert(ok, true)
		t.Assert(qn_conv.String(i2), "a")
		t.Assert(array1.Len(), 2)
		t.Assert(array1.Contains("a"), false)

		i3, ok := array1.Remove(1)
		t.Assert(ok, true)
		t.Assert(qn_conv.String(i3), "d")
		t.Assert(array1.Len(), 1)
		t.Assert(array1.Contains("d"), false)
	})

}

func TestSortedArray_PopLeft(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		array1 := qn_array.NewSortedArrayFrom(
			[]interface{}{"a", "d", "c", "b"},
			qn_util.ComparatorString,
		)
		i1, ok := array1.PopLeft()
		t.Assert(ok, true)
		t.Assert(qn_conv.String(i1), "a")
		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"b", "c", "d"})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFrom(qn.Slice{1, 2, 3}, qn_util.ComparatorInt)
		v, ok := array.PopLeft()
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array.Len(), 2)
		v, ok = array.PopLeft()
		t.Assert(v, 2)
		t.Assert(ok, true)
		t.Assert(array.Len(), 1)
		v, ok = array.PopLeft()
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array.Len(), 0)
	})
}

func TestSortedArray_PopRight(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		array1 := qn_array.NewSortedArrayFrom(
			[]interface{}{"a", "d", "c", "b"},
			qn_util.ComparatorString,
		)
		i1, ok := array1.PopRight()
		t.Assert(ok, true)
		t.Assert(qn_conv.String(i1), "d")
		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"a", "b", "c"})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFrom(qn.Slice{1, 2, 3}, qn_util.ComparatorInt)
		v, ok := array.PopRight()
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array.Len(), 2)

		v, ok = array.PopRight()
		t.Assert(v, 2)
		t.Assert(ok, true)
		t.Assert(array.Len(), 1)

		v, ok = array.PopRight()
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array.Len(), 0)
	})
}

func TestSortedArray_PopRand(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1, ok := array1.PopRand()
		t.Assert(ok, true)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.Len(), 3)

	})
}

func TestSortedArray_PopRands(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1 := array1.PopRands(2)
		t.Assert(len(i1), 2)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.Len(), 2)

		i2 := array1.PopRands(3)
		t.Assert(len(i1), 2)
		t.AssertIN(i2, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.Len(), 0)

	})
}

func TestSortedArray_Empty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArray(qn_util.ComparatorInt)
		v, ok := array.PopLeft()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.PopLefts(10), nil)

		v, ok = array.PopRight()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.PopRights(10), nil)

		v, ok = array.PopRand()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.PopRands(10), nil)
	})
}

func TestSortedArray_PopLefts(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1 := array1.PopLefts(2)
		t.Assert(len(i1), 2)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b", "e", "f"})
		t.Assert(array1.Len(), 4)

		i2 := array1.PopLefts(5)
		t.Assert(len(i2), 4)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b", "e", "f"})
		t.Assert(array1.Len(), 0)
	})
}

func TestSortedArray_PopRights(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1 := array1.PopRights(2)
		t.Assert(len(i1), 2)
		t.Assert(i1, []interface{}{"e", "f"})
		t.Assert(array1.Len(), 4)

		i2 := array1.PopRights(10)
		t.Assert(len(i2), 4)
		t.Assert(array1.Len(), 0)
	})
}

func TestSortedArray_Range(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array2 := qn_array.NewSortedArrayFrom(a1, func1, true)
		i1 := array1.Range(2, 5)
		t.Assert(i1, []interface{}{"c", "d", "e"})
		t.Assert(array1.Len(), 6)

		i2 := array1.Range(7, 5)
		t.Assert(len(i2), 0)
		i2 = array1.Range(-1, 2)
		t.Assert(i2, []interface{}{"a", "b"})

		i2 = array1.Range(4, 10)
		t.Assert(len(i2), 2)
		t.Assert(i2, []interface{}{"e", "f"})

		t.Assert(array2.Range(1, 3), []interface{}{"b", "c"})

	})
}

func TestSortedArray_Sum(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		a2 := []interface{}{"1", "2", "3", "b", "e", "f"}
		a3 := []interface{}{"4", "5", "6"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array2 := qn_array.NewSortedArrayFrom(a2, func1)
		array3 := qn_array.NewSortedArrayFrom(a3, func1)
		t.Assert(array1.Sum(), 0)
		t.Assert(array2.Sum(), 6)
		t.Assert(array3.Sum(), 15)

	})
}

func TestSortedArray_Clone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array2 := array1.Clone()
		t.Assert(array1, array2)
		array1.Remove(1)
		t.AssertNE(array1, array2)

	})
}

func TestSortedArray_Clear(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		t.Assert(array1.Len(), 6)
		array1.Clear()
		t.Assert(array1.Len(), 0)

	})
}

func TestSortedArray_Chunk(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1 := array1.Chunk(2)
		t.Assert(len(i1), 3)
		t.Assert(i1[0], []interface{}{"a", "b"})
		t.Assert(i1[2], []interface{}{"e"})

		i1 = array1.Chunk(0)
		t.Assert(len(i1), 0)
	})
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{1, 2, 3, 4, 5}
		array1 := qn_array.NewSortedArrayFrom(a1, qn_util.ComparatorInt)
		chunks := array1.Chunk(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5})
		t.Assert(array1.Chunk(0), nil)
	})
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := qn_array.NewSortedArrayFrom(a1, qn_util.ComparatorInt)
		chunks := array1.Chunk(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []interface{}{1, 2})
		t.Assert(chunks[1], []interface{}{3, 4})
		t.Assert(chunks[2], []interface{}{5, 6})
		t.Assert(array1.Chunk(0), nil)
	})
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := qn_array.NewSortedArrayFrom(a1, qn_util.ComparatorInt)
		chunks := array1.Chunk(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5, 6})
		t.Assert(array1.Chunk(0), nil)
	})
}

func TestSortedArray_SubSlice(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array2 := qn_array.NewSortedArrayFrom(a1, func1, true)
		i1 := array1.SubSlice(2, 3)
		t.Assert(len(i1), 3)
		t.Assert(i1, []interface{}{"c", "d", "e"})

		i1 = array1.SubSlice(2, 6)
		t.Assert(len(i1), 3)
		t.Assert(i1, []interface{}{"c", "d", "e"})

		i1 = array1.SubSlice(7, 2)
		t.Assert(len(i1), 0)

		s1 := array1.SubSlice(1, -2)
		t.Assert(s1, nil)

		s1 = array1.SubSlice(-9, 2)
		t.Assert(s1, nil)
		t.Assert(array2.SubSlice(1, 3), []interface{}{"b", "c", "d"})

	})
}

func TestSortedArray_Rand(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1, ok := array1.Rand()
		t.Assert(ok, true)
		t.AssertIN(i1, []interface{}{"a", "d", "c"})
		t.Assert(array1.Len(), 3)
	})
}

func TestSortedArray_Rands(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		i1 := array1.Rands(2)
		t.AssertIN(i1, []interface{}{"a", "d", "c"})
		t.Assert(len(i1), 2)
		t.Assert(array1.Len(), 3)

		i1 = array1.Rands(4)
		t.Assert(len(i1), 4)
	})
}

func TestSortedArray_Join(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		t.Assert(array1.Join(","), `a,c,d`)
		t.Assert(array1.Join("."), `a.c.d`)
	})

	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{0, 1, `"a"`, `\a`}
		array1 := qn_array.NewSortedArrayFrom(a1, qn_util.ComparatorString)
		t.Assert(array1.Join("."), `"a".0.1.\a`)
	})
}

func TestSortedArray_String(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{0, 1, "a", "b"}
		array1 := qn_array.NewSortedArrayFrom(a1, qn_util.ComparatorString)
		t.Assert(array1.String(), `[0,1,"a","b"]`)
	})
}

func TestSortedArray_CountValues(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		m1 := array1.CountValues()
		t.Assert(len(m1), 3)
		t.Assert(m1["c"], 2)
		t.Assert(m1["a"], 1)

	})
}

func TestSortedArray_SetUnique(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"a", "d", "c", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array1 := qn_array.NewSortedArrayFrom(a1, func1)
		array1.SetUnique(true)
		t.Assert(array1.Len(), 3)
		t.Assert(array1, []interface{}{"a", "c", "d"})
	})
}

func TestSortedArray_LockFunc(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := qn_array.NewSortedArrayFrom(s1, func1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		//go1
		go a1.LockFunc(func(n1 []interface{}) { //读写锁
			time.Sleep(2 * time.Second) //暂停2秒
			n1[2] = "g"
			ch2 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
		})

		//go2
		go func() {
			time.Sleep(100 * time.Millisecond) //故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
			a1.Len()
			ch1 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 //等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertGT(t2-t1, 20) //go1加的读写互斥锁，所go2读的时候被阻塞。
		t.Assert(a1.Contains("g"), true)
	})
}

func TestSortedArray_RLockFunc(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := qn_array.NewSortedArrayFrom(s1, func1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		//go1
		go a1.RLockFunc(func(n1 []interface{}) { //读写锁
			time.Sleep(2 * time.Second) //暂停2秒
			n1[2] = "g"
			ch2 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
		})

		//go2
		go func() {
			time.Sleep(100 * time.Millisecond) //故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
			a1.Len()
			ch1 <- qn_conv.Int64(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 //等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertLT(t2-t1, 20) //go1加的读锁，所go2读的时候不会被阻塞。
		t.Assert(a1.Contains("g"), true)
	})
}

func TestSortedArray_Merge(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		func1 := func(v1, v2 interface{}) int {
			if qn_conv.Int(v1) < qn_conv.Int(v2) {
				return 0
			}
			return 1
		}

		s1 := []interface{}{"a", "b", "c", "d"}
		s2 := []string{"e", "f"}
		i1 := qn_array.NewIntArrayFrom([]int{1, 2, 3})
		i2 := qn_array.NewArrayFrom([]interface{}{3})
		s3 := qn_array.NewStrArrayFrom([]string{"g", "h"})
		s4 := qn_array.NewSortedArrayFrom([]interface{}{4, 5}, func1)
		s5 := qn_array.NewSortedStrArrayFrom(s2)
		s6 := qn_array.NewSortedIntArrayFrom([]int{1, 2, 3})

		a1 := qn_array.NewSortedArrayFrom(s1, func1)

		t.Assert(a1.Merge(s2).Len(), 6)
		t.Assert(a1.Merge(i1).Len(), 9)
		t.Assert(a1.Merge(i2).Len(), 10)
		t.Assert(a1.Merge(s3).Len(), 12)
		t.Assert(a1.Merge(s4).Len(), 14)
		t.Assert(a1.Merge(s5).Len(), 16)
		t.Assert(a1.Merge(s6).Len(), 19)
	})
}

func TestSortedArray_Json(t *testing.T) {
	// array pointer
	qn_test.C(t, func(t *qn_test.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		s2 := []interface{}{"a", "b", "c", "d"}
		a1 := qn_array.NewSortedArrayFrom(s1, qn_util.ComparatorString)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := qn_array.NewSortedArray(qn_util.ComparatorString)
		err1 = json.Unmarshal(b2, &a2)
		t.Assert(a2.Slice(), s2)

		var a3 qn_array.SortedArray
		err := json.Unmarshal(b2, &a3)
		t.Assert(err, nil)
		t.Assert(a3.Slice(), s1)
		t.Assert(a3.Interfaces(), s1)
	})
	// array value
	qn_test.C(t, func(t *qn_test.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		s2 := []interface{}{"a", "b", "c", "d"}
		a1 := *qn_array.NewSortedArrayFrom(s1, qn_util.ComparatorString)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := qn_array.NewSortedArray(qn_util.ComparatorString)
		err1 = json.Unmarshal(b2, &a2)
		t.Assert(a2.Slice(), s2)

		var a3 qn_array.SortedArray
		err := json.Unmarshal(b2, &a3)
		t.Assert(err, nil)
		t.Assert(a3.Slice(), s1)
		t.Assert(a3.Interfaces(), s1)
	})
	// array pointer
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Name   string
			Scores *qn_array.SortedArray
		}
		data := qn.Map{
			"Name":   "john",
			"Scores": []int{99, 100, 98},
		}
		b, err := json.Marshal(data)
		t.Assert(err, nil)

		user := new(User)
		err = json.Unmarshal(b, user)
		t.Assert(err, nil)
		t.Assert(user.Name, data["Name"])
		t.AssertNE(user.Scores, nil)
		t.Assert(user.Scores.Len(), 3)

		v, ok := user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.Assert(v, nil)
		t.Assert(ok, false)
	})
	// array value
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Name   string
			Scores qn_array.SortedArray
		}
		data := qn.Map{
			"Name":   "john",
			"Scores": []int{99, 100, 98},
		}
		b, err := json.Marshal(data)
		t.Assert(err, nil)

		user := new(User)
		err = json.Unmarshal(b, user)
		t.Assert(err, nil)
		t.Assert(user.Name, data["Name"])
		t.AssertNE(user.Scores, nil)
		t.Assert(user.Scores.Len(), 3)

		v, ok := user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.PopLeft()
		t.Assert(v, nil)
		t.Assert(ok, false)
	})
}

func TestSortedArray_Iterator(t *testing.T) {
	slice := qn.Slice{"a", "b", "d", "c"}
	array := qn_array.NewSortedArrayFrom(slice, qn_util.ComparatorString)
	qn_test.C(t, func(t *qn_test.T) {
		array.Iterator(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array.IteratorAsc(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array.IteratorDesc(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		index := 0
		array.Iterator(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		index := 0
		array.IteratorAsc(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		index := 0
		array.IteratorDesc(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestSortedArray_RemoveValue(t *testing.T) {
	slice := qn.Slice{"a", "b", "d", "c"}
	array := qn_array.NewSortedArrayFrom(slice, qn_util.ComparatorString)
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(array.RemoveValue("e"), false)
		t.Assert(array.RemoveValue("b"), true)
		t.Assert(array.RemoveValue("a"), true)
		t.Assert(array.RemoveValue("c"), true)
		t.Assert(array.RemoveValue("f"), false)
	})
}

func TestSortedArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *qn_array.SortedArray
	}
	// JSON
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(qn.Map{
			"name":  "john",
			"array": []byte(`[2,3,1]`),
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.Slice(), qn.Slice{1, 2, 3})
	})
	// Map
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(qn.Map{
			"name":  "john",
			"array": qn.Slice{2, 3, 1},
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.Slice(), qn.Slice{1, 2, 3})
	})
}

func TestSortedArray_FilterNil(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		values := qn.Slice{0, 1, 2, 3, 4, "", qn.Slice{}}
		array := qn_array.NewSortedArrayFromCopy(values, qn_util.ComparatorInt)
		t.Assert(array.FilterNil().Slice(), qn.Slice{0, "", qn.Slice{}, 1, 2, 3, 4})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFromCopy(qn.Slice{nil, 1, 2, 3, 4, nil}, qn_util.ComparatorInt)
		t.Assert(array.FilterNil(), qn.Slice{1, 2, 3, 4})
	})
}

func TestSortedArray_FilterEmpty(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFrom(qn.Slice{0, 1, 2, 3, 4, "", qn.Slice{}}, qn_util.ComparatorInt)
		t.Assert(array.FilterEmpty(), qn.Slice{1, 2, 3, 4})
	})
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFrom(qn.Slice{1, 2, 3, 4}, qn_util.ComparatorInt)
		t.Assert(array.FilterEmpty(), qn.Slice{1, 2, 3, 4})
	})
}

func TestSortedArray_Walk(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		array := qn_array.NewSortedArrayFrom(qn.Slice{"1", "2"}, qn_util.ComparatorString)
		t.Assert(array.Walk(func(value interface{}) interface{} {
			return "key-" + qn_conv.String(value)
		}), qn.Slice{"key-1", "key-2"})
	})
}
