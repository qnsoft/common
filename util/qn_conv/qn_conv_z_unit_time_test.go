// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_Time(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t1 := "2011-10-10 01:02:03.456"
		t.AssertEQ(qn_conv.qn_time(t1), qn_time.NewFromStr(t1))
		t.AssertEQ(qn_conv.Time(t1), qn_time.NewFromStr(t1).Time)
		t.AssertEQ(qn_conv.Duration(100), 100*time.Nanosecond)
	})
}
