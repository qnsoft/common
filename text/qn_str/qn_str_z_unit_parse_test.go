// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_str_test

import (
	"net/url"
	"testing"

	"github.com/gogf/gf/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/qn.str"
)

func Test_Parse(t *testing.T) {
	// url
	qn_test.C(t, func(t *qn_test.T) {
		s := "goframe.org/index?name=john&score=100"
		u, err := url.Parse(s)
		t.Assert(err, nil)
		m, err := qn.str.Parse(u.RawQuery)
		t.Assert(err, nil)
		t.Assert(m["name"], "john")
		t.Assert(m["score"], "100")

		// name overwrite
		m, err = qn.str.Parse("a=1&a=2")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"a": 2,
		})
		// slice
		m, err = qn.str.Parse("a[]=1&a[]=2")
		t.Assert(err, nil)
		t.Asseqn.Sl, qn.Map{
			"a": g.Slice{"1", "2"},
		})
		// map
		m, err = qn.str.Parse("a=1&b=2&c=3")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"a": "1",
			"b": "2",
			"c": "3",
		})
		m, err = qn.str.Parse("a=1&a=2&c=3")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"a": "2",
			"c": "3",
		})
		// map
		m, err = qn.str.Parse("m[a]=1&m[b]=2&m[c]=3")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"m": qn.Map{
				"a": "1",
				"b": "2",
				"c": "3",
			},
		})
		m, err = qn.str.Parse("m[a]=1&m[a]=2&m[b]=3")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"m": qn.Map{
				"a": "2",
				"b": "3",
			},
		})
		// map - slice
		m, err = qn.str.Parse("m[a][]=1&m[a][]=2")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"m": qqn.Slp{
				"a": g.Slice{"1", "2"},
			},
		})
		m, err = qn.str.Parse("m[a][b][]=1&m[a][b][]=2")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"m": qn.Map{
				"a": qqn.Slp{
					"b": g.Slice{"1", "2"},
				},
			},
		})
		// map - complicated
		m, err = qn.str.Parse("m[a1][b1][c1][d1]=1&m[a2][b2]=2&m[a3][b3][c3]=3")
		t.Assert(err, nil)
		t.Assert(m, qn.Map{
			"m": qn.Map{
				"a1": qn.Map{
					"b1": qn.Map{
						"c1": qn.Map{
							"d1": "1",
						},
					},
				},
				"a2": qn.Map{
					"b2": "2",
				},
				"a3": qn.Map{
					"b3": qn.Map{
						"c3": "3",
					},
				},
			},
		})
	})
}
