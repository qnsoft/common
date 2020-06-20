// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_set_test

import (
	"fmt"

	"github.com/qnsoft/common/container/qn_set"
)

func ExampleIntSet_Contains() {
	var set qn_set.IntSet
	set.Add(1)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(2))

	// Output:
	// true
	// false
}
