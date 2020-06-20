// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_mlock_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/gmlock"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Locker_RLock(t *testing.T) {
	//RLock before Lock
	qn_test.C(t, func(t *qn_test.T) {
		key := "testRLockBeforeLock"
		array := qn_array.New(true)
		go func() {
			gmlock.RLock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.RUnlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.Lock(key)
			array.Append(1)
			gmlock.Unlock(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	//Lock before RLock
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeRLock"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLock(key)
			array.Append(1)
			gmlock.RUnlock(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	//Lock before RLocks
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeRLocks"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.RUnlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.RUnlock(key)
		}()
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_Locker_TryRLock(t *testing.T) {
	//Lock before TryRLock
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeTryRLock"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if gmlock.TryRLock(key) {
				array.Append(1)
				gmlock.RUnlock(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})

	//Lock before TryRLocks
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeTryRLocks"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if gmlock.TryRLock(key) {
				array.Append(1)
				gmlock.RUnlock(key)
			}
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			if gmlock.TryRLock(key) {
				array.Append(1)
				gmlock.RUnlock(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_Locker_RLockFunc(t *testing.T) {
	//RLockFunc before Lock
	qn_test.C(t, func(t *qn_test.T) {
		key := "testRLockFuncBeforeLock"
		array := qn_array.New(true)
		go func() {
			gmlock.RLockFunc(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.Lock(key)
			array.Append(1)
			gmlock.Unlock(key)
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	//Lock before RLockFunc
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeRLockFunc"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLockFunc(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	//Lock before RLockFuncs
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeRLockFuncs"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLockFunc(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.RLockFunc(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_Locker_TryRLockFunc(t *testing.T) {
	//Lock before TryRLockFunc
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeTryRLockFunc"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.TryRLockFunc(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})

	//Lock before TryRLockFuncs
	qn_test.C(t, func(t *qn_test.T) {
		key := "testLockBeforeTryRLockFuncs"
		array := qn_array.New(true)
		go func() {
			gmlock.Lock(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.Unlock(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.TryRLockFunc(key, func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			gmlock.TryRLockFunc(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}
