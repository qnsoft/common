// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_time_test

import (
	"testing"
	"time"

	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_SetTimeZone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_time.SetTimeZone("Asia/Shanghai")
		t.Assert(time.Local.String(), "Asia/Shanghai")
	})
}

func Test_Nanosecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		nanos := qn_time.TimestampNano()
		timeTemp := time.Unix(0, nanos)
		t.Assert(nanos, timeTemp.UnixNano())
	})
}

func Test_Microsecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		micros := qn_time.TimestampMicro()
		timeTemp := time.Unix(0, micros*1e3)
		t.Assert(micros, timeTemp.UnixNano()/1e3)
	})
}

func Test_Millisecond(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		millis := qn_time.TimestampMilli()
		timeTemp := time.Unix(0, millis*1e6)
		t.Assert(millis, timeTemp.UnixNano()/1e6)
	})
}

func Test_Second(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		s := qn_time.Timestamp()
		timeTemp := time.Unix(s, 0)
		t.Assert(s, timeTemp.Unix())
	})
}

func Test_Date(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_time.Date(), time.Now().Format("2006-01-02"))
	})
}

func Test_Datetime(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		datetime := qn_time.Datetime()
		timeTemp, err := qn_time.StrToTime(datetime, "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		t.Assert(datetime, timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}

func Test_ISO8601(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		iso8601 := qn_time.ISO8601()
		t.Assert(iso8601, qn_time.Now().Format("c"))
	})
}

func Test_RFC822(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		rfc822 := qn_time.RFC822()
		t.Assert(rfc822, qn_time.Now().Format("r"))
	})
}

func Test_StrToTime(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		//正常日期列表
		//正则的原因，日期"06.01.02"，"2006.01"，"2006..01"无法覆盖qn_time.go的百分百
		var testDatetimes = []string{
			"2006-01-02 15:04:05",
			"2006/01/02 15:04:05",
			"2006.01.02 15:04:05.000",
			"2006.01.02 - 15:04:05",
			"2006.01.02 15:04:05 +0800 CST",
			"2006-01-02T20:05:06+05:01:01",
			"2006-01-02T14:03:04Z01:01:01",
			"2006-01-02T15:04:05Z",
			"02-jan-2006 15:04:05",
			"02/jan/2006 15:04:05",
			"02.jan.2006 15:04:05",
			"02.jan.2006:15:04:05",
		}

		for _, item := range testDatetimes {
			timeTemp, err := qn_time.StrToTime(item)
			if err != nil {
				t.Error("test fail")
			}
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")
		}

		//正常日期列表，时间00:00:00
		var testDates = []string{
			"2006.01.02",
			"2006.01.02 00:00",
			"2006.01.02 00:00:00.000",
		}

		for _, item := range testDates {
			timeTemp, err := qn_time.StrToTime(item)
			if err != nil {
				t.Error("test fail")
			}
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 00:00:00")
		}

		//测试格式化formatToStdLayout
		var testDateFormats = []string{
			"Y-m-d H:i:s",
			"\\T\\i\\m\\e Y-m-d H:i:s",
			"Y-m-d H:i:s\\",
			"Y-m-j G:i:s.u",
			"Y-m-j G:i:su",
		}

		var testDateFormatsResult = []string{
			"2007-01-02 15:04:05",
			"Time 2007-01-02 15:04:05",
			"2007-01-02 15:04:05",
			"2007-01-02 15:04:05.000",
			"2007-01-02 15:04:05.000",
		}

		for index, item := range testDateFormats {
			timeTemp, err := qn_time.StrToTime(testDateFormatsResult[index], item)
			if err != nil {
				t.Error("test fail")
			}
			t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05.000"), "2007-01-02 15:04:05.000")
		}

		//异常日期列表
		var testDatesFail = []string{
			"2006.01",
			"06..02",
			"20060102",
		}

		for _, item := range testDatesFail {
			_, err := qn_time.StrToTime(item)
			if err == nil {
				t.Error("test fail")
			}
		}

		//test err
		_, err := qn_time.StrToTime("2006-01-02 15:04:05", "aabbccdd")
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_ConvertZone(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		//现行时间
		nowUTC := time.Now().UTC()
		testZone := "America/Los_Angeles"

		//转换为洛杉矶时间
		t1, err := qn_time.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err != nil {
			t.Error("test fail")
		}

		//使用洛杉矶时区解析上面转换后的时间
		laStr := t1.Time.Format("2006-01-02 15:04:05")
		loc, err := time.LoadLocation(testZone)
		t2, err := time.ParseInLocation("2006-01-02 15:04:05", laStr, loc)

		//判断是否与现行时间匹配
		t.Assert(t2.UTC().Unix(), nowUTC.Unix())

	})

	//test err
	qn_test.C(t, func(t *qn_test.T) {
		//现行时间
		nowUTC := time.Now().UTC()
		//t.Log(nowUTC.Unix())
		testZone := "errZone"

		//错误时间输入
		_, err := qn_time.ConvertZone(nowUTC.Format("06..02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		//错误时区输入
		_, err = qn_time.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, "")
		if err == nil {
			t.Error("test fail")
		}
		//错误时区输入
		_, err = qn_time.ConvertZone(nowUTC.Format("2006-01-02 15:04:05"), testZone, testZone)
		if err == nil {
			t.Error("test fail")
		}
	})
}

func Test_StrToTimeFormat(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {

	})
}

func Test_ParseTimeFromContent(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		timeTemp := qn_time.ParseTimeFromContent("我是中文2006-01-02 15:04:05我也是中文", "Y-m-d H:i:s")
		t.Assert(timeTemp.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp1 := qn_time.ParseTimeFromContent("我是中文2006-01-02 15:04:05我也是中文")
		t.Assert(timeTemp1.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		timeTemp2 := qn_time.ParseTimeFromContent("我是中文02.jan.2006 15:04:05我也是中文")
		t.Assert(timeTemp2.Time.Format("2006-01-02 15:04:05"), "2006-01-02 15:04:05")

		//test err
		timeTempErr := qn_time.ParseTimeFromContent("我是中文", "Y-m-d H:i:s")
		if timeTempErr != nil {
			t.Error("test fail")
		}
	})
}

func Test_FuncCost(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		qn_time.FuncCost(func() {

		})
	})
}
