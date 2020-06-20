// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_valid

import (
	"strconv"
	"strings"

	"github.com/qnsoft/common/util/qn_conv"
)

// checkLength checks <value> using length rules.
// The length is calculated using unicode string, which means one chinese character or letter
// both has the length of 1.
func checkLength(value, ruleKey, ruleVal string, customMsqn_map map[string]string) string {
	var (
		msg       = ""
		runeArray = qn_conv.Runes(value)
		valueLen  = len(runeArray)
	)
	switch ruleKey {
	case "length":
		var (
			min   = 0
			max   = 0
			array = strings.Split(ruleVal, ",")
		)
		if len(array) > 0 {
			if v, err := strconv.Atoi(strings.TrimSpace(array[0])); err == nil {
				min = v
			}
		}
		if len(array) > 1 {
			if v, err := strconv.Atoi(strings.TrimSpace(array[1])); err == nil {
				max = v
			}
		}
		if valueLen < min || valueLen > max {
			msg = getErrorMessageByRule(ruleKey, customMsqn_map)
			msg = strings.Replace(msg, ":min", strconv.Itoa(min), -1)
			msg = strings.Replace(msg, ":max", strconv.Itoa(max), -1)
			return msg
		}

	case "min-length":
		min, err := strconv.Atoi(ruleVal)
		if valueLen < min || err != nil {
			msg = getErrorMessageByRule(ruleKey, customMsqn_map)
			msg = strings.Replace(msg, ":min", strconv.Itoa(min), -1)
		}

	case "max-length":
		max, err := strconv.Atoi(ruleVal)
		if valueLen > max || err != nil {
			msg = getErrorMessageByRule(ruleKey, customMsqn_map)
			msg = strings.Replace(msg, ":max", strconv.Itoa(max), -1)
		}
	}
	return msg
}
