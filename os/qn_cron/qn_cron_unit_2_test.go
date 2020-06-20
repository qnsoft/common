// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cron_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/gcron"
	"github.com/qnsoft/common/os/qn_log"
	"github.com/qnsoft/common/test/qn_test"
)

func TestCron_Entry_Operations(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		cron := gcron.New()
		array := qn_array.New(true)
		cron.DelayAddTimes(500*time.Millisecond, "* * * * * *", 2, func() {
			qn_log.Println("add times")
			array.Append(1)
		})
		t.Assert(cron.Size(), 0)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(3000 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 0)
	})

	qn_test.C(t, func(t *qn_test.T) {
		cron := gcron.New()
		array := qn_array.New(true)
		entry, err1 := cron.Add("* * * * * *", func() {
			qn_log.Println("add")
			array.Append(1)
		})
		t.Assert(err1, nil)
		t.Assert(array.Len(), 0)
		t.Assert(cron.Size(), 1)
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
		entry.Stop()
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		t.Assert(cron.Size(), 1)
		entry.Start()
		qn_log.Println("start")
		time.Sleep(1200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		t.Assert(cron.Size(), 1)
		entry.Close()
		time.Sleep(1200 * time.Millisecond)
		t.Assert(cron.Size(), 0)
	})
}
