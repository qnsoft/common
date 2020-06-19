// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_udp_test

import (
	"github.com/qnsoft/common/container/qn_array"
)

var (
	ports = qn_array.NewIntArray(true)
)

func init() {
	for i := 9000; i <= 10000; i++ {
		ports.Append(i)
	}
}
