// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_str_test

import (
	"testing"

	"github.com/gogf/gf/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/qn.str"
)

func Test_Replace(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(qn.str.Replace(s1, "ab", "AB"), "ABcdEFG乱入的中文ABcdefg")
		t.Assert(qn.str.Replace(s1, "EF", "ef"), "abcdefG乱入的中文abcdefg")
		t.Assert(qn.str.Replace(s1, "MN", "mn"), s1)

		t.Assert(qn.str.ReplaceByArray(s1, g.ArrayStr{
			"a", "A",
			"A", "-",
			"a",
		}), "-bcdEFG乱入的中文-bcdefg")

		t.Assert(qn.str.ReplaceByMap(s1, qn.MapStrStr{
			"a": "A",
			"G": "g",
		}), "AbcdEFg乱入的中文Abcdefg")
	})
}

func Test_ReplaceI_1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcd乱入的中文ABCD"
		s2 := "a"
		t.Assert(qn.str.ReplaceI(s1, "ab", "aa"), "aacd乱入的中文aaCD")
		t.Assert(qn.str.ReplaceI(s1, "ab", "aa", 0), "abcd乱入的中文ABCD")
		t.Assert(qn.str.ReplaceI(s1, "ab", "aa", 1), "aacd乱入的中文ABCD")

		t.Assert(qn.str.ReplaceI(s1, "abcd", "-"), "-乱入的中文-")
		t.Assert(qn.str.ReplaceI(s1, "abcd", "-", 1), "-乱入的中文ABCD")

		t.Assert(qn.str.ReplaceI(s1, "abcd乱入的", ""), "中文ABCD")
		t.Assert(qn.str.ReplaceI(s1, "ABCD乱入的", ""), "中文ABCD")

		t.Assert(qn.str.ReplaceI(s2, "A", "-"), "-")
		t.Assert(qn.str.ReplaceI(s2, "a", "-"), "-")

		t.Assert(qn.str.ReplaceIByArray(s1, g.ArrayStr{
			"abcd乱入的", "-",
			"-", "=",
			"a",
		}), "=中文ABCD")

		t.Assert(qn.str.ReplaceIByMap(s1, qn.MapStrStr{
			"ab": "-",
			"CD": "=",
		}), "-=乱入的中文-=")
	})
}

func Test_ToLower(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "abcdefg乱入的中文abcdefg"
		t.Assert(qn.str.ToLower(s1), e1)
	})
}

func Test_ToUpper(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "ABCDEFG乱入的中文ABCDEFG"
		t.Assert(qn.str.ToUpper(s1), e1)
	})
}

func Test_UcFirst(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "abcdEFG乱入的中文abcdefg"
		e1 := "AbcdEFG乱入的中文abcdefg"
		t.Assert(qn.str.UcFirst(""), "")
		t.Assert(qn.str.UcFirst(s1), e1)
		t.Assert(qn.str.UcFirst(e1), e1)
	})
}

func Test_LcFirst(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "AbcdEFG乱入的中文abcdefg"
		e1 := "abcdEFG乱入的中文abcdefg"
		t.Assert(qn.str.LcFirst(""), "")
		t.Assert(qn.str.LcFirst(s1), e1)
		t.Assert(qn.str.LcFirst(e1), e1)
	})
}

func Test_UcWords(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s1 := "我爱GF: i love go frame"
		e1 := "我爱GF: I Love Go Frame"
		t.Assert(qn.str.UcWords(s1), e1)
	})
}

func Test_IsLetterLower(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.IsLetterLower('a'), true)
		t.Assert(qn.str.IsLetterLower('A'), false)
		t.Assert(qn.str.IsLetterLower('1'), false)
	})
}

func Test_IsLetterUpper(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.IsLetterUpper('a'), false)
		t.Assert(qn.str.IsLetterUpper('A'), true)
		t.Assert(qn.str.IsLetterUpper('1'), false)
	})
}

func Test_IsNumeric(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.IsNumeric("1a我"), false)
		t.Assert(qn.str.IsNumeric("0123"), true)
		t.Assert(qn.str.IsNumeric("我是中国人"), false)
	})
}

func Test_SubStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.SubStr("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(qn.str.SubStr("我爱GoFrame", 6), "GoFrame")
		t.Assert(qn.str.SubStr("我爱GoFrame", 6, 2), "Go")
		t.Assert(qn.str.SubStr("我爱GoFrame", -1, 30), "我爱GoFrame")
		t.Assert(qn.str.SubStr("我爱GoFrame", 30, 30), "")
	})
}

func Test_SubStrRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.SubStrRune("我爱GoFrame", 0), "我爱GoFrame")
		t.Assert(qn.str.SubStrRune("我爱GoFrame", 2), "GoFrame")
		t.Assert(qn.str.SubStrRune("我爱GoFrame", 2, 2), "Go")
		t.Assert(qn.str.SubStrRune("我爱GoFrame", -1, 30), "我爱GoFrame")
		t.Assert(qn.str.SubStrRune("我爱GoFrame", 30, 30), "")
	})
}

