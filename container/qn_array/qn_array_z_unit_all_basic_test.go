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

	"github.com/qnsoft/common/util/qn_conv"
	"github.com/qnsoft/common/util/qn_util"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Array_Var(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.Array
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.IntArray
		expect := []int{2, 3, 1}
		array.Append(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.StrArray
		expect := []string{"b", "a"}
		array.Append("b", "a")
		t.Assert(array.Slice(), expect)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.SortedArray
		array.SetComparator(qn_util.ComparatorInt)
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.SortedStrArray
		expect := []string{"a", "b", "c"}
		array.Add("c", "a", "b")
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray_Var(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var array qn_array.SortedIntArray
		expect := []int{1, 2, 3}
		array.Add(2, 3, 1)
		t.Assert(array.Slice(), expect)
	})
}

func Test_IntArray_Unique(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		array := qn_array.NewIntArray()
		array.Append(1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6)
		array.Unique()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedIntArray1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := qn_array.NewSortedIntArray()
		for i := 10; i > -1; i-- {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
	})
}

func Test_SortedIntArray2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		array := qn_array.NewSortedIntArray()
		for i := 0; i <= 10; i++ {
			array.Add(i)
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedStrArray1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array1 := qn_array.NewSortedStrArray()
		array2 := qn_array.NewSortedStrArray(true)
		for i := 10; i > -1; i-- {
			array1.Add(qn_conv.String(i))
			array2.Add(qn_conv.String(i))
		}
		t.Assert(array1.Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})

}

func Test_SortedStrArray2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := qn_array.NewSortedStrArray()
		for i := 0; i <= 10; i++ {
			array.Add(qn_conv.String(i))
		}
		t.Assert(array.Slice(), expect)
		array.Add()
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		array := qn_array.NewSortedArray(func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		})
		for i := 10; i > -1; i-- {
			array.Add(qn_conv.String(i))
		}
		t.Assert(array.Slice(), expect)
	})
}

func Test_SortedArray2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		expect := []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(qn_conv.String(v1), qn_conv.String(v2))
		}
		array := qn_array.NewSortedArray(func1)
		array2 := qn_array.NewSortedArray(func1, true)
		for i := 0; i <= 10; i++ {
			array.Add(qn_conv.String(i))
			array2.Add(qn_conv.String(i))
		}
		t.Assert(array.Slice(), expect)
		t.Assert(array.Add().Slice(), expect)
		t.Assert(array2.Slice(), expect)
	})
}

func TestNewFromCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		a1 := []interface{}{"100", "200", "300", "400", "500", "600"}
		array1 := qn_array.NewFromCopy(a1)
		t.AssertIN(array1.PopRands(2), a1)
		t.Assert(len(array1.PopRands(1)), 1)
		t.Assert(len(array1.PopRands(9)), 3)
	})
}
