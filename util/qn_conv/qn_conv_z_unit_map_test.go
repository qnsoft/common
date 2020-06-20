// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_conv_test

import (
	"testing"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/util/qn_conv"
	"github.com/qnsoft/common/util/qn_util"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_Map_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m1 := map[string]string{
			"k": "v",
		}
		m2 := map[int]string{
			3: "v",
		}
		m3 := map[float64]float32{
			1.22: 3.1,
		}
		t.Assert(qn_conv.Map(m1), qn.Map{
			"k": "v",
		})
		t.Assert(qn_conv.Map(m2), qn.Map{
			"3": "v",
		})
		t.Assert(qn_conv.Map(m3), qn.Map{
			"1.22": "3.1",
		})
	})
}

func Test_Map_Slice(t *testing.T) {
	qn_test.C(tqn.Slnc(t *qn_test.T) {
		slice1 := qn.Slice{"1", "2", "3", "4"}
		slice2 := qn.Slice{"1", "2", "3"}
		slice3 := g.Slice{}
		t.Assert(qn_conv.Map(slice1), qn.Map{
			"1": "2",
			"3": "4",
		})
		t.Assert(qn_conv.Map(slice2), qn.Map{
			"1": "2",
			"3": nil,
		})
		t.Assert(qn_conv.Map(slice3), qn.Map{})
	})
}

func Test_Mqn.SlBasic(t *testing.T) {
	params := g.Slice{
		qn.Map{"id": 100, "name": "john"},
		qn.Map{"id": 200, "name": "smith"},
	}
	qn_test.C(t, func(t *qn_test.T) {
		list := qn_conv.Maps(params)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
}

func Test_Maps_JsonStr(t *testing.T) {
	jsonStr := `[{"id":100, "name":"john"},{"id":200, "name":"smith"}]`
	qn_test.C(t, func(t *qn_test.T) {
		list := qn_conv.Maps(jsonStr)
		t.Assert(len(list), 2)
		t.Assert(list[0]["id"], 100)
		t.Assert(list[1]["id"], 200)
	})
}

func Test_Map_StructWithqn_convTag(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `qn_conv:"-"`
			NickName string `qn_conv:"nickname, omitempty"`
			Pass1    string `qn_conv:"password1"`
			Pass2    string `qn_conv:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := qn_conv.Map(user1)
		map2 := qn_conv.Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithJsonTag(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `json:"-"`
			NickName string `json:"nickname, omitempty"`
			Pass1    string `json:"password1"`
			Pass2    string `json:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := qn_conv.Map(user1)
		map2 := qn_conv.Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithCTag(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `c:"-"`
			NickName string `c:"nickname, omitempty"`
			Pass1    string `c:"password1"`
			Pass2    string `c:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := qn_conv.Map(user1)
		map2 := qn_conv.Map(user2)
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["SiteUrl"], nil)
		t.Assert(map1["NickName"], nil)
		t.Assert(map1["nickname"], nil)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		t.Assert(map2["Uid"], 100)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["SiteUrl"], nil)
		t.Assert(map2["NickName"], nil)
		t.Assert(map2["nickname"], nil)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")
	})
}

func Test_Map_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	qn_test.C(t, func(t *qn_test.T) {
		user := &User{1, "john"}
		t.Assert(qn_conv.Map(user), qn.Map{"Id": 1})
	})
}

func Test_MapDeep1(t *testing.T) {
	type Ids struct {
		Id  int `c:"id"`
		Uid int `c:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `c:"create_time"`
	}
	type User struct {
		Base
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}
	qn_test.C(t, func(t *qn_test.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := qn_conv.Map(user)
		t.Assert(m["id"], "")
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], "")
	})
	qn_test.C(t, func(t *qn_test.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := qn_conv.MapDeep(user)
		t.Assert(m["id"], user.Id)
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], user.CreateTime)
	})
}

func Test_MapDeep2(t *testing.T) {
	type A struct {
		F string
		G string
	}

	type B struct {
		A
		H string
	}

	type C struct {
		A A
		F string
	}

	type D struct {
		I A
		F string
	}

	qn_test.C(t, func(t *qn_test.T) {
		b := new(B)
		c := new(C)
		d := new(D)
		mb := qn_conv.MapDeep(b)
		mc := qn_conv.MapDeep(c)
		md := qn_conv.MapDeep(d)
		t.Assert(qn_util.MapContains(mb, "F"), true)
		t.Assert(qn_util.MapContains(mb, "G"), true)
		t.Assert(qn_util.MapContains(mb, "H"), true)
		t.Assert(qn_util.MapContains(mc, "A"), true)
		t.Assert(qn_util.MapContains(mc, "F"), true)
		t.Assert(qn_util.MapContains(mc, "G"), false)
		t.Assert(qn_util.MapContains(md, "F"), true)
		t.Assert(qn_util.MapContains(md, "I"), true)
		t.Assert(qn_util.MapContains(md, "H"), false)
		t.Assert(qn_util.MapContains(md, "G"), false)
	})
}

func Test_MapDeepWithAttributeTag(t *testing.T) {
	type Ids struct {
		Id  int `c:"id"`
		Uid int `c:"uid"`
	}
	type Base struct {
		Ids        `json:"ids"`
		CreateTime string `c:"create_time"`
	}
	type User struct {
		Base     `json:"base"`
		Passport string `c:"passport"`
		Password string `c:"password"`
		Nickname string `c:"nickname"`
	}
	qn_test.C(t, func(t *qn_test.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := qn_conv.Map(user)
		t.Assert(m["id"], "")
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["create_time"], "")
	})
	qn_test.C(t, func(t *qn_test.T) {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := qn_conv.MapDeep(user)
		t.Assert(m["base"].(map[string]interface{})["ids"].(map[string]interface{})["id"], user.Id)
		t.Assert(m["nickname"], user.Nickname)
		t.Assert(m["base"].(map[string]interface{})["create_time"], user.CreateTime)
	})
}
