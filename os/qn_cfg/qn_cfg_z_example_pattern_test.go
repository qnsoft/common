// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cfg_test

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/os/qn_cfg"
)

func Example_mapSliceChange() {
	intlog.SetEnabled(false)
	defer intlog.SetEnabled(true)
	// For testing/example only.
	content := `{"map":{"key":"value"}, "slice":[59,90]}`
	qn_cfg.SetContent(content)
	defer qn_cfg.RemoveContent()

	m := g.Cfg().GetMap("map")
	fmt.Println(m)

	// Change the key-value pair.
	m["key"] = "john"

	// It changes the underlying key-value pair.
	fmt.Println(g.Cfg().GetMap("map"))

	s := g.Cfg().GetArray("slice")
	fmt.Println(s)

	// Change the value of specified index.
	s[0] = 100

	// It changes the underlying slice.
	fmt.Println(g.Cfg().GetArray("slice"))

	// output:
	// map[key:value]
	// map[key:john]
	// [59 90]
	// [100 90]
}
