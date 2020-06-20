// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_type_test

import (
	"sync"
	"testing"

	"github.com/qnsoft/common/container/qn_type"
	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/util/qn_conv"
)

func Test_Uint(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var wg sync.WaitGroup
		addTimes := 1000
		i := qn_type.NewUint(0)
		iClone := i.Clone()
		t.AssertEQ(iClone.Set(1), uint(0))
		t.AssertEQ(iClone.Val(), uint(1))
		for index := 0; index < addTimes; index++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i.Add(1)
			}()
		}
		wg.Wait()
		t.AssertEQ(uint(addTimes), i.Val())

		//空参测试
		i1 := qn_type.NewUint()
		t.AssertEQ(i1.Val(), uint(0))
	})
}

func Test_Uint_JSON(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		i := qn_type.NewUint(666)
		b1, err1 := json.Marshal(i)
		b2, err2 := json.Marshal(i.Val())
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.Assert(b1, b2)

		i2 := qn_type.NewUint()
		err := json.Unmarshal(b2, &i2)
		t.Assert(err, nil)
		t.Assert(i2.Val(), i)
	})
}

func Test_Uint_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Var  *qn_type.Uint
	}
	qn_test.C(t, func(t *qn_test.T) {
		var v *V
		err := qn_conv.Struct(map[string]interface{}{
			"name": "john",
			"var":  "123",
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.Val(), "123")
	})
}
