// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// Entry Operations

package qn_timer_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/qn_timer"
	"github.com/qnsoft/common/test/qn_test"
)

func TestEntry_Start_Stop_Close(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timer := New()
		array := qn_array.New(true)
		entry := timer.Add(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
		entry.Stop()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
		entry.Start()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)
		entry.Close()
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 2)

		t.Assert(entry.Status(), qn_timer.STATUS_CLOSED)
	})
}

func TestEntry_Singleton(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timer := New()
		array := qn_array.New(true)
		entry := timer.Add(200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10 * time.Second)
		})
		t.Assert(entry.IsSingleton(), false)
		entry.SetSingleton(true)
		t.Assert(entry.IsSingleton(), true)
		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)

		time.Sleep(250 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestEntry_SetTimes(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timer := New()
		array := qn_array.New(true)
		entry := timer.Add(200*time.Millisecond, func() {
			array.Append(1)
		})
		entry.SetTimes(2)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestEntry_Run(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timer := New()
		array := qn_array.New(true)
		entry := timer.Add(1000*time.Millisecond, func() {
			array.Append(1)
		})
		entry.Run()
		t.Assert(array.Len(), 1)
	})
}
