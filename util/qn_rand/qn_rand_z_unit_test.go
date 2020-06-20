// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_rand_test

import (
	"testing"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_rand"
)

func Test_Intn(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 1000000; i++ {
			n := qn_rand.Intn(100)
			t.AssertLT(n, 100)
			t.AssertGE(n, 0)
		}
		for i := 0; i < 1000000; i++ {
			n := qn_rand.Intn(-100)
			t.AssertLE(n, 0)
			t.Assert(n, -100)
		}
	})
}

func Test_Meet(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.Meet(100, 100), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.Meet(0, 100), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.Meet(50, 100), []bool{true, false})
		}
	})
}

func Test_MeetProb(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.MeetProb(1), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.MeetProb(0), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.MeetProb(0.5), []bool{true, false})
		}
	})
}

func Test_N(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.N(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.N(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.N(1, 2), []int{1, 2})
		}
	})
}

func Test_Rand(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.N(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(qn_rand.N(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.N(1, 2), []int{1, 2})
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.N(-1, 2), []int{-1, 0, 1, 2})
		}
	})
}

func Test_S(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.S(5)), 5)
		}
	})
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.S(5, true)), 5)
		}
	})
}

func Test_B(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			b := qn_rand.B(5)
			t.Assert(len(b), 5)
			t.AssertNE(b, make([]byte, 5))
		}
	})
}

func Test_Str(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.S(5)), 5)
		}
	})
}

func Test_RandStr(t *testing.T) {
	str := "我爱GoFrame"
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 10; i++ {
			s := qn_rand.Str(str, 100000)
			t.Assert(qn.str.Contains(s, "我"), true)
			t.Assert(qn.str.Contains(s, "爱"), true)
			t.Assert(qn.str.Contains(s, "G"), true)
			t.Assert(qn.str.Contains(s, "o"), true)
			t.Assert(qn.str.Contains(s, "F"), true)
			t.Assert(qn.str.Contains(s, "r"), true)
			t.Assert(qn.str.Contains(s, "a"), true)
			t.Assert(qn.str.Contains(s, "m"), true)
			t.Assert(qn.str.Contains(s, "e"), true)
			t.Assert(qn.str.Contains(s, "w"), false)
		}
	})
}

func Test_Digits(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.Digits(5)), 5)
		}
	})
}

func Test_RandDigits(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.Digits(5)), 5)
		}
	})
}

func Test_Letters(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.Letters(5)), 5)
		}
	})
}

func Test_RandLetters(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(qn_rand.Letters(5)), 5)
		}
	})
}

func Test_Perm(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			t.AssertIN(qn_rand.Perm(5), []int{0, 1, 2, 3, 4})
		}
	})
}
