// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/debug/qn_debug"
	"github.com/qnsoft/common/frame/g"

	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_BindMiddleware_Basic1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("/test", func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	}, func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.BindMiddleware("/test/:name", func(r *qn_http.Request) {
		r.Response.Write("5")
		r.Middleware.Next()
		r.Response.Write("6")
	}, func(r *qn_http.Request) {
		r.Response.Write("7")
		r.Middleware.Next()
		r.Response.Write("8")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/test"), "1342")
		t.Assert(client.GetContent("/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("/*", func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "12")
		t.Assert(client.GetContent("/test"), "12")
		t.Assert(client.GetContent("/test/test"), "1test2")
	})
}

func Test_BindMiddleware_Basic3(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("PUT:/test", func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	}, func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.BindMiddleware("POST:/test/:name", func(r *qn_http.Request) {
		r.Response.Write("5")
		r.Middleware.Next()
		r.Response.Write("6")
	}, func(r *qn_http.Request) {
		r.Response.Write("7")
		r.Middleware.Next()
		r.Response.Write("8")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/test"), "Not Found")
		t.Assert(client.PutContent("/test"), "1342")
		t.Assert(client.PostContent("/test"), "Not Found")
		t.Assert(client.GetContent("/test/test"), "test")
		t.Assert(client.PutContent("/test/test"), "test")
		t.Assert(client.PostContent("/test/test"), "57test86")
	})
}

func Test_BindMiddleware_Basic4(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Middleware(func(r *qn_http.Request) {
			r.Response.Write("1")
			r.Middleware.Next()
		})
		group.Middleware(func(r *qn_http.Request) {
			r.Middleware.Next()
			r.Response.Write("2")
		})
		group.ALL("/test", func(r *qn_http.Request) {
			r.Response.Write("test")
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/test"), "1test2")
		t.Assert(client.PutContent("/test/none"), "Not Found")
	})
}

func Test_Middleware_With_Static(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Middleware(func(r *qn_http.Request) {
			r.Response.Write("1")
			r.Middleware.Next()
			r.Response.Write("2")
		})
		group.ALL("/user/list", func(r *qn_http.Request) {
			r.Response.Write("list")
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.SetServerRoot(qn_debug.TestDataPath("static1"))
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "index")
		t.Assert(client.GetContent("/test.html"), "test")
		t.Assert(client.GetContent("/none"), "Not Found")
		t.Assert(client.GetContent("/user/list"), "1list2")
	})
}

func Test_Middleware_Status(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Middleware(func(r *qn_http.Request) {
			r.Middleware.Next()
			r.Response.WriteOver(r.Response.Status)
		})
		group.ALL("/user/list", func(r *qn_http.Request) {
			r.Response.Write("list")
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/user/list"), "200")

		resp, err := client.Get("/")
		defer resp.Close()
		t.Assert(err, nil)
		t.Assert(resp.StatusCode, 404)
	})
}

func Test_Middleware_Hook_With_Static(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	a := qn_array.New(true)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Hook("/*", qn_http.HOOK_BEFORE_SERVE, func(r *qn_http.Request) {
			a.Append(1)
			fmt.Println("HOOK_BEFORE_SERVE")
			r.Response.Write("a")
		})
		group.Hook("/*", qn_http.HOOK_AFTER_SERVE, func(r *qn_http.Request) {
			a.Append(1)
			fmt.Println("HOOK_AFTER_SERVE")
			r.Response.Write("b")
		})
		group.Middleware(func(r *qn_http.Request) {
			r.Response.Write("1")
			r.Middleware.Next()
			r.Response.Write("2")
		})
		group.ALL("/user/list", func(r *qn_http.Request) {
			r.Response.Write("list")
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.SetServerRoot(qn_debug.TestDataPath("static1"))
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		// The length assert sometimes fails, so I added time.Sleep here for debug purpose.

		t.Assert(client.GetContent("/"), "index")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.Len(), 2)

		t.Assert(client.GetContent("/test.html"), "test")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.Len(), 4)

		t.Assert(client.GetContent("/none"), "ab")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.Len(), 6)

		t.Assert(client.GetContent("/user/list"), "a1list2b")
		time.Sleep(100 * time.Millisecond)
		t.Assert(a.Len(), 8)
	})
}

func Test_BindMiddleware_Status(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddleware("/test/*any", func(r *qn_http.Request) {
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/test"), "Not Found")
		t.Assert(client.GetContent("/test/test"), "test")
		t.Assert(client.GetContent("/test/test/test"), "Not Found")
	})
}

func Test_BindMiddlewareDefault_Basic1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "1342")
		t.Assert(client.GetContent("/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic2(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("PUT:/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "1342")
		t.Assert(client.PutContent("/"), "1342")
		t.Assert(client.GetContent("/test/test"), "1342")
		t.Assert(client.PutContent("/test/test"), "13test42")
	})
}

func Test_BindMiddlewareDefault_Basic3(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "12")
		t.Assert(client.GetContent("/test/test"), "1test2")
	})
}

func Test_BindMiddlewareDefault_Basic4(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Middleware.Next()
		r.Response.Write("1")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("2")
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "21")
		t.Assert(client.GetContent("/test/test"), "2test1")
	})
}

func Test_BindMiddlewareDefault_Basic5(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("2")
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "12")
		t.Assert(client.GetContent("/test/test"), "12test")
	})
}

func Test_BindMiddlewareDefault_Status(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Middleware.Next()
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/test/test"), "test")
	})
}

type ObjectMiddleware struct{}

func (o *ObjectMiddleware) Init(r *qn_http.Request) {
	r.Response.Write("100")
}

func (o *ObjectMiddleware) Shut(r *qn_http.Request) {
	r.Response.Write("200")
}

func (o *ObjectMiddleware) Index(r *qn_http.Request) {
	r.Response.Write("Object Index")
}

func (o *ObjectMiddleware) Show(r *qn_http.Request) {
	r.Response.Write("Object Show")
}

func (o *ObjectMiddleware) Info(r *qn_http.Request) {
	r.Response.Write("Object Info")
}

func Test_BindMiddlewareDefault_Basic6(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindObject("/", new(ObjectMiddleware))
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "13100Object Index20042")
		t.Assert(client.GetContent("/init"), "1342")
		t.Assert(client.GetContent("/shut"), "1342")
		t.Assert(client.GetContent("/index"), "13100Object Index20042")
		t.Assert(client.GetContent("/show"), "13100Object Show20042")
		t.Assert(client.GetContent("/none-exist"), "1342")
	})
}

func Test_Hook_Middleware_Basic1(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/test/test", func(r *qn_http.Request) {
		r.Response.Write("test")
	})
	s.BindHookHandler("/*", qn_http.HOOK_BEFORE_SERVE, func(r *qn_http.Request) {
		r.Response.Write("a")
	})
	s.BindHookHandler("/*", qn_http.HOOK_AFTER_SERVE, func(r *qn_http.Request) {
		r.Response.Write("b")
	})
	s.BindHookHandler("/*", qn_http.HOOK_BEFORE_SERVE, func(r *qn_http.Request) {
		r.Response.Write("c")
	})
	s.BindHookHandler("/*", qn_http.HOOK_AFTER_SERVE, func(r *qn_http.Request) {
		r.Response.Write("d")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("1")
		r.Middleware.Next()
		r.Response.Write("2")
	})
	s.BindMiddlewareDefault(func(r *qn_http.Request) {
		r.Response.Write("3")
		r.Middleware.Next()
		r.Response.Write("4")
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "ac1342bd")
		t.Assert(client.GetContent("/test/test"), "ac13test42bd")
	})
}

func MiddlewareAuth(r *qn_http.Request) {
	token := r.Get("token")
	if token == "123456" {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func MiddlewareCORS(r *qn_http.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func Test_Middleware_CORSAndAuth(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Use(MiddlewareCORS)
	s.Group("/api.v2", func(group *qn_http.RouterGroup) {
		group.Middleware(MiddlewareAuth)
		group.POST("/user/list", func(r *qn_http.Request) {
			r.Response.Write("list")
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		// Common Checks.
		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/api.v2"), "Not Found")
		// Auth Checks.
		t.Assert(client.PostContent("/api.v2/user/list"), "Forbidden")
		t.Assert(client.PostContent("/api.v2/user/list", "token=123456"), "list")
		// CORS Checks.
		resp, err := client.Post("/api.v2/user/list", "token=123456")
		t.Assert(err, nil)
		t.Assert(len(resp.Header["Access-Control-Allow-Headers"]), 1)
		t.Assert(resp.Header["Access-Control-Allow-Headers"][0], "Origin,Content-Type,Accept,User-Agent,Cookie,Authorization,X-Auth-Token,X-Requested-With")
		t.Assert(resp.Header["Access-Control-Allow-Methods"][0], "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE")
		t.Assert(resp.Header["Access-Control-Allow-Origin"][0], "*")
		t.Assert(resp.Header["Access-Control-Max-Age"][0], "3628800")
		resp.Close()
	})
}

func MiddlewareScope1(r *qn_http.Request) {
	r.Response.Write("a")
	r.Middleware.Next()
	r.Response.Write("b")
}

func MiddlewareScope2(r *qn_http.Request) {
	r.Response.Write("c")
	r.Middleware.Next()
	r.Response.Write("d")
}

func MiddlewareScope3(r *qn_http.Request) {
	r.Response.Write("e")
	r.Middleware.Next()
	r.Response.Write("f")
}

func Test_Middleware_Scope(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Middleware(MiddlewareScope1)
		group.ALL("/scope1", func(r *qn_http.Request) {
			r.Response.Write("1")
		})
		group.Group("/", func(group *qn_http.RouterGroup) {
			group.Middleware(MiddlewareScope2)
			group.ALL("/scope2", func(r *qn_http.Request) {
				r.Response.Write("2")
			})
		})
		group.Group("/", func(group *qn_http.RouterGroup) {
			group.Middleware(MiddlewareScope3)
			group.ALL("/scope3", func(r *qn_http.Request) {
				r.Response.Write("3")
			})
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "Not Found")
		t.Assert(client.GetContent("/scope1"), "a1b")
		t.Assert(client.GetContent("/scope2"), "ac2db")
		t.Assert(client.GetContent("/scope3"), "ae3fb")
	})
}

func Test_Middleware_Panic(t *testing.T) {
	p, _ := ports.PopRand()
	s := g.Server(p)
	i := 0
	s.Group("/", func(group *qn_http.RouterGroup) {
		group.Group("/", func(group *qn_http.RouterGroup) {
			group.Middleware(func(r *qn_http.Request) {
				i++
				panic("error")
				r.Middleware.Next()
			}, func(r *qn_http.Request) {
				i++
				r.Middleware.Next()
			})
			group.ALL("/", func(r *qn_http.Request) {
				r.Response.Write(i)
			})
		})
	})
	s.SetPort(p)
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		client := qn_http.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		t.Assert(client.GetContent("/"), "error")
	})
}
