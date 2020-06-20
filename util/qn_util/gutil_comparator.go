// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_util

import (
	"strings"

	"github.com/qnsoft/common/util/qn_conv"
)

// Comparator is a function that compare a and b, and returns the result as int.
//
// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator func(a, b interface{}) int

// ComparatorString provides a fast comparison on strings.
func ComparatorString(a, b interface{}) int {
	return strings.Compare(qn_conv.String(a), qn_conv.String(b))
}

// ComparatorInt provides a basic comparison on int.
func ComparatorInt(a, b interface{}) int {
	return qn_conv.Int(a) - qn_conv.Int(b)
}

// ComparatorInt8 provides a basic comparison on int8.
func ComparatorInt8(a, b interface{}) int {
	return int(qn_conv.Int8(a) - qn_conv.Int8(b))
}

// ComparatorInt16 provides a basic comparison on int16.
func ComparatorInt16(a, b interface{}) int {
	return int(qn_conv.Int16(a) - qn_conv.Int16(b))
}

// ComparatorInt32 provides a basic comparison on int32.
func ComparatorInt32(a, b interface{}) int {
	return int(qn_conv.Int32(a) - qn_conv.Int32(b))
}

// ComparatorInt64 provides a basic comparison on int64.
func ComparatorInt64(a, b interface{}) int {
	return int(qn_conv.Int64(a) - qn_conv.Int64(b))
}

// ComparatorUint provides a basic comparison on uint.
func ComparatorUint(a, b interface{}) int {
	return int(qn_conv.Uint(a) - qn_conv.Uint(b))
}

// ComparatorUint8 provides a basic comparison on uint8.
func ComparatorUint8(a, b interface{}) int {
	return int(qn_conv.Uint8(a) - qn_conv.Uint8(b))
}

// ComparatorUint16 provides a basic comparison on uint16.
func ComparatorUint16(a, b interface{}) int {
	return int(qn_conv.Uint16(a) - qn_conv.Uint16(b))
}

// ComparatorUint32 provides a basic comparison on uint32.
func ComparatorUint32(a, b interface{}) int {
	return int(qn_conv.Uint32(a) - qn_conv.Uint32(b))
}

// ComparatorUint64 provides a basic comparison on uint64.
func ComparatorUint64(a, b interface{}) int {
	return int(qn_conv.Uint64(a) - qn_conv.Uint64(b))
}

// ComparatorFloat32 provides a basic comparison on float32.
func ComparatorFloat32(a, b interface{}) int {
	return int(qn_conv.Float32(a) - qn_conv.Float32(b))
}

// ComparatorFloat64 provides a basic comparison on float64.
func ComparatorFloat64(a, b interface{}) int {
	return int(qn_conv.Float64(a) - qn_conv.Float64(b))
}

// ComparatorByte provides a basic comparison on byte.
func ComparatorByte(a, b interface{}) int {
	return int(qn_conv.Byte(a) - qn_conv.Byte(b))
}

// ComparatorRune provides a basic comparison on rune.
func ComparatorRune(a, b interface{}) int {
	return int(qn_conv.Rune(a) - qn_conv.Rune(b))
}

// ComparatorTime provides a basic comparison on time.Time.
func ComparatorTime(a, b interface{}) int {
	aTime := qn_conv.Time(a)
	bTime := qn_conv.Time(b)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}
