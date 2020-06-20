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

	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Static_ServerRoot(t *testing.T) {
	// SetServerRoot with absolute path
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/index.htm", "index")
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "index")
		t.Assert(client.GetContent("/index.htm"), "index")
	})

	// SetServerRoot with relative path
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`static/test/%d`, p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/index.htm", "index")
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "index")
		t.Assert(client.GetContent("/index.htm"), "index")
	})
}

func Test_Static_ServerRoot_Security(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		s.SetServerRoot(qn_debug.TestDataPath("static1"))
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "index")
		t.Assert(client.GetContent("/index.htm"), "Not Found")
		t.Assert(client.GetContent("/index.html"), "index")
		t.Assert(client.GetContent("/test.html"), "test")
		t.Assert(client.GetContent("/../main.html"), "Not Found")
		t.Assert(client.GetContent("/..%2Fmain.html"), "Not Found")
	})
}

func Test_Static_Folder_Forbidden(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/test.html", "test")
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/index.html"), "Not Found")
		t.Assert(client.GetContent("/test.html"), "test")
	})
}

func Test_Static_IndexFolder(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/test.html", "test")
		s.SetIndexFolder(true)
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.AssertNE(client.GetContent("/"), "Forbidden")
		t.AssertNE(qn.str.Pos(client.GetContent("/"), `<a href="/test.html"`), -1)
		t.Assert(client.GetContent("/index.html"), "Not Found")
		t.Assert(client.GetContent("/test.html"), "test")
	})
}

func Test_Static_IndexFiles1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/index.html", "index")
		qn_file.PutContents(path+"/test.html", "test")
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "index")
		t.Assert(client.GetContent("/index.html"), "index")
		t.Assert(client.GetContent("/test.html"), "test")
	})
}

func Test_Static_IndexFiles2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/test.html", "test")
		s.SetIndexFiles([]string{"index.html", "test.html"})
		s.SetServerRoot(path)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "test")
		t.Assert(client.GetContent("/index.html"), "Not Found")
		t.Assert(client.GetContent("/test.html"), "test")
	})
}

func Test_Static_AddSearchPath1(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path1 := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		path2 := fmt.Sprintf(`%s/qn_http/static/test/%d/%d`, qn_file.TempDir(), p, p)
		defer qn_file.Remove(path1)
		defer qn_file.Remove(path2)
		qn_file.PutContents(path2+"/test.html", "test")
		s.SetServerRoot(path1)
		s.AddSearchPath(path2)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/test.html"), "test")
	})
}

func Test_Static_AddSearchPath2(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path1 := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		path2 := fmt.Sprintf(`%s/qn_http/static/test/%d/%d`, qn_file.TempDir(), p, p)
		defer qn_file.Remove(path1)
		defer qn_file.Remove(path2)
		qn_file.PutContents(path1+"/test.html", "test1")
		qn_file.PutContents(path2+"/test.html", "test2")
		s.SetServerRoot(path1)
		s.AddSearchPath(path2)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/test.html"), "test1")
	})
}

func Test_Static_AddStaticPath(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path1 := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		path2 := fmt.Sprintf(`%s/qn_http/static/test/%d/%d`, qn_file.TempDir(), p, p)
		defer qn_file.Remove(path1)
		defer qn_file.Remove(path2)
		qn_file.PutContents(path1+"/test.html", "test1")
		qn_file.PutContents(path2+"/test.html", "test2")
		s.SetServerRoot(path1)
		s.AddStaticPath("/my-test", path2)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/test.html"), "test1")
		t.Assert(client.GetContent("/my-test/test.html"), "test2")
	})
}

func Test_Static_AddStaticPath_Priority(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path1 := fmt.Sprintf(`%s/qn_http/static/test/%d/test`, qn_file.TempDir(), p)
		path2 := fmt.Sprintf(`%s/qn_http/static/test/%d/%d/test`, qn_file.TempDir(), p, p)
		defer qn_file.Remove(path1)
		defer qn_file.Remove(path2)
		qn_file.PutContents(path1+"/test.html", "test1")
		qn_file.PutContents(path2+"/test.html", "test2")
		s.SetServerRoot(path1)
		s.AddStaticPath("/test", path2)
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/test.html"), "test1")
		t.Assert(client.GetContent("/test/test.html"), "test2")
	})
}

func Test_Static_Rewrite(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		p, _ := ports.PopRand()
		s := g.Server(p)
		path := fmt.Sprintf(`%s/qn_http/static/test/%d`, qn_file.TempDir(), p)
		defer qn_file.Remove(path)
		qn_file.PutContents(path+"/test1.html", "test1")
		qn_file.PutContents(path+"/test2.html", "test2")
		s.SetServerRoot(path)
		s.SetRewrite("/test.html", "/test1.html")
		s.SetRewriteMap(g.MapStrStr{
			"/my-test1": "/test1.html",
			"/my-test2": "/test2.html",
		})
		s.SetPort(p)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Forbidden")
		t.Assert(client.GetContent("/test.html"), "test1")
		t.Assert(client.GetContent("/test1.html"), "test1")
		t.Assert(client.GetContent("/test2.html"), "test2")
		t.Assert(client.GetContent("/my-test1"), "test1")
		t.Assert(client.GetContent("/my-test2"), "test2")
	})
}
