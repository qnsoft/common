// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qn_queue_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/gqueue"
	"github.com/qnsoft/common/test/qn_test"
)

func TestQueue_Len(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		max := 100
		for n := 10; n < max; n++ {
			q1 := gqueue.New(max)
			for i := 0; i < max; i++ {
				q1.Push(i)
			}
			t.Assert(q1.Len(), max)
			t.Assert(q1.Size(), max)
		}
	})
}

func TestQueue_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		q := gqueue.New()
		for i := 0; i < 100; i++ {
			q.Push(i)
		}
		t.Assert(q.Pop(), 0)
		t.Assert(q.Pop(), 1)
	})
}

func TestQueue_Pop(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		q1 := gqueue.New()
		q1.Push(1)
		q1.Push(2)
		q1.Push(3)
		q1.Push(4)
		i1 := q1.Pop()
		t.Assert(i1, 1)
	})
}

func TestQueue_Close(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		q1 := gqueue.New()
		q1.Push(1)
		q1.Push(2)
		time.Sleep(time.Millisecond)
		t.Assert(q1.Len(), 2)
		q1.Close()
	})
}
