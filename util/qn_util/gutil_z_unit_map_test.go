// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_util_test

import (
	"testing"

	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_util"
)

func Test_MapCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := qn_util.MapCopy(m1)
		m2["k2"] = "v2"

		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], "v1")
		t.Assert(m2["k2"], "v2")
	})
}

func Test_MapContains(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		t.Assert(qn_util.MapContains(m1, "k1"), true)
		t.Assert(qn_util.MapContains(m1, "K1"), false)
		t.Assert(qn_util.MapContains(m1, "k2"), false)
	})
}

func Test_MapMerge(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		qn_util.MapMerge(m1, m2, m3, nil)
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], "v2")
		t.Assert(m1["k3"], "v3")
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapMergeCopy(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		m := qn_util.MapMergeCopy(m1, m2, m3, nil)
		t.Assert(m["k1"], "v1")
		t.Assert(m["k2"], "v2")
		t.Assert(m["k3"], "v3")
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapPossibleItemByKey(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		k, v := qn_util.MapPossibleItemByKey(m, "NAME")
		t.Assert(k, "name")
		t.Assert(v, "guo")

		k, v = qn_util.MapPossibleItemByKey(m, "nick name")
		t.Assert(k, "NickName")
		t.Assert(v, "john")

		k, v = qn_util.MapPossibleItemByKey(m, "none")
		t.Assert(k, "")
		t.Assert(v, nil)
	})
}

func Test_MapContainsPossibleKey(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		t.Assert(qn_util.MapContainsPossibleKey(m, "name"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "NAME"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "nickname"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "nick name"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "nick_name"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "nick-name"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "nick.name"), true)
		t.Assert(qn_util.MapContainsPossibleKey(m, "none"), false)
	})
}
