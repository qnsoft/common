// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_time_test

import (
	"testing"

	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Json_Pointer(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time *qn_time.Time
		}
		t1 := new(T)
		s := "2006-01-02 15:04:05"
		t1.Time = qn_time.NewFromStr(s)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{"Time":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time *qn_time.Time
		}
		t1 := new(T)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{"Time":null}`)
	})
	// Marshal nil omitempty
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time *qn_time.Time `json:"time,omitempty"`
		}
		t1 := new(T)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{}`)
	})
	// Unmarshal
	qn_test.C(t, func(t *qn_test.T) {
		var t1 qn_time.Time
		s := []byte(`"2006-01-02 15:04:05"`)
		err := json.Unmarshal(s, &t1)
		t.Assert(err, nil)
		t.Assert(t1.String(), "2006-01-02 15:04:05")
	})
}

func Test_Json_Struct(t *testing.T) {
	// Marshal
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time qn_time.Time
		}
		t1 := new(T)
		s := "2006-01-02 15:04:05"
		t1.Time = *qn_time.NewFromStr(s)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{"Time":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time qn_time.Time
		}
		t1 := new(T)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{"Time":""}`)
	})
	// Marshal nil omitempty
	qn_test.C(t, func(t *qn_test.T) {
		type T struct {
			Time qn_time.Time `json:"time,omitempty"`
		}
		t1 := new(T)
		j, err := json.Marshal(t1)
		t.Assert(err, nil)
		t.Assert(j, `{"time":""}`)
	})

}
