// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_binary_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_binary"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_BeEncodeAndBeDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for k, v := range testData {
			ve := qn_binary.BeEncode(v)
			ve1 := qn_binary.BeEncodeByLength(len(ve), v)

			//t.Logf("%s:%v, encoded:%v\n", k, v, ve)
			switch v.(type) {
			case int:
				t.Assert(qn_binary.BeDecodeToInt(ve), v)
				t.Assert(qn_binary.BeDecodeToInt(ve1), v)
			case int8:
				t.Assert(qn_binary.BeDecodeToInt8(ve), v)
				t.Assert(qn_binary.BeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(qn_binary.BeDecodeToInt16(ve), v)
				t.Assert(qn_binary.BeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(qn_binary.BeDecodeToInt32(ve), v)
				t.Assert(qn_binary.BeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(qn_binary.BeDecodeToInt64(ve), v)
				t.Assert(qn_binary.BeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(qn_binary.BeDecodeToUint(ve), v)
				t.Assert(qn_binary.BeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(qn_binary.BeDecodeToUint8(ve), v)
				t.Assert(qn_binary.BeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(qn_binary.BeDecodeToUint16(ve1), v)
				t.Assert(qn_binary.BeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(qn_binary.BeDecodeToUint32(ve1), v)
				t.Assert(qn_binary.BeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(qn_binary.BeDecodeToUint64(ve), v)
				t.Assert(qn_binary.BeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(qn_binary.BeDecodeToBool(ve), v)
				t.Assert(qn_binary.BeDecodeToBool(ve1), v)
			case string:
				t.Assert(qn_binary.BeDecodeToString(ve), v)
				t.Assert(qn_binary.BeDecodeToString(ve1), v)
			case float32:
				t.Assert(qn_binary.BeDecodeToFloat32(ve), v)
				t.Assert(qn_binary.BeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(qn_binary.BeDecodeToFloat64(ve), v)
				t.Assert(qn_binary.BeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := qn_binary.BeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_BeEncodeStruct(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := qn_binary.BeEncode(user)
		s := qn_binary.BeDecodeToString(ve)
		t.Assert(string(s), s)
	})
}
