// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// static service testing.

package qn_http_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Log(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		logDir := qn_file.TempDir(qn_time.TimestampNanoStr())
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.BindHandler("/hello", func(r *qn_http.Request) {
			r.Response.Write("hello")
		})
		s.BindHandler("/error", func(r *qn_http.Request) {
			panic("custom error")
		})
		s.SetLogPath(logDir)
		s.SetAccessLogEnabled(true)
		s.SetErrorLogEnabled(true)
		s.SetLogStdout(false)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		defer qn_file.Remove(logDir)
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/hello"), "hello")
		t.Assert(client.GetContent("/error"), "custom error")

		logPath1 := qn_file.Join(logDir, qn_time.Now().Format("Y-m-d")+".log")
		t.Assert(qn.str.Contains(qn_file.GetContents(logPath1), "http server started listening on"), true)
		t.Assert(qn.str.Contains(qn_file.GetContents(logPath1), "HANDLER"), true)

		logPath2 := qn_file.Join(logDir, "access-"+qn_time.Now().Format("Ymd")+".log")
		//fmt.Println(qn_file.GetContents(logPath2))
		t.Assert(qn.str.Contains(qn_file.GetContents(logPath2), " /hello "), true)

		logPath3 := qn_file.Join(logDir, "error-"+qn_time.Now().Format("Ymd")+".log")
		//fmt.Println(qn_file.GetContents(logPath3))
		t.Assert(qn.str.Contains(qn_file.GetContents(logPath3), "custom error"), true)
	})
}
