// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv

import (
	"time"

	"github.com/qnsoft/common/internal/utils"
	"github.com/qnsoft/common/os/qn_time"
)

// Time converts <i> to time.Time.
func Time(i interface{}, format ...string) time.Time {
	// It's already this type.
	if len(format) == 0 {
		if v, ok := i.(time.Time); ok {
			return v
		}
	}
	if t := qn_time(i, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration converts <i> to time.Duration.
// If <i> is string, then it uses time.ParseDuration to convert it.
// If <i> is numeric, then it converts <i> as nanoseconds.
func Duration(i interface{}) time.Duration {
	// It's already this type.
	if v, ok := i.(time.Duration); ok {
		return v
	}
	s := String(i)
	if !utils.IsNumeric(s) {
		d, _ := time.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(i))
}

// qn_time converts <i> to *qn_time.Time.
// The parameter <format> can be used to specify the format of <i>.
// If no <format> given, it converts <i> using qn_time.NewFromTimeStamp if <i> is numeric,
// or using qn_time.StrToTime if <i> is string.
func qn_time(i interface{}, format ...string) *qn_time.Time {
	if i == nil {
		return nil
	}
	// It's already this type.
	if len(format) == 0 {
		if v, ok := i.(*qn_time.Time); ok {
			return v
		}
	}
	s := String(i)
	if len(s) == 0 {
		return qn_time.New()
	}
	// Priority conversion using given format.
	if len(format) > 0 {
		t, _ := qn_time.StrToTimeFormat(s, format[0])
		return t
	}
	if utils.IsNumeric(s) {
		return qn_time.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := qn_time.StrToTime(s)
		return t
	}
}
