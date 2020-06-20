// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_util_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_util"
)

func Test_ComparatorString(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorString(1, 1), 0)
		t.Assert(qn_util.ComparatorString(1, 2), -1)
		t.Assert(qn_util.ComparatorString(2, 1), 1)
	})
}

func Test_ComparatorInt(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorInt(1, 1), 0)
		t.Assert(qn_util.ComparatorInt(1, 2), -1)
		t.Assert(qn_util.ComparatorInt(2, 1), 1)
	})
}

func Test_ComparatorInt8(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorInt8(1, 1), 0)
		t.Assert(qn_util.ComparatorInt8(1, 2), -1)
		t.Assert(qn_util.ComparatorInt8(2, 1), 1)
	})
}

func Test_ComparatorInt16(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorInt16(1, 1), 0)
		t.Assert(qn_util.ComparatorInt16(1, 2), -1)
		t.Assert(qn_util.ComparatorInt16(2, 1), 1)
	})
}

func Test_ComparatorInt32(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorInt32(1, 1), 0)
		t.Assert(qn_util.ComparatorInt32(1, 2), -1)
		t.Assert(qn_util.ComparatorInt32(2, 1), 1)
	})
}

func Test_ComparatorInt64(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorInt64(1, 1), 0)
		t.Assert(qn_util.ComparatorInt64(1, 2), -1)
		t.Assert(qn_util.ComparatorInt64(2, 1), 1)
	})
}

func Test_ComparatorUint(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorUint(1, 1), 0)
		t.Assert(qn_util.ComparatorUint(1, 2), -1)
		t.Assert(qn_util.ComparatorUint(2, 1), 1)
	})
}

func Test_ComparatorUint8(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorUint8(1, 1), 0)
		t.Assert(qn_util.ComparatorUint8(2, 6), 252)
		t.Assert(qn_util.ComparatorUint8(2, 1), 1)
	})
}

func Test_ComparatorUint16(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorUint16(1, 1), 0)
		t.Assert(qn_util.ComparatorUint16(1, 2), 65535)
		t.Assert(qn_util.ComparatorUint16(2, 1), 1)
	})
}

func Test_ComparatorUint32(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorUint32(1, 1), 0)
		t.Assert(qn_util.ComparatorUint32(-1000, 2147483640), 2147482656)
		t.Assert(qn_util.ComparatorUint32(2, 1), 1)
	})
}

func Test_ComparatorUint64(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorUint64(1, 1), 0)
		t.Assert(qn_util.ComparatorUint64(1, 2), -1)
		t.Assert(qn_util.ComparatorUint64(2, 1), 1)
	})
}

func Test_ComparatorFloat32(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorFloat32(1, 1), 0)
		t.Assert(qn_util.ComparatorFloat32(1, 2), -1)
		t.Assert(qn_util.ComparatorFloat32(2, 1), 1)
	})
}

func Test_ComparatorFloat64(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorFloat64(1, 1), 0)
		t.Assert(qn_util.ComparatorFloat64(1, 2), -1)
		t.Assert(qn_util.ComparatorFloat64(2, 1), 1)
	})
}

func Test_ComparatorByte(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorByte(1, 1), 0)
		t.Assert(qn_util.ComparatorByte(1, 2), 255)
		t.Assert(qn_util.ComparatorByte(2, 1), 1)
	})
}

func Test_ComparatorRune(t *testing.T) {

	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_util.ComparatorRune(1, 1), 0)
		t.Assert(qn_util.ComparatorRune(1, 2), -1)
		t.Assert(qn_util.ComparatorRune(2, 1), 1)
	})
}

func Test_ComparatorTime(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_util.ComparatorTime("2019-06-14", "2019-06-14")
		t.Assert(j, 0)

		k := qn_util.ComparatorTime("2019-06-15", "2019-06-14")
		t.Assert(k, 1)

		l := qn_util.ComparatorTime("2019-06-13", "2019-06-14")
		t.Assert(l, -1)
	})
}
