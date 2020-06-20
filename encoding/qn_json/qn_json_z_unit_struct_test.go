// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_GetScan(t *testing.T) {
	type User struct {
		Name  string
		Score float64
	}
	j := qn_json.New(`[{"name":"john", "score":"100"},{"name":"smith", "score":"60"}]`)
	qn_test.C(t, func(t *qn_test.T) {
		var user *User
		err := j.GetScan("1", &user)
		t.Assert(err, nil)
		t.Assert(user, &User{
			Name:  "smith",
			Score: 60,
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var users []User
		err := j.GetScan(".", &users)
		t.Assert(err, nil)
		t.Assert(users, []User{
			{
				Name:  "john",
				Score: 100,
			},
			{
				Name:  "smith",
				Score: 60,
			},
		})
	})
}

func Test_GetScanDeep(t *testing.T) {
	type User struct {
		Name  string
		Score float64
	}
	j := qn_json.New(`[{"name":"john", "score":"100"},{"name":"smith", "score":"60"}]`)
	qn_test.C(t, func(t *qn_test.T) {
		var user *User
		err := j.GetScanDeep("1", &user)
		t.Assert(err, nil)
		t.Assert(user, &User{
			Name:  "smith",
			Score: 60,
		})
	})
	qn_test.C(t, func(t *qn_test.T) {
		var users []User
		err := j.GetScanDeep(".", &users)
		t.Assert(err, nil)
		t.Assert(users, []User{
			{
				Name:  "john",
				Score: 100,
			},
			{
				Name:  "smith",
				Score: 60,
			},
		})
	})
}

func Test_ToScan(t *testing.T) {
	type User struct {
		Name  string
		Score float64
	}
	j := qn_json.New(`[{"name":"john", "score":"100"},{"name":"smith", "score":"60"}]`)
	qn_test.C(t, func(t *qn_test.T) {
		var users []User
		err := j.ToScan(&users)
		t.Assert(err, nil)
		t.Assert(users, []User{
			{
				Name:  "john",
				Score: 100,
			},
			{
				Name:  "smith",
				Score: 60,
			},
		})
	})
}

func Test_ToScanDeep(t *testing.T) {
	type User struct {
		Name  string
		Score float64
	}
	j := qn_json.New(`[{"name":"john", "score":"100"},{"name":"smith", "score":"60"}]`)
	qn_test.C(t, func(t *qn_test.T) {
		var users []User
		err := j.ToScanDeep(&users)
		t.Assert(err, nil)
		t.Assert(users, []User{
			{
				Name:  "john",
				Score: 100,
			},
			{
				Name:  "smith",
				Score: 60,
			},
		})
	})
}

func Test_ToStruct1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type BaseInfoItem struct {
			IdCardNumber        string `db:"id_card_number" json:"idCardNumber" field:"id_card_number"`
			IsHouseholder       bool   `db:"is_householder" json:"isHouseholder" field:"is_householder"`
			HouseholderRelation string `db:"householder_relation" json:"householderRelation" field:"householder_relation"`
			UserName            string `db:"user_name" json:"userName" field:"user_name"`
			UserSex             string `db:"user_sex" json:"userSex" field:"user_sex"`
			UserAge             int    `db:"user_age" json:"userAge" field:"user_age"`
			UserNation          string `db:"user_nation" json:"userNation" field:"user_nation"`
		}

		type UserCollectionAddReq struct {
			BaseInfo []BaseInfoItem `db:"_" json:"baseInfo" field:"_"`
		}
		jsonContent := `{
	"baseInfo": [{
		"idCardNumber": "520101199412141111",
		"isHouseholder": true,
		"householderRelation": "户主",
		"userName": "李四",
		"userSex": "男",
		"userAge": 32,
		"userNation": "苗族",
		"userPhone": "13084183323",
		"liveAddress": {},
		"occupationInfo": [{
			"occupationType": "经商",
			"occupationBusinessInfo": [{
				"occupationClass": "制造业",
				"businessLicenseNumber": "32020000012300",
				"businessName": "土灶柴火鸡",
				"spouseName": "",
				"spouseIdCardNumber": "",
				"businessLicensePhotoId": 125,
				"businessPlace": "租赁房产",
				"hasGoodsInsurance": true,
				"businessScopeStr": "柴火鸡;烧烤",
				"businessAddress": {},
				"businessPerformAbility": {
					"businessType": "服务业",
					"businessLife": 5,
					"salesRevenue": 8000,
					"familyEquity": 6000
				}
			}],
			"occupationWorkInfo": {
				"occupationClass": "",
				"companyName": "",
				"companyType": "",
				"workYearNum": 0,
				"spouseName": "",
				"spouseIdCardNumber": "",
				"spousePhone": "",
				"spouseEducation": "",
				"spouseCompanyName": "",
				"workLevel": "",
				"workAddress": {},
				"workPerformAbility": {
					"familyAnnualIncome": 0,
					"familyEquity": 0,
					"workCooperationState": "",
					"workMoneyCooperationState": ""
				}
			},
			"occupationAgricultureInfo": []
		}],
		"assetsInfo": [],
		"expenditureInfo": [],
		"incomeInfo": [],
		"liabilityInfo": []
	}]
}`
		data := new(UserCollectionAddReq)
		j, err := qn_json.LoadJson(jsonContent)
		t.Assert(err, nil)
		err = j.ToStruct(data)
		t.Assert(err, nil)
	})
}

func Test_ToStructDeep(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		type Item struct {
			Title string `json:"title"`
			Key   string `json:"key"`
		}

		type M struct {
			Id    string                 `json:"id"`
			Me    map[string]interface{} `json:"me"`
			Txt   string                 `json:"txt"`
			Items []*Item                `json:"items"`
		}

		txt := `{
		  "id":"88888",
		  "me":{"name":"mikey","day":"20009"},
		  "txt":"hello",
		  "items":null
		 }`

		j, err := qn_json.LoadContent(txt)
		t.Assert(err, nil)
		t.Assert(j.GetString("me.name"), "mikey")
		t.Assert(j.GetString("items"), "")
		t.Assert(j.GetBool("items"), false)
		t.Assert(j.GetArray("items"), nil)
		m := new(M)
		err = j.ToStructDeep(m)
		t.Assert(err, nil)
		t.AssertNE(m.Me, nil)
		t.Assert(m.Me["day"], "20009")
		t.Assert(m.Items, nil)
	})
}
