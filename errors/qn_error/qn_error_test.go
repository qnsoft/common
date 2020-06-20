// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_error_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/qnsoft/common/errors/qn_error"
	"github.com/qnsoft/common/test/qn_test"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_error.New(""), nil)
		t.Assert(qn_error.Wrap(nilError(), "test"), nil)
	})
}

func Test_Wrap(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(err.Error(), "3: 2: 1")
	})
}

func Test_Cause(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		t.Assert(qn_error.Cause(err), err)
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.Assert(qn_error.Cause(err), "1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		t.Assert(qn_error.Cause(err), "1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.Assert(qn_error.Cause(err), "1")
	})
}

func Test_Format(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%s", err), "3: 2: 1")
		t.Assert(fmt.Sprintf("%v", err), "3: 2: 1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		t.Assert(fmt.Sprintf("%-s", err), "3")
		t.Assert(fmt.Sprintf("%-v", err), "3")
	})
}

func Test_Stack(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		t.Assert(fmt.Sprintf("%+v", err), "1")
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := errors.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		//fmt.Printf("%+v", err)
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		t.AssertNE(fmt.Sprintf("%+v", err), "1")
		//fmt.Printf("%+v", err)
	})

	qn_test.C(t, func(t *qn_test.T) {
		err := qn_error.New("1")
		err = qn_error.Wrap(err, "2")
		err = qn_error.Wrap(err, "3")
		t.AssertNE(err, nil)
		//fmt.Printf("%+v", err)
	})
}
