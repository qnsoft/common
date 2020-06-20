// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.
package qn_html_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_html"
	"github.com/qnsoft/common/test/qn_test"
)

func TestStripTags(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
		dst := `Test paragraph.  Other text`
		t.Assert(qn_html.StripTags(src), dst)
	})
}

func TestEntities(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(qn_html.Entities(src), dst)
		t.Assert(qn_html.EntitiesDecode(dst), src)
	})
}

func TestSpecialChars(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		src := `A 'quote' "is" <b>bold</b>`
		dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
		t.Assert(qn_html.SpecialChars(src), dst)
		t.Assert(qn_html.SpecialCharsDecode(dst), src)
	})
}
