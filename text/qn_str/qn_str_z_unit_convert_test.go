// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_str_test

import (
	"testing"

	"github.com/gogf/gf/text/gstr"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_OctStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(gstr.OctStr(`\346\200\241`), "怡")
	})
}
