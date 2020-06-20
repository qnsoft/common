// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package utils_test

import (
	"io/ioutil"
	"testing"

	"github.com/qnsoft/common/internal/utils"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_ReadCloser(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var (
			n    int
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var (
			r    []byte
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var (
			n    int
			r    []byte
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, true)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = ioutil.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
	})
}
