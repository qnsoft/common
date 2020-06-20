// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_regex_test

import (
	"strings"
	"testing"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/qn_regex"
)

var (
	PatternErr = `([\d+`
)

func Test_Quote(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := `[foo]` //`\[foo\]`
		t.Assert(qn_regex.Quote(s1), `\[foo\]`)
	})
}

func Test_Validate(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var s1 = `(.+):(\d+)`
		t.Assert(qn_regex.Validate(s1), nil)
		s1 = `((.+):(\d+)`
		t.Assert(qn_regex.Validate(s1) == nil, false)
	})
}

func Test_IsMatch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var pattern = `(.+):(\d+)`
		s1 := []byte(`sfs:2323`)
		t.Assert(qn_regex.IsMatch(pattern, s1), true)
		s1 = []byte(`sfs2323`)
		t.Assert(qn_regex.IsMatch(pattern, s1), false)
		s1 = []byte(`sfs:`)
		t.Assert(qn_regex.IsMatch(pattern, s1), false)
		// error pattern
		t.Assert(qn_regex.IsMatch(PatternErr, s1), false)
	})
}

func Test_IsMatchString(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var pattern = `(.+):(\d+)`
		s1 := `sfs:2323`
		t.Assert(qn_regex.IsMatchString(pattern, s1), true)
		s1 = `sfs2323`
		t.Assert(qn_regex.IsMatchString(pattern, s1), false)
		s1 = `sfs:`
		t.Assert(qn_regex.IsMatchString(pattern, s1), false)
		// error pattern
		t.Assert(qn_regex.IsMatchString(PatternErr, s1), false)
	})
}

func Test_Match(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := qn_regex.Match(re, []byte(s))
		t.Assert(err, nil)
		if string(subs[0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
		}
		if string(subs[1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
		}
		// error pattern
		_, err = qn_regex.Match(PatternErr, []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_MatchString(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := qn_regex.MatchString(re, s)
		t.Assert(err, nil)
		if string(subs[0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0], wantSubs)
		}
		if string(subs[1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1], "aab")
		}
		// error pattern
		_, err = qn_regex.MatchString(PatternErr, s)
		t.AssertNE(err, nil)
	})
}

