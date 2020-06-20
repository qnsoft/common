// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/text/gstr"
	"github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/frame/g"
	"github.com/qnsoft/common/net/qn_http"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_ConfigFromMap(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := g.Map{
			"address":         ":8199",
			"readTimeout":     "60s",
			"indexFiles":      g.Slice{"index.php", "main.php"},
			"errorLogEnabled": true,
			"cookieMaxAge":    "1y",
		}
		config, err := qn_http.ConfigFromMap(m)
		t.Assert(err, nil)
		d1, _ := time.ParseDuration(qn_conv.String(m["readTimeout"]))
		d2, _ := time.ParseDuration(qn_conv.String(m["cookieMaxAge"]))
		t.Assert(config.Address, m["address"])
		t.Assert(config.ReadTimeout, d1)
		t.Assert(config.CookieMaxAge, d2)
		t.Assert(config.IndexFiles, m["indexFiles"])
		t.Assert(config.ErrorLogEnabled, m["errorLogEnabled"])
	})
}

func Test_SetConfigWithMap(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		m := g.Map{
			"Address": ":8199",
			//"ServerRoot":       "/var/www/MyServerRoot",
			"IndexFiles":       g.Slice{"index.php", "main.php"},
			"AccessLogEnabled": true,
			"ErrorLogEnabled":  true,
			"PProfEnabled":     true,
			"LogPath":          "/var/log/MyServerLog",
			"SessionIdName":    "MySessionId",
			"SessionPath":      "/tmp/MySessionStoragePath",
			"SessionMaxAge":    24 * time.Hour,
		}
		s := g.Server()
		err := s.SetConfigWithMap(m)
		t.Assert(err, nil)
	})
}

func Test_ClientMaxBodySize(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.POST("/", func(r *qn_http.Request) {
			r.Response.Write(r.GetBodyString())
		})
	})
	m := g.Map{
		"Address":           p,
		"ClientMaxBodySize": "1k",
	}
	qn_test.Assert(s.SetConfigWithMap(m), nil)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(
			gstr.Trim(c.PostContent("/", data)),
			data[:1024],
		)
	})
}

func Test_ClientMaxBodySize_File(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.POST("/", func(r *qn_http.Request) {
			r.GetUploadFile("file")
			r.Response.Write("ok")
		})
	})
	m := g.Map{
		"Address":           p,
		"ErrorLogEnabled":   false,
		"ClientMaxBodySize": "1k",
	}
	qn_test.Assert(s.SetConfigWithMap(m), nil)
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	// ok
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		data := make([]byte, 512)
		for i := 0; i < 512; i++ {
			data[i] = 'a'
		}
		t.Assert(qn_file.PutBytes(path, data), nil)
		defer qn_file.Remove(path)
		t.Assert(
			gstr.Trim(c.PostContent("/", "name=john&file=@file:"+path)),
			"ok",
		)
	})

	// too large
	qn_test.C(t, func(t *qn_test.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		data := make([]byte, 1056)
		for i := 0; i < 1056; i++ {
			data[i] = 'a'
		}
		t.Assert(qn_file.PutBytes(path, data), nil)
		defer qn_file.Remove(path)
		t.Assert(
			gstr.Trim(c.PostContent("/", "name=john&file=@file:"+path)),
			"http: request body too large",
		)
	})
}
