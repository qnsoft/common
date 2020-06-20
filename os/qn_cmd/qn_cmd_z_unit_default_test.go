// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*" -benchmem

package qn_cmd_test

import (
	"os"
	"testing"

	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/os/gcmd"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Default(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		os.Args = []string{"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"}
		t.Assert(len(gcmd.GetArgAll()), 4)
		t.Assert(gcmd.GetArg(1), "remove")
		t.Assert(gcmd.GetArg(100, "test"), "test")
		t.Assert(gcmd.GetOpt("n"), "")
		t.Assert(gcmd.ContainsOpt("p"), true)
		t.Assert(gcmd.ContainsOpt("n"), true)
		t.Assert(gcmd.ContainsOpt("none"), false)
		t.Assert(gcmd.GetOpt("none", "value"), "value")

	})
}

func Test_BuildOptions(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := gcmd.BuildOptions(qn.MapStrStr{
			"n": "john",
		})
		t.Assert(s, "-n=john")
	})

	qn_test.C(t, func(t *qn_test.T) {
		s := gcmd.BuildOptions(qn.MapStrStr{
			"n": "john",
		}, "-test")
		t.Assert(s, "-testn=john")
	})
}
