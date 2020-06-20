// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_uid_test

import (
	"testing"

	"github.com/qnsoft/common/container/gset"
	"github.com/qnsoft/common/util/guid"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_S(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		set := gset.NewStrSet()
		for i := 0; i < 1000000; i++ {
			s := guid.S()
			t.Assert(set.AddIfNotExist(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(len(guid.S([]byte("123"))), 32)
	})
}