func Test_MatchAll(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		s = s + `其他的` + s
		subs, err := qn_regex.MatchAll(re, []byte(s))
		t.Assert(err, nil)
		if string(subs[0][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
		}
		if string(subs[0][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
		}

		if string(subs[1][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
		}
		if string(subs[1][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
		}
		// error pattern
		_, err = qn_regex.MatchAll(PatternErr, []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_MatchAllString(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		s := "acbb" + wantSubs + "dd"
		subs, err := qn_regex.MatchAllString(re, s+`其他的`+s)
		t.Assert(err, nil)
		if string(subs[0][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[0][0], wantSubs)
		}
		if string(subs[0][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[0][1], "aab")
		}

		if string(subs[1][0]) != wantSubs {
			t.Fatalf("regex:%s,Match(%q)[0] = %q; want %q", re, s, subs[1][0], wantSubs)
		}
		if string(subs[1][1]) != "aab" {
			t.Fatalf("Match(%q)[1] = %q; want %q", s, subs[1][1], "aab")
		}
		// error pattern
		_, err = qn_regex.MatchAllString(PatternErr, s)
		t.AssertNE(err, nil)
	})
}

func Test_Replace(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		replace := "12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb" + replace + "dd"
		replacedStr, err := qn_regex.Replace(re, []byte(replace), []byte(s))
		t.Assert(err, nil)
		if string(replacedStr) != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = qn_regex.Replace(PatternErr, []byte(replace), []byte(s))
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceString(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		replace := "12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb" + replace + "dd"
		replacedStr, err := qn_regex.ReplaceString(re, replace, s)
		t.Assert(err, nil)
		if replacedStr != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = qn_regex.ReplaceString(PatternErr, replace, s)
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceFun(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		//replace :="12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb[x" + wantSubs + "y]dd"
		wanted = "acbb" + "3个a" + "dd"
		replacedStr, err := qn_regex.ReplaceFunc(re, []byte(s), func(s []byte) []byte {
			if strings.Index(string(s), "aaa") >= 0 {
				return []byte("3个a")
			}
			return []byte("[x" + string(s) + "y]")
		})
		t.Assert(err, nil)
		if string(replacedStr) != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = qn_regex.ReplaceFunc(PatternErr, []byte(s), func(s []byte) []byte {
			return []byte("")
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceFuncMatch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := []byte("1234567890")
		p := `(\d{3})(\d{3})(.+)`
		s0, e0 := qn_regex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[0]
		})
		t.Assert(e0, nil)
		t.Assert(s0, s)
		s1, e1 := qn_regex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[1]
		})
		t.Assert(e1, nil)
		t.Assert(s1, []byte("123"))
		s2, e2 := qn_regex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[2]
		})
		t.Assert(e2, nil)
		t.Assert(s2, []byte("456"))
		s3, e3 := qn_regex.ReplaceFuncMatch(p, s, func(match [][]byte) []byte {
			return match[3]
		})
		t.Assert(e3, nil)
		t.Assert(s3, []byte("7890"))
		// error pattern
		_, err := qn_regex.ReplaceFuncMatch(PatternErr, s, func(match [][]byte) []byte {
			return match[3]
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceStringFunc(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		wantSubs := "aaabb"
		//replace :="12345"
		s := "acbb" + wantSubs + "dd"
		wanted := "acbb[x" + wantSubs + "y]dd"
		wanted = "acbb" + "3个a" + "dd"
		replacedStr, err := qn_regex.ReplaceStringFunc(re, s, func(s string) string {
			if strings.Index(s, "aaa") >= 0 {
				return "3个a"
			}
			return "[x" + s + "y]"
		})
		t.Assert(err, nil)
		if replacedStr != wanted {
			t.Fatalf("regex:%s,old:%s; want %q", re, s, wanted)
		}
		// error pattern
		_, err = qn_regex.ReplaceStringFunc(PatternErr, s, func(s string) string {
			return ""
		})
		t.AssertNE(err, nil)
	})
}

func Test_ReplaceStringFuncMatch(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := "1234567890"
		p := `(\d{3})(\d{3})(.+)`
		s0, e0 := qn_regex.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[0]
		})
		t.Assert(e0, nil)
		t.Assert(s0, s)
		s1, e1 := qn_regex.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[1]
		})
		t.Assert(e1, nil)
		t.Assert(s1, "123")
		s2, e2 := qn_regex.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[2]
		})
		t.Assert(e2, nil)
		t.Assert(s2, "456")
		s3, e3 := qn_regex.ReplaceStringFuncMatch(p, s, func(match []string) string {
			return match[3]
		})
		t.Assert(e3, nil)
		t.Assert(s3, "7890")
		// error pattern
		_, err := qn_regex.ReplaceStringFuncMatch(PatternErr, s, func(match []string) string {
			return ""
		})
		t.AssertNE(err, nil)
	})
}

func Test_Split(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		matched := "aaabb"
		item0 := "acbb"
		item1 := "dd"
		s := item0 + matched + item1
		t.Assert(qn_regex.IsMatchString(re, matched), true)
		items := qn_regex.Split(re, s) //split string with matched
		if items[0] != item0 {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
		if items[1] != item1 {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
	})

	qn_test.C(t, func(t *qn_test.T) {
		re := "a(a+b+)b"
		notmatched := "aaxbb"
		item0 := "acbb"
		item1 := "dd"
		s := item0 + notmatched + item1
		t.Assert(qn_regex.IsMatchString(re, notmatched), false)
		items := qn_regex.Split(re, s) //split string with notmatched then nosplitting
		if items[0] != s {
			t.Fatalf("regex:%s,Split(%q) want %q", re, s, item0)
		}
		// error pattern
		items = qn_regex.Split(PatternErr, s)
		t.AssertEQ(items, nil)

	})
}