func Test_StrLimit(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.StrLimit("我爱GoFrame", 6), "我爱...")
		t.Assert(qn.str.StrLimit("我爱GoFrame", 6, ""), "我爱")
		t.Assert(qn.str.StrLimit("我爱GoFrame", 6, "**"), "我爱**")
		t.Assert(qn.str.StrLimit("我爱GoFrame", 8, ""), "我爱Go")
		t.Assert(qn.str.StrLimit("*", 4, ""), "*")
	})
}

func Test_StrLimitRune(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.StrLimitRune("我爱GoFrame", 2), "我爱...")
		t.Assert(qn.str.StrLimitRune("我爱GoFrame", 2, ""), "我爱")
		t.Assert(qn.str.StrLimitRune("我爱GoFrame", 2, "**"), "我爱**")
		t.Assert(qn.str.StrLimitRune("我爱GoFrame", 4, ""), "我爱Go")
		t.Assert(qn.str.StrLimitRune("*", 4, ""), "*")
	})
}

func Test_HasPrefix(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.HasPrefix("我爱GoFrame", "我爱"), true)
		t.Assert(qn.str.HasPrefix("en我爱GoFrame", "我爱"), false)
		t.Assert(qn.str.HasPrefix("en我爱GoFrame", "en"), true)
	})
}

func Test_HasSuffix(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.HasSuffix("我爱GoFrame", "GoFrame"), true)
		t.Assert(qn.str.HasSuffix("en我爱GoFrame", "a"), false)
		t.Assert(qn.str.HasSuffix("GoFrame很棒", "棒"), true)
	})
}

func Test_Reverse(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Reverse("我爱123"), "321爱我")
	})
}

func Test_NumberFormat(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.NumberFormat(1234567.8910, 2, ".", ","), "1,234,567.89")
		t.Assert(qn.str.NumberFormat(1234567.8910, 2, "#", "/"), "1/234/567#89")
		t.Assert(qn.str.NumberFormat(-1234567.8910, 2, "#", "/"), "-1/234/567#89")
	})
}

func Test_ChunkSplit(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.ChunkSplit("1234", 1, "#"), "1#2#3#4#")
		t.Assert(qn.str.ChunkSplit("我爱123", 1, "#"), "我#爱#1#2#3#")
		t.Assert(qn.str.ChunkSplit("1234", 1, ""), "1\r\n2\r\n3\r\n4\r\n")
	})
}

func Test_SplitAndTrim(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := `

010    

020  

`
		a := qn.str.SplitAndTrim(s, "\n", "0")
		t.Assert(len(a), 2)
		t.Assert(a[0], "1")
		t.Assert(a[1], "2")
	})
}

func Test_Fields(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Fields("我爱 Go Frame"), []string{
			"我爱", "Go", "Frame",
		})
	})
}

func Test_CountWords(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.CountWords("我爱 Go Go Go"), map[string]int{
			"Go": 3,
			"我爱": 1,
		})
	})
}

func Test_CountChars(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.CountChars("我爱 Go Go Go"), map[string]int{
			" ": 3,
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
		t.Assert(qn.str.CountChars("我爱 Go Go Go", true), map[string]int{
			"G": 3,
			"o": 3,
			"我": 1,
			"爱": 1,
		})
	})
}

func Test_WordWrap(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.WordWrap("12 34", 2, "<br>"), "12<br>34")
		t.Assert(qn.str.WordWrap("12 34", 2, "\n"), "12\n34")
		t.Assert(qn.str.WordWrap("我爱 GF", 2, "\n"), "我爱\nGF")
		t.Assert(qn.str.WordWrap("A very long woooooooooooooooooord. and something", 7, "<br>"),
			"A very<br>long<br>woooooooooooooooooord.<br>and<br>something")
	})
}

func Test_RuneLen(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.RuneLen("1234"), 4)
		t.Assert(qn.str.RuneLen("我爱GoFrame"), 9)
	})
}

func Test_Repeat(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Repeat("go", 3), "gogogo")
		t.Assert(qn.str.Repeat("好的", 3), "好的好的好的")
	})
}

func Test_Str(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Str("name@example.com", "@"), "@example.com")
		t.Assert(qn.str.Str("name@example.com", ""), "")
		t.Assert(qn.str.Str("name@example.com", "z"), "")
	})
}

func Test_Shuffle(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(len(qn.str.Shuffle("123456")), 6)
	})
}

func Test_Split(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Split("1.2", "."), []string{"1", "2"})
		t.Assert(qn.str.Split("我爱 - GoFrame", " - "), []string{"我爱", "GoFrame"})
	})
}

