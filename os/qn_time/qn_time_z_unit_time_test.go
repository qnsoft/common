// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_time_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_New(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeNow := time.Now()
		timeTemp := qn_time.New(timeNow)
		t.Assert(timeTemp.Time.UnixNano(), timeNow.UnixNano())

		timeTemp1 := qn_time.New()
		t.Assert(timeTemp1.Time, time.Time{})
	})
}

func Test_Nil(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		var t1 *qn_time.Time
		t.Assert(t1.String(), "")
	})
	qn_test.C(t, func(t *qn_test.T) {
		var t1 qn_time.Time
		t.Assert(t1.String(), "")
	})
}

func Test_NewFromStr(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromStr("2006-01-02 15:04:05")
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := qn_time.NewFromStr("20060102")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_String(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t1 := qn_time.NewFromStr("2006-01-02 15:04:05")
		t.Assert(t1.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("%s", t1), "2006-01-02 15:04:05")

		t2 := *t1
		t.Assert(t2.String(), "2006-01-02 15:04:05")
		t.Assert(fmt.Sprintf("%s", t2), "{2006-01-02 15:04:05}")
	})
}

func Test_NewFromStrFormat(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromStrFormat("2006-01-02 15:04:05", "Y-m-d H:i:s")
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := qn_time.NewFromStrFormat("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})

	qn_test.C(t, func(t *qn_test.T) {
		t1 := qn_time.NewFromStrFormat("2019/2/1", "Y/n/j")
		t.Assert(t1.Format("Y-m-d"), "2019-02-01")

		t2 := qn_time.NewFromStrFormat("2019/10/12", "Y/n/j")
		t.Assert(t2.Format("Y-m-d"), "2019-10-12")
	})
}

func Test_NewFromStrLayout(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromStrLayout("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2006-01-02 15:04:05")

		timeTemp1 := qn_time.NewFromStrLayout("2006-01-02 15:04:05", "aabbcc")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_NewFromTimeStamp(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromTimeStamp(1554459846000)
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2019-04-05 18:24:06")
		timeTemp1 := qn_time.NewFromTimeStamp(0)
		t.Assert(timeTemp1.Format("Y-m-d H:i:s"), "0001-01-01 00:00:00")
	})
}

func Test_Time_Second(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		t.Assert(timeTemp.Second(), timeTemp.Time.Second())
	})
}

func Test_Time_Nanosecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		t.Assert(timeTemp.Nanosecond(), timeTemp.Time.Nanosecond())
	})
}

func Test_Time_Microsecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		t.Assert(timeTemp.Microsecond(), timeTemp.Time.Nanosecond()/1e3)
	})
}

func Test_Time_Millisecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		t.Assert(timeTemp.Millisecond(), timeTemp.Time.Nanosecond()/1e6)
	})
}

func Test_Time_String(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		t.Assert(timeTemp.String(), timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}

func Test_Time_ISO8601(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		now := qn_time.Now()
		t.Assert(now.ISO8601(), now.Format("c"))
	})
}

func Test_Time_RFC822(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		now := qn_time.Now()
		t.Assert(now.RFC822(), now.Format("r"))
	})
}

func Test_Clone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Clone()
		t.Assert(timeTemp.Time.Unix(), timeTemp1.Time.Unix())
	})
}

func Test_ToTime(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Time
		t.Assert(timeTemp.Time.UnixNano(), timeTemp1.UnixNano())
	})
}

func Test_Add(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromStr("2006-01-02 15:04:05")
		timeTemp.Add(time.Second)
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2006-01-02 15:04:06")
	})
}

func Test_ToZone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		//
		timeTemp.ToZone("America/Los_Angeles")
		t.Assert(timeTemp.Time.Location().String(), "America/Los_Angeles")

		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			t.Error("test fail")
		}
		timeTemp.ToLocation(loc)
		t.Assert(timeTemp.Time.Location().String(), "Asia/Shanghai")

		timeTemp1, _ := timeTemp.ToZone("errZone")
		if timeTemp1 != nil {
			t.Error("test fail")
		}
	})
}

func Test_AddDate(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.NewFromStr("2006-01-02 15:04:05")
		timeTemp.AddDate(1, 2, 3)
		t.Assert(timeTemp.Format("Y-m-d H:i:s"), "2007-03-05 15:04:05")
	})
}

func Test_UTC(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Time
		timeTemp.UTC()
		t.Assert(timeTemp.UnixNano(), timeTemp1.UTC().UnixNano())
	})
}

func Test_Local(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Time
		timeTemp.Local()
		t.Assert(timeTemp.UnixNano(), timeTemp1.Local().UnixNano())
	})
}

func Test_Round(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Time
		timeTemp.Round(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Round(time.Hour).UnixNano())
	})
}

func Test_Truncate(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.Now()
		timeTemp1 := timeTemp.Time
		timeTemp.Truncate(time.Hour)
		t.Assert(timeTemp.UnixNano(), timeTemp1.Truncate(time.Hour).UnixNano())
	})
}
