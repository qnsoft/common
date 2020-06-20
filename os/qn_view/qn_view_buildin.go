// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_view

import (
	"fmt"
	"strings"

	"github.com/qnsoft/common/os/qn_time"
	qn_util "github.com/qnsoft/common/util/qn_util"

	"github.com/qnsoft/common/encoding/qn_html"
	"github.com/qnsoft/common/encoding/qn_url"
	qn_conv "github.com/qnsoft/common/util/qn_conv"

	htmltpl "html/template"
)

// funcDump implements build-in template function: dump
func (view *View) funcDump(values ...interface{}) (result string) {
	result += "<!--\n"
	for _, v := range values {
		result += qn_util.Export(v) + "\n"
	}
	result += "-->\n"
	return result
}

// funcEq implements build-in template function: eq
func (view *View) funcEq(value interface{}, others ...interface{}) bool {
	s := qn_conv.String(value)
	for _, v := range others {
		if strings.Compare(s, qn_conv.String(v)) != 0 {
			return false
		}
	}
	return true
}

// funcNe implements build-in template function: ne
func (view *View) funcNe(value, other interface{}) bool {
	return strings.Compare(qn_conv.String(value), qn_conv.String(other)) != 0
}

// funcLt implements build-in template function: lt
func (view *View) funcLt(value, other interface{}) bool {
	s1 := qn_conv.String(value)
	s2 := qn_conv.String(other)
	if qn.str.IsNumeric(s1) && qn.str.IsNumeric(s2) {
		return qn_conv.Int64(value) < qn_conv.Int64(other)
	}
	return strings.Compare(s1, s2) < 0
}

// funcLe implements build-in template function: le
func (view *View) funcLe(value, other interface{}) bool {
	s1 := qn_conv.String(value)
	s2 := qn_conv.String(other)
	if qn.str.IsNumeric(s1) && qn.str.IsNumeric(s2) {
		return qn_conv.Int64(value) <= qn_conv.Int64(other)
	}
	return strings.Compare(s1, s2) <= 0
}

// funcGt implements build-in template function: gt
func (view *View) funcGt(value, other interface{}) bool {
	s1 := qn_conv.String(value)
	s2 := qn_conv.String(other)
	if qn.str.IsNumeric(s1) && qn.str.IsNumeric(s2) {
		return qn_conv.Int64(value) > qn_conv.Int64(other)
	}
	return strings.Compare(s1, s2) > 0
}

// funcGe implements build-in template function: ge
func (view *View) funcGe(value, other interface{}) bool {
	s1 := qn_conv.String(value)
	s2 := qn_conv.String(other)
	if qn.str.IsNumeric(s1) && qn.str.IsNumeric(s2) {
		return qn_conv.Int64(value) >= qn_conv.Int64(other)
	}
	return strings.Compare(s1, s2) >= 0
}

// funcInclude implements build-in template function: include
// Note that configuration AutoEncode does not affect the output of this function.
func (view *View) funcInclude(file interface{}, data ...map[string]interface{}) htmltpl.HTML {
	var m map[string]interface{} = nil
	if len(data) > 0 {
		m = data[0]
	}
	path := qn_conv.String(file)
	if path == "" {
		return ""
	}
	// It will search the file internally.
	content, err := view.Parse(path, m)
	if err != nil {
		return htmltpl.HTML(err.Error())
	}
	return htmltpl.HTML(content)
}

// funcText implements build-in template function: text
func (view *View) funcText(html interface{}) string {
	return qn_html.StripTags(qn_conv.String(html))
}

// funcHtmlEncode implements build-in template function: html
func (view *View) funcHtmlEncode(html interface{}) string {
	return qn_html.Entities(qn_conv.String(html))
}

// funcHtmlDecode implements build-in template function: htmldecode
func (view *View) funcHtmlDecode(html interface{}) string {
	return qn_html.EntitiesDecode(qn_conv.String(html))
}

// funcUrlEncode implements build-in template function: url
func (view *View) funcUrlEncode(url interface{}) string {
	return qn_url.Encode(qn_conv.String(url))
}

// funcUrlDecode implements build-in template function: urldecode
func (view *View) funcUrlDecode(url interface{}) string {
	if content, err := qn_url.Decode(qn_conv.String(url)); err == nil {
		return content
	} else {
		return err.Error()
	}
}

// funcDate implements build-in template function: date
func (view *View) funcDate(format interface{}, timestamp ...interface{}) string {
	t := int64(0)
	if len(timestamp) > 0 {
		t = qn_conv.Int64(timestamp[0])
	}
	if t == 0 {
		t = qn_time.Timestamp()
	}
	return qn_time.NewFromTimeStamp(t).Format(qn_conv.String(format))
}

// funcCompare implements build-in template function: compare
func (view *View) funcCompare(value1, value2 interface{}) int {
	return strings.Compare(qn_conv.String(value1), qn_conv.String(value2))
}

// funcSubStr implements build-in template function: substr
func (view *View) funcSubStr(start, end, str interface{}) string {
	return qn.str.SubStrRune(qn_conv.String(str), qn_conv.Int(start), qn_conv.Int(end))
}

// funcStrLimit implements build-in template function: strlimit
func (view *View) funcStrLimit(length, suffix, str interface{}) string {
	return qn.str.StrLimitRune(qn_conv.String(str), qn_conv.Int(length), qn_conv.String(suffix))
}

// funcConcat implements build-in template function: concat
func (view *View) funcConcat(str ...interface{}) string {
	var s string
	for _, v := range str {
		s += qn_conv.String(v)
	}
	return s
}

// funcReplace implements build-in template function: replace
func (view *View) funcReplace(search, replace, str interface{}) string {
	return qn.str.Replace(qn_conv.String(str), qn_conv.String(search), qn_conv.String(replace), -1)
}

// funcHighlight implements build-in template function: highlight
func (view *View) funcHighlight(key, color, str interface{}) string {
	return qn.str.Replace(qn_conv.String(str), qn_conv.String(key), fmt.Sprintf(`<span style="color:%v;">%v</span>`, color, key))
}

// funcHideStr implements build-in template function: hidestr
func (view *View) funcHideStr(percent, hide, str interface{}) string {
	return qn.str.HideStr(qn_conv.String(str), qn_conv.Int(percent), qn_conv.String(hide))
}

// funcToUpper implements build-in template function: toupper
func (view *View) funcToUpper(str interface{}) string {
	return qn.str.ToUpper(qn_conv.String(str))
}

// funcToLower implements build-in template function: toupper
func (view *View) funcToLower(str interface{}) string {
	return qn.str.ToLower(qn_conv.String(str))
}

// funcNl2Br implements build-in template function: nl2br
func (view *View) funcNl2Br(str interface{}) string {
	return qn.str.Nl2Br(qn_conv.String(str))
}
