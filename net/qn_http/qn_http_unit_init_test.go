// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/genv"
)

var (
	ports = qn_array.NewIntArray(true)
)

func init() {
	genv.Set("UNDER_TEST", "1")
	for i := 8000; i <= 9000; i++ {
		ports.Append(i)
	}
}