func Test_Join(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Join([]string{"我爱", "GoFrame"}, " - "), "我爱 - GoFrame")
	})
}

func Test_Explode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Explode(" - ", "我爱 - GoFrame"), []string{"我爱", "GoFrame"})
	})
}

func Test_Implode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Implode(" - ", []string{"我爱", "GoFrame"}), "我爱 - GoFrame")
	})
}

func Test_Chr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Chr(65), "A")
	})
}

func Test_Ord(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Ord("A"), 65)
	})
}

func Test_HideStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.HideStr("15928008611", 40, "*"), "159****8611")
		t.Assert(qn.str.HideStr("john@kohg.cn", 40, "*"), "jo*n@kohg.cn")
		t.Assert(qn.str.HideStr("张三", 50, "*"), "张*")
		t.Assert(qn.str.HideStr("张小三", 50, "*"), "张*三")
		t.Assert(qn.str.HideStr("欧阳小三", 50, "*"), "欧**三")
	})
}

func Test_Nl2Br(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Nl2Br("1\n2"), "1<br>2")
		t.Assert(qn.str.Nl2Br("1\r\n2"), "1<br>2")
		t.Assert(qn.str.Nl2Br("1\r\n2", true), "1<br />2")
	})
}

func Test_AddSlashes(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.AddSlashes(`1'2"3\`), `1\'2\"3\\`)
	})
}

func Test_StripSlashes(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.StripSlashes(`1\'2\"3\\`), `1'2"3\`)
	})
}

func Test_QuoteMeta(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.QuoteMeta(`.\+*?[^]($)`), `\.\\\+\*\?\[\^\]\(\$\)`)
		t.Assert(qn.str.QuoteMeta(`.\+*中国?[^]($)`), `\.\\\+\*中国\?\[\^\]\(\$\)`)
		t.Assert(qn.str.QuoteMeta(`.''`, `'`), `.\'\'`)
		t.Assert(qn.str.QuoteMeta(`中国.''`, `'`), `中国.\'\'`)
	})
}

func Test_Count(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := "abcdaAD"
		t.Assert(qn.str.Count(s, "0"), 0)
		t.Assert(qn.str.Count(s, "a"), 2)
		t.Assert(qn.str.Count(s, "b"), 1)
		t.Assert(qn.str.Count(s, "d"), 1)
	})
}

func Test_CountI(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := "abcdaAD"
		t.Assert(qn.str.CountI(s, "0"), 0)
		t.Assert(qn.str.CountI(s, "a"), 3)
		t.Assert(qn.str.CountI(s, "b"), 1)
		t.Assert(qn.str.CountI(s, "d"), 2)
	})
}

func Test_Compare(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Compare("a", "b"), -1)
		t.Assert(qn.str.Compare("a", "a"), 0)
		t.Assert(qn.str.Compare("b", "a"), 1)
	})
}

func Test_Equal(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Equal("a", "A"), true)
		t.Assert(qn.str.Equal("a", "a"), true)
		t.Assert(qn.str.Equal("b", "a"), false)
	})
}

func Test_Contains(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.Contains("abc", "a"), true)
		t.Assert(qn.str.Contains("abc", "A"), false)
		t.Assert(qn.str.Contains("abc", "ab"), true)
		t.Assert(qn.str.Contains("abc", "abc"), true)
	})
}

func Test_ContainsI(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.ContainsI("abc", "a"), true)
		t.Assert(qn.str.ContainsI("abc", "A"), true)
		t.Assert(qn.str.ContainsI("abc", "Ab"), true)
		t.Assert(qn.str.ContainsI("abc", "ABC"), true)
		t.Assert(qn.str.ContainsI("abc", "ABCD"), false)
		t.Assert(qn.str.ContainsI("abc", "D"), false)
	})
}

func Test_ContainsAny(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn.str.ContainsAny("abc", "a"), true)
		t.Assert(qn.str.ContainsAny("abc", "cd"), true)
		t.Assert(qn.str.ContainsAny("abc", "de"), false)
		t.Assert(qn.str.ContainsAny("abc", "A"), false)
	})
}

func Test_SearchArray(t *testing.T) {
	qn_tesqn.Slt, func(t *qn_test.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(qn.str.SearchArray(a, "a"), 0)
		t.AssertEQ(qn.str.SearchArray(a, "b"), 1)
		t.AssertEQ(qn.str.SearchArray(a, "c"), 2)
		t.AssertEQ(qn.str.SearchArray(a, "d"), -1)
	})
}

func Test_InArray(t *testing.T) {
	qn_tesqn.Slt, func(t *qn_test.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(qn.str.InArray(a, "a"), true)
		t.AssertEQ(qn.str.InArray(a, "b"), true)
		t.AssertEQ(qn.str.InArray(a, "c"), true)
		t.AssertEQ(qn.str.InArray(a, "d"), false)
	})
}
