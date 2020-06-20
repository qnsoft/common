// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package cmdenv

import (
	"os"
	"testing"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Get(t *testing.T) {
	os.Args = []string{"--qn.test.value1=111"}
	os.Setenv("GF_TEST_VALUE1", "222")
	os.Setenv("GF_TEST_VALUE2", "333")
	doInit()
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(Get("qn.test.value1").String(), "111")
		t.Assert(Get("qn.test.value2").String(), "333")
		t.Assert(Get("qn.test.value3").String(), "")
		t.Assert(Get("qn.test.value3", 1).String(), "1")
	})
}
