// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_session

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_NewSessionId(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		id1 := NewSessionId()
		id2 := NewSessionId()
		t.AssertNE(id1, id2)
		t.Assert(len(id1), 32)
	})
}
