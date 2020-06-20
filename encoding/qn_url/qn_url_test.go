// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.
package qn_url_test

import (
	"net/url"
	"testing"

	"github.com/qnsoft/common/encoding/qn_url"
	"github.com/qnsoft/common/test/qn_test"
)

var urlStr string = `https://golang.org/x/crypto?go-get=1 +`
var urlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1+%2B`
var rawUrlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1%20%2B`

func TestEncodeAndDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_url.Encode(urlStr), urlEncode)

		res, err := qn_url.Decode(urlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestRowEncodeAndDecode(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		t.Assert(qn_url.RawEncode(urlStr), rawUrlEncode)

		res, err := qn_url.RawDecode(rawUrlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(res, urlStr)
	})
}

func TestBuildQuery(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		src := url.Values{
			"a": {"a2", "a1"},
			"b": {"b2", "b1"},
			"c": {"c1", "c2"},
		}
		expect := "a=a2&a=a1&b=b2&b=b1&c=c1&c=c2"
		t.Assert(qn_url.BuildQuery(src), expect)
	})
}

func TestParseURL(t *testing.T) {
	src := `http://username:password@hostname:9090/path?arg=value#anchor`
	expect := map[string]string{
		"scheme":   "http",
		"host":     "hostname",
		"port":     "9090",
		"user":     "username",
		"pass":     "password",
		"path":     "/path",
		"query":    "arg=value",
		"fragment": "anchor",
	}

	qn_test.C(t, func(t *qn_test.T) {
		component := 0
		for k, v := range []string{"all", "scheme", "host", "port", "user", "pass", "path", "query", "fragment"} {
			if v == "all" {
				component = -1
			} else {
				component = 1 << (uint(k - 1))
			}

			res, err := qn_url.ParseURL(src, component)
			if err != nil {
				t.Errorf("ParseURL failed. component:%v, err:%v", component, err)
				return
			}

			if v == "all" {
				t.Assert(res, expect)
			} else {
				t.Assert(res[v], expect[v])
			}

		}
	})
}
