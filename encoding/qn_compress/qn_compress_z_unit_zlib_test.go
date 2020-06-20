// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_compress_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_compress"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Zlib_UnZlib(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := qn_compress.Zlib([]byte(src))
		t.Assert(data, dst)

		data, _ = qn_compress.UnZlib(dst)
		t.Assert(data, []byte(src))

		data, _ = qn_compress.Zlib(nil)
		t.Assert(data, nil)
		data, _ = qn_compress.UnZlib(nil)
		t.Assert(data, nil)

		data, _ = qn_compress.UnZlib(dst[1:])
		t.Assert(data, nil)
	})
}
