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

func Test_LeEncodeAndLeDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for k, v := range testData {
			ve := qn_binary.LeEncode(v)
			ve1 := qn_binary.LeEncodeByLength(len(ve), v)

			//t.Logf("%s:%v, encoded:%v\n", k, v, ve)
			switch v.(type) {
			case int:
				t.Assert(qn_binary.LeDecodeToInt(ve), v)
				t.Assert(qn_binary.LeDecodeToInt(ve1), v)
			case int8:
				t.Assert(qn_binary.LeDecodeToInt8(ve), v)
				t.Assert(qn_binary.LeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(qn_binary.LeDecodeToInt16(ve), v)
				t.Assert(qn_binary.LeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(qn_binary.LeDecodeToInt32(ve), v)
				t.Assert(qn_binary.LeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(qn_binary.LeDecodeToInt64(ve), v)
				t.Assert(qn_binary.LeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(qn_binary.LeDecodeToUint(ve), v)
				t.Assert(qn_binary.LeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(qn_binary.LeDecodeToUint8(ve), v)
				t.Assert(qn_binary.LeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(qn_binary.LeDecodeToUint16(ve1), v)
				t.Assert(qn_binary.LeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(qn_binary.LeDecodeToUint32(ve1), v)
				t.Assert(qn_binary.LeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(qn_binary.LeDecodeToUint64(ve), v)
				t.Assert(qn_binary.LeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(qn_binary.LeDecodeToBool(ve), v)
				t.Assert(qn_binary.LeDecodeToBool(ve1), v)
			case string:
				t.Assert(qn_binary.LeDecodeToString(ve), v)
				t.Assert(qn_binary.LeDecodeToString(ve1), v)
			case float32:
				t.Assert(qn_binary.LeDecodeToFloat32(ve), v)
				t.Assert(qn_binary.LeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(qn_binary.LeDecodeToFloat64(ve), v)
				t.Assert(qn_binary.LeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := qn_binary.LeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_LeEncodeStruct(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := qn_binary.LeEncode(user)
		s := qn_binary.LeDecodeToString(ve)
		t.Assert(s, s)
	})
}
