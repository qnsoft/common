// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_str_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Trim(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Trim(" 123456\n "), "123456")
		t.Assert(qn.str.Trim("#123456#;", "#;"), "123456")
	})
}

func Test_TrimStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimStr("gogo我爱gogo", "go"), "我爱")
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimStr("啊我爱中国人啊", "啊"), "我爱中国人")
	})
}

func Test_TrimRight(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimRight(" 123456\n "), " 123456")
		t.Assert(qn.str.TrimRight("#123456#;", "#;"), "#123456")
	})
}

func Test_TrimRightStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimRightStr("gogo我爱gogo", "go"), "gogo我爱")
		t.Assert(qn.str.TrimRightStr("gogo我爱gogo", "go我爱gogo"), "go")
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimRightStr("我爱中国人", "人"), "我爱中国")
		t.Assert(qn.str.TrimRightStr("我爱中国人", "爱中国人"), "我")
	})
}

func Test_TrimLeft(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimLeft(" \r123456\n "), "123456\n ")
		t.Assert(qn.str.TrimLeft("#;123456#;", "#;"), "123456#;")
	})
}

func Test_TrimLeftStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimLeftStr("gogo我爱gogo", "go"), "我爱gogo")
		t.Assert(qn.str.TrimLeftStr("gogo我爱gogo", "gogo我爱go"), "go")
	})
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.TrimLeftStr("我爱中国人", "我爱"), "中国人")
		t.Assert(qn.str.TrimLeftStr("我爱中国人", "我爱中国"), "人")
	})
}
