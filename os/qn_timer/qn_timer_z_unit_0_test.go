// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// Package functions

package qn_timer_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/gtimer"
	"github.com/qnsoft/common/test/gtest"
)

func TestSetTimeout(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.SetTimeout(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestSetInterval(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.SetInterval(200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.Len(), 5)
	})
}

func TestAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.AddEntry(200*time.Millisecond, func() {
			array.Append(1)
		}, false, 2, gtimer.STATUS_READY)
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.AddSingleton(200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(1100 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.AddTimes(200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestDelayAdd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.DelayAdd(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddEntry(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.DelayAddEntry(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		}, false, 2, gtimer.STATUS_READY)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func TestDelayAddSingleton(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.DelayAddSingleton(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
			time.Sleep(10000 * time.Millisecond)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.DelayAddOnce(200*time.Millisecond, 200*time.Millisecond, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})
}

func TestDelayAddTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := qn_array.New(true)
		gtimer.DelayAddTimes(200*time.Millisecond, 200*time.Millisecond, 2, func() {
			array.Append(1)
		})
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 0)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}
