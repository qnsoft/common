// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_snotify_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/gtype"
	"github.com/qnsoft/common/os/gfsnotify"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

func TestWatcher_AddOnce(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := gtype.New()
		path := qn_file.TempDir(qn_conv.String(qn_time.TimestampNano()))
		err := qn_file.PutContents(path, "init")
		t.Assert(err, nil)
		defer qn_file.Remove(path)

		time.Sleep(100 * time.Millisecond)
		callback1, err := gfsnotify.AddOnce("mywatch", path, func(event *gfsnotify.Event) {
			value.Set(1)
		})
		t.Assert(err, nil)
		callback2, err := gfsnotify.AddOnce("mywatch", path, func(event *gfsnotify.Event) {
			value.Set(2)
		})
		t.Assert(err, nil)
		t.Assert(callback2, nil)

		err = qn_file.PutContents(path, "1")
		t.Assert(err, nil)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)

		err = gfsnotify.RemoveCallback(callback1.Id)
		t.Assert(err, nil)

		err = qn_file.PutContents(path, "2")
		t.Assert(err, nil)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)
	})
}

func TestWatcher_AddRemove(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		path1 := qn_file.TempDir() + qn_file.Separator + qn_conv.String(qn_time.TimestampNano())
		path2 := qn_file.TempDir() + qn_file.Separator + qn_conv.String(qn_time.TimestampNano()) + "2"
		qn_file.PutContents(path1, "1")
		defer func() {
			qn_file.Remove(path1)
			qn_file.Remove(path2)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
			if event.IsRename() {
				v.Set(3)
				gfsnotify.Exit()
				return
			}
		})
		t.Assert(err, nil)
		t.AssertNE(callback, nil)

		qn_file.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 2)

		qn_file.Rename(path1, path2)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 3)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path1 := qn_file.TempDir() + qn_file.Separator + qn_conv.String(qn_time.TimestampNano())
		qn_file.PutContents(path1, "1")
		defer func() {
			qn_file.Remove(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
			if event.IsRemove() {
				v.Set(4)
				return
			}
		})
		t.Assert(err, nil)
		t.AssertNE(callback, nil)

		qn_file.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 2)

		qn_file.Remove(path1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 4)

		qn_file.PutContents(path1, "1")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 4)
	})
}

func TestWatcher_Callback1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		path1 := qn_file.TempDir(qn_time.TimestampNanoStr())
		qn_file.PutContents(path1, "1")
		defer func() {
			qn_file.Remove(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
		})
		t.Assert(err, nil)
		t.AssertNE(callback, nil)

		qn_file.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 2)

		v.Set(3)
		gfsnotify.RemoveCallback(callback.Id)
		qn_file.PutContents(path1, "3")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.Val(), 3)
	})
}

func TestWatcher_Callback2(t *testing.T) {
	// multiple callbacks
	qn_test.C(t, func(t *qn_test.T) {
		path1 := qn_file.TempDir(qn_time.TimestampNanoStr())
		t.Assert(qn_file.PutContents(path1, "1"), nil)
		defer func() {
			qn_file.Remove(path1)
		}()
		v1 := gtype.NewInt(1)
		v2 := gtype.NewInt(1)
		callback1, err1 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v1.Set(2)
				return
			}
		})
		callback2, err2 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v2.Set(2)
				return
			}
		})
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.AssertNE(callback1, nil)
		t.AssertNE(callback2, nil)

		t.Assert(qn_file.PutContents(path1, "2"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.Val(), 2)
		t.Assert(v2.Val(), 2)

		v1.Set(3)
		v2.Set(3)
		gfsnotify.RemoveCallback(callback1.Id)
		t.Assert(qn_file.PutContents(path1, "3"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.Val(), 3)
		t.Assert(v2.Val(), 2)
	})
}
