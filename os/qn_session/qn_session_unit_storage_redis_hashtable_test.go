// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_session_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/database/gredis"
	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/os/gsession"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_StorageRedisHashTable(t *testing.T) {
	redis, err := gredis.NewFromStr("127.0.0.1:6379,0")
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(err, nil)
	})

	storage := gsession.NewStorageRedisHashTable(redis)
	manager := gsession.New(time.Second, storage)
	sessionId := ""
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New()
		defer s.Close()
		s.Set("k1", "v1")
		s.Set("k2", "v2")
		s.Sets(g.Map{
			"k3": "v3",
			"k4": "v4",
		})
		t.Assert(s.IsDirty(), true)
		sessionId = s.Id()
	})
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New(sessionId)
		t.Assert(s.Get("k1"), "v1")
		t.Assert(s.Get("k2"), "v2")
		t.Assert(s.Get("k3"), "v3")
		t.Assert(s.Get("k4"), "v4")
		t.Assert(len(s.Map()), 4)
		t.Assert(s.Map()["k1"], "v1")
		t.Assert(s.Map()["k4"], "v4")
		t.Assert(s.Id(), sessionId)
		t.Assert(s.Size(), 4)
		t.Assert(s.Contains("k1"), true)
		t.Assert(s.Contains("k3"), true)
		t.Assert(s.Contains("k5"), false)
		s.Remove("k4")
		t.Assert(s.Size(), 3)
		t.Assert(s.Contains("k3"), true)
		t.Assert(s.Contains("k4"), false)
		s.RemoveAll()
		t.Assert(s.Size(), 0)
		t.Assert(s.Contains("k1"), false)
		t.Assert(s.Contains("k2"), false)
		s.Sets(g.Map{
			"k5": "v5",
			"k6": "v6",
		})
		t.Assert(s.Size(), 2)
		t.Assert(s.Contains("k5"), true)
		t.Assert(s.Contains("k6"), true)
		s.Close()
	})

	time.Sleep(1500 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New(sessionId)
		t.Assert(s.Size(), 0)
		t.Assert(s.Get("k5"), nil)
		t.Assert(s.Get("k6"), nil)
	})
}

func Test_StorageRedisHashTablePrefix(t *testing.T) {
	redis, err := gredis.NewFromStr("127.0.0.1:6379,0")
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(err, nil)
	})

	prefix := "s_"
	storage := gsession.NewStorageRedisHashTable(redis, prefix)
	manager := gsession.New(time.Second, storage)
	sessionId := ""
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New()
		defer s.Close()
		s.Set("k1", "v1")
		s.Set("k2", "v2")
		s.Sets(g.Map{
			"k3": "v3",
			"k4": "v4",
		})
		t.Assert(s.IsDirty(), true)
		sessionId = s.Id()
	})
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New(sessionId)
		t.Assert(s.Get("k1"), "v1")
		t.Assert(s.Get("k2"), "v2")
		t.Assert(s.Get("k3"), "v3")
		t.Assert(s.Get("k4"), "v4")
		t.Assert(len(s.Map()), 4)
		t.Assert(s.Map()["k1"], "v1")
		t.Assert(s.Map()["k4"], "v4")
		t.Assert(s.Id(), sessionId)
		t.Assert(s.Size(), 4)
		t.Assert(s.Contains("k1"), true)
		t.Assert(s.Contains("k3"), true)
		t.Assert(s.Contains("k5"), false)
		s.Remove("k4")
		t.Assert(s.Size(), 3)
		t.Assert(s.Contains("k3"), true)
		t.Assert(s.Contains("k4"), false)
		s.RemoveAll()
		t.Assert(s.Size(), 0)
		t.Assert(s.Contains("k1"), false)
		t.Assert(s.Contains("k2"), false)
		s.Sets(g.Map{
			"k5": "v5",
			"k6": "v6",
		})
		t.Assert(s.Size(), 2)
		t.Assert(s.Contains("k5"), true)
		t.Assert(s.Contains("k6"), true)
		s.Close()
	})

	time.Sleep(1500 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		s := manager.New(sessionId)
		t.Assert(s.Size(), 0)
		t.Assert(s.Get("k5"), nil)
		t.Assert(s.Get("k6"), nil)
	})
}
