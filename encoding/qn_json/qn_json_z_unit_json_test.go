// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_json_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/frame/qn"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/qn.str"
)

func Test_ToJson(t *testing.T) {
	type ModifyFieldInfoType struct {
		Id  int64  `json:"id"`
		New string `json:"new"`
	}
	type ModifyFieldInfosType struct {
		Duration ModifyFieldInfoType `json:"duration"`
		OMLevel  ModifyFieldInfoType `json:"om_level"`
	}

	type MediaRequestModifyInfo struct {
		Modify ModifyFieldInfosType `json:"modifyFieldInfos"`
		Field  ModifyFieldInfosType `json:"fieldInfos"`
		FeedID string               `json:"feed_id"`
		Vid    string               `json:"id"`
	}

	qn_test.C(t, func(t *qn_test.T) {
		jsonContent := `{"dataSetId":2001,"fieldInfos":{"duration":{"id":80079,"value":"59"},"om_level":{"id":2409,"value":"4"}},"id":"g0936lt1u0f","modifyFieldInfos":{"om_level":{"id":2409,"new":"4","old":""}},"timeStamp":1584599734}`
		var info MediaRequestModifyInfo
		err := qn_json.DecodeTo(jsonContent, &info)
		t.Assert(err, nil)
		content := qn_json.New(info).MustToJsonString()
		t.Assert(qn.str.Contains(content, `"feed_id":""`), true)
		t.Assert(qn.str.Contains(content, `"fieldInfos":{`), true)
		t.Assert(qn.str.Contains(content, `"id":80079`), true)
		t.Assert(qn.str.Contains(content, `"om_level":{`), true)
		t.Assert(qn.str.Contains(content, `"id":2409,`), true)
		t.Assert(qn.str.Contains(content, `"id":"g0936lt1u0f"`), true)
		t.Assert(qn.str.Contains(content, `"new":"4"`), true)
	})
}

func Test_MapAttributeConvert(t *testing.T) {
	var data = `
 {
   "title": {"l1":"标签1","l2":"标签2"}
}
`
	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		qn_test.Assert(err, nil)

		tx := struct {
			Title map[string]interface{}
		}{}

		err = j.ToStruct(&tx)
		qn_test.Assert(err, nil)
		t.Assert(tx.Title, qn.Map{
			"l1": "标签1", "l2": "标签2",
		})
	})

	qn_test.C(t, func(t *qn_test.T) {
		j, err := qn_json.LoadContent(data)
		qn_test.Assert(err, nil)

		tx := struct {
			Title map[string]string
		}{}

		err = j.ToStruct(&tx)
		qn_test.Assert(err, nil)
		t.Assert(tx.Title, qn.Map{
			"l1": "标签1", "l2": "标签2",
		})
	})
}
