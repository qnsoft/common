// Copyright 2017-2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_env_test

import (
	"os"
	"testing"

	"github.com/qnsoft/common/os/genv"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

func Test_GEnv_All(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(os.Environ(), genv.All())
	})
}

func Test_GEnv_Map(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := qn_conv.String(qn_time.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.Assert(err, nil)
		t.Assert(genv.Map()[key], "TEST")
	})
}

func Test_GEnv_Get(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := qn_conv.String(qn_time.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.Assert(err, nil)
		t.AssertEQ(genv.Get(key), "TEST")
	})
}

func Test_GEnv_Contains(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := qn_conv.String(qn_time.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.Assert(err, nil)
		t.AssertEQ(genv.Contains(key), true)
		t.AssertEQ(genv.Contains("none"), false)
	})
}

func Test_GEnv_Set(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := qn_conv.String(qn_time.TimestampNano())
		key := "TEST_ENV_" + value
		err := genv.Set(key, "TEST")
		t.Assert(err, nil)
		t.AssertEQ(os.Getenv(key), "TEST")
	})
}

func Test_GEnv_Build(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := genv.Build(map[string]string{
			"k1": "v1",
			"k2": "v2",
		})
		t.AssertIN("k1=v1", s)
		t.AssertIN("k2=v2", s)
	})
}

func Test_GEnv_Remove(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		value := qn_conv.String(qn_time.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.Assert(err, nil)
		err = genv.Remove(key)
		t.Assert(err, nil)
		t.AssertEQ(os.Getenv(key), "")
	})
}
