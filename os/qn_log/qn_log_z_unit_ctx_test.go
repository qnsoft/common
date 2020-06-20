// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_log_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/os/qn_log"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/qn.str"
)

func Test_Ctx(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		w := bytes.NewBuffer(nil)
		l := qn_log.NewWithWriter(w)
		l.SetCtxKeys("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Ctx(ctx).Print(1, 2, 3)
		t.Assert(qn.str.Count(w.String(), "Trace-Id"), 1)
		t.Assert(qn.str.Count(w.String(), "1234567890"), 1)
		t.Assert(qn.str.Count(w.String(), "Span-Id"), 1)
		t.Assert(qn.str.Count(w.String(), "abcdefg"), 1)
		t.Assert(qn.str.Count(w.String(), "1 2 3"), 1)
	})
}

func Test_Ctx_Config(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		w := bytes.NewBuffer(nil)
		l := qn_log.NewWithWriter(w)
		m := map[string]interface{}{
			"CtxKeys": qn.SliceStr{"Trace-Id", "Span-Id", "Test"},
		}
		err := l.SetConfigWithMap(m)
		t.Assert(err, nil)
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.Ctx(ctx).Print(1, 2, 3)
		t.Assert(qn.str.Count(w.String(), "Trace-Id"), 1)
		t.Assert(qn.str.Count(w.String(), "1234567890"), 1)
		t.Assert(qn.str.Count(w.String(), "Span-Id"), 1)
		t.Assert(qn.str.Count(w.String(), "abcdefg"), 1)
		t.Assert(qn.str.Count(w.String(), "1 2 3"), 1)
	})
}
