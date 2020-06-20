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

func Test_IsSubDomain(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		main := "goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8080", main), true)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	qn_test.C(t, func(t *qn_test.T) {
		main := "*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org:80", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	qn_test.C(t, func(t *qn_test.T) {
		main := "*.*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	qn_test.C(t, func(t *qn_test.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
}
