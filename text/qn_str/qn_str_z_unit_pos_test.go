// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_str_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Pos(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(qn.str.Pos(s1, "ab"), 0)
		t.Assert(qn.str.Pos(s1, "ab", 2), 7)
		t.Assert(qn.str.Pos(s1, "abd", 0), -1)
		t.Assert(qn.str.Pos(s1, "e", -4), 11)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.Pos(s1, "爱"), 3)
		t.Assert(qn.str.Pos(s1, "C"), 6)
		t.Assert(qn.str.Pos(s1, "China"), 6)
	})
}

func Test_PosRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(qn.str.PosRune(s1, "ab"), 0)
		t.Assert(qn.str.PosRune(s1, "ab", 2), 7)
		t.Assert(qn.str.PosRune(s1, "abd", 0), -1)
		t.Assert(qn.str.PosRune(s1, "e", -4), 11)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosRune(s1, "爱"), 1)
		t.Assert(qn.str.PosRune(s1, "C"), 2)
		t.Assert(qn.str.PosRune(s1, "China"), 2)
	})
}

func Test_PosI(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(qn.str.PosI(s1, "zz"), -1)
		t.Assert(qn.str.PosI(s1, "ab"), 0)
		t.Assert(qn.str.PosI(s1, "ef", 2), 4)
		t.Assert(qn.str.PosI(s1, "abd", 0), -1)
		t.Assert(qn.str.PosI(s1, "E", -4), 11)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosI(s1, "爱"), 3)
		t.Assert(qn.str.PosI(s1, "c"), 6)
		t.Assert(qn.str.PosI(s1, "china"), 6)
	})
}

func Test_PosIRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		t.Assert(qn.str.PosIRune(s1, "zz"), -1)
		t.Assert(qn.str.PosIRune(s1, "ab"), 0)
		t.Assert(qn.str.PosIRune(s1, "ef", 2), 4)
		t.Assert(qn.str.PosIRune(s1, "abd", 0), -1)
		t.Assert(qn.str.PosIRune(s1, "E", -4), 11)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosIRune(s1, "爱"), 1)
		t.Assert(qn.str.PosIRune(s1, "c"), 2)
		t.Assert(qn.str.PosIRune(s1, "china"), 2)
	})
}

func Test_PosR(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(qn.str.PosR(s1, "zz"), -1)
		t.Assert(qn.str.PosR(s1, "ab"), 7)
		t.Assert(qn.str.PosR(s2, "ab", -2), 0)
		t.Assert(qn.str.PosR(s1, "ef"), 11)
		t.Assert(qn.str.PosR(s1, "abd", 0), -1)
		t.Assert(qn.str.PosR(s1, "e", -4), -1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosR(s1, "爱"), 3)
		t.Assert(qn.str.PosR(s1, "C"), 6)
		t.Assert(qn.str.PosR(s1, "China"), 6)
	})
}

func Test_PosRRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(qn.str.PosRRune(s1, "zz"), -1)
		t.Assert(qn.str.PosRRune(s1, "ab"), 7)
		t.Assert(qn.str.PosRRune(s2, "ab", -2), 0)
		t.Assert(qn.str.PosRRune(s1, "ef"), 11)
		t.Assert(qn.str.PosRRune(s1, "abd", 0), -1)
		t.Assert(qn.str.PosRRune(s1, "e", -4), -1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosRRune(s1, "爱"), 1)
		t.Assert(qn.str.PosRRune(s1, "C"), 2)
		t.Assert(qn.str.PosRRune(s1, "China"), 2)
	})
}

func Test_PosRI(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(qn.str.PosRI(s1, "zz"), -1)
		t.Assert(qn.str.PosRI(s1, "AB"), 7)
		t.Assert(qn.str.PosRI(s2, "AB", -2), 0)
		t.Assert(qn.str.PosRI(s1, "EF"), 11)
		t.Assert(qn.str.PosRI(s1, "abd", 0), -1)
		t.Assert(qn.str.PosRI(s1, "e", -5), 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosRI(s1, "爱"), 3)
		t.Assert(qn.str.PosRI(s1, "C"), 19)
		t.Assert(qn.str.PosRI(s1, "China"), 6)
	})
}

func Test_PosRIRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFGabcdefg"
		s2 := "abcdEFGz1cdeab"
		t.Assert(qn.str.PosRIRune(s1, "zz"), -1)
		t.Assert(qn.str.PosRIRune(s1, "AB"), 7)
		t.Assert(qn.str.PosRIRune(s2, "AB", -2), 0)
		t.Assert(qn.str.PosRIRune(s1, "EF"), 11)
		t.Assert(qn.str.PosRIRune(s1, "abd", 0), -1)
		t.Assert(qn.str.PosRIRune(s1, "e", -5), 4)
	})
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱China very much"
		t.Assert(qn.str.PosRIRune(s1, "爱"), 1)
		t.Assert(qn.str.PosRIRune(s1, "C"), 15)
		t.Assert(qn.str.PosRIRune(s1, "China"), 2)
	})
}
