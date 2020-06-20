// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_set_test

import (
	"fmt"

	"github.com/qnsoft/common/container/qn_set"
	"github.com/qnsoft/common/frame/qn"
)

func ExampleSet_Intersect() {
	s1 := qn_set.NewFrom(qn.Slice{1, 2, 3})
	s2 := qn_set.NewFrom(qn.Slice{4, 5, 6})
	s3 := qn_set.NewFrom(qn.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Diff() {
	s1 := qn_set.NewFrom(qn.Slice{1, 2, 3})
	s2 := qn_set.NewFrom(qn.Slice{4, 5, 6})
	s3 := qn_set.NewFrom(qn.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Union() {
	s1 := qn_set.NewFrom(qn.Slice{1, 2, 3})
	s2 := qn_set.NewFrom(qn.Slice{4, 5, 6})
	s3 := qn_set.NewFrom(qn.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Complement() {
	s1 := qn_set.NewFrom(qn.Slice{1, 2, 3})
	s2 := qn_set.NewFrom(qn.Slice{4, 5, 6})
	s3 := qn_set.NewFrom(qn.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_IsSubsetOf() {
	var s1, s2 qn_set.Set
	s1.Add(qn.Slice{1, 2, 3}...)
	s2.Add(qn.Slice{2, 3}...)
	fmt.Println(s1.IsSubsetOf(&s2))
	fmt.Println(s2.IsSubsetOf(&s1))

	// Output:
	// false
	// true
}

func ExampleSet_AddIfNotExist() {
	var set qn_set.Set
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.Slice())

	// Output:
	// true
	// false
	// [1]
}

func ExampleSet_Pop() {
	var set qn_set.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Pops() {
	var set qn_set.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Join() {
	var set qn_set.Set
	set.Add("a", "b", "c", "d")
	fmt.Println(set.Join(","))

	// May Output:
	// a,b,c,d
}

func ExampleSet_Contains() {
	var set qn_set.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleSet_ContainsI() {
	var set qn_set.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}
