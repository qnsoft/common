// Copyright 2018-2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_var

import (
	"github.com/qnsoft/common/internal/json"
	"time"

	"github.com/qnsoft/common/internal/empty"

	"github.com/qnsoft/common/container/gtype"
	"github.com/qnsoft/common/os/gtime"
	"github.com/qnsoft/common/util/gconv"
)

// VarImp is an universal variable type implementer.
type VarImp struct {
	value interface{} // Underlying value.
	safe  bool        // Concurrent safe or not.
}

// New creates and returns a new Var with given <value>.
// The optional parameter <safe> specifies whether Var is used in concurrent-safety,
// which is false in default.
func New(value interface{}, safe ...bool) Var {
	v := VarImp{}
	if len(safe) > 0 && !safe[0] {
		v.safe = true
		v.value = gtype.NewInterface(value)
	} else {
		v.value = value
	}
	return &v
}

// Create creates and returns a new VarImp with given <value>.
// The optional parameter <safe> specifies whether Var is used in concurrent-safety,
// which is false in default.
// Deprecated.
func Create(value interface{}, safe ...bool) VarImp {
	v := VarImp{}
	if len(safe) > 0 && !safe[0] {
		v.safe = true
		v.value = gtype.NewInterface(value)
	} else {
		v.value = value
	}
	return v
}

// Clone does a shallow copy of current Var and returns a pointer to this Var.
func (v *VarImp) Clone() Var {
	return New(v.Val(), v.safe)
}

// Set sets <value> to <v>, and returns the old value.
func (v *VarImp) Set(value interface{}) (old interface{}) {
	if v.safe {
		if t, ok := v.value.(*gtype.Interface); ok {
			old = t.Set(value)
			return
		}
	}
	old = v.value
	v.value = value
	return
}

// Val returns the current value of <v>.
func (v *VarImp) Val() interface{} {
	if v == nil {
		return nil
	}
	if v.safe {
		if t, ok := v.value.(*gtype.Interface); ok {
			return t.Val()
		}
	}
	return v.value
}

// Interface is alias of Val.
func (v *VarImp) Interface() interface{} {
	return v.Val()
}

// IsNil checks whether <v> is nil.
func (v *VarImp) IsNil() bool {
	return v.Val() == nil
}

// IsEmpty checks whether <v> is empty.
func (v *VarImp) IsEmpty() bool {
	return empty.IsEmpty(v.Val())
}

// Bytes converts and returns <v> as []byte.
func (v *VarImp) Bytes() []byte {
	return gconv.Bytes(v.Val())
}

// String converts and returns <v> as string.
func (v *VarImp) String() string {
	return gconv.String(v.Val())
}

// Bool converts and returns <v> as bool.
func (v *VarImp) Bool() bool {
	return gconv.Bool(v.Val())
}

// Int converts and returns <v> as int.
func (v *VarImp) Int() int {
	return gconv.Int(v.Val())
}

// Ints converts and returns <v> as []int.
func (v *VarImp) Ints() []int {
	return gconv.Ints(v.Val())
}

// Int8 converts and returns <v> as int8.
func (v *VarImp) Int8() int8 {
	return gconv.Int8(v.Val())
}

// Int16 converts and returns <v> as int16.
func (v *VarImp) Int16() int16 {
	return gconv.Int16(v.Val())
}

// Int32 converts and returns <v> as int32.
func (v *VarImp) Int32() int32 {
	return gconv.Int32(v.Val())
}

// Int64 converts and returns <v> as int64.
func (v *VarImp) Int64() int64 {
	return gconv.Int64(v.Val())
}

// Uint converts and returns <v> as uint.
func (v *VarImp) Uint() uint {
	return gconv.Uint(v.Val())
}

// Uints converts and returns <v> as []uint.
func (v *VarImp) Uints() []uint {
	return gconv.Uints(v.Val())
}

// Uint8 converts and returns <v> as uint8.
func (v *VarImp) Uint8() uint8 {
	return gconv.Uint8(v.Val())
}

// Uint16 converts and returns <v> as uint16.
func (v *VarImp) Uint16() uint16 {
	return gconv.Uint16(v.Val())
}

// Uint32 converts and returns <v> as uint32.
func (v *VarImp) Uint32() uint32 {
	return gconv.Uint32(v.Val())
}

// Uint64 converts and returns <v> as uint64.
func (v *VarImp) Uint64() uint64 {
	return gconv.Uint64(v.Val())
}

// Float32 converts and returns <v> as float32.
func (v *VarImp) Float32() float32 {
	return gconv.Float32(v.Val())
}

// Float64 converts and returns <v> as float64.
func (v *VarImp) Float64() float64 {
	return gconv.Float64(v.Val())
}

// Floats converts and returns <v> as []float64.
func (v *VarImp) Floats() []float64 {
	return gconv.Floats(v.Val())
}

// Strings converts and returns <v> as []string.
func (v *VarImp) Strings() []string {
	return gconv.Strings(v.Val())
}

// Interfaces converts and returns <v> as []interfaces{}.
func (v *VarImp) Interfaces() []interface{} {
	return gconv.Interfaces(v.Val())
}

// Slice is alias of Interfaces.
func (v *VarImp) Slice() []interface{} {
	return v.Interfaces()
}

// Array is alias of Interfaces.
func (v *VarImp) Array() []interface{} {
	return v.Interfaces()
}

// Vars converts and returns <v> as []Var.
func (v *VarImp) Vars() []Var {
	array := gconv.Interfaces(v.Val())
	if len(array) == 0 {
		return nil
	}
	vars := make([]Var, len(array))
	for k, v := range array {
		vars[k] = New(v)
	}
	return vars
}

// Time converts and returns <v> as time.Time.
// The parameter <format> specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
func (v *VarImp) Time(format ...string) time.Time {
	return gconv.Time(v.Val(), format...)
}

// Duration converts and returns <v> as time.Duration.
// If value of <v> is string, then it uses time.ParseDuration for conversion.
func (v *VarImp) Duration() time.Duration {
	return gconv.Duration(v.Val())
}

// GTime converts and returns <v> as *gtime.Time.
// The parameter <format> specifies the format of the time string using gtime,
// eg: Y-m-d H:i:s.
func (v *VarImp) GTime(format ...string) *gtime.Time {
	return gconv.GTime(v.Val(), format...)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (v *VarImp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (v *VarImp) UnmarshalJSON(b []byte) error {
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	v.Set(i)
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for Var.
func (v *VarImp) UnmarshalValue(value interface{}) error {
	v.Set(value)
	return nil
}