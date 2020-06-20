// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_log_test

import (
	"sync"
	"testing"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_log"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
	"github.com/qnsoft/common/text/gstr"
)

func Test_Concurrent(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		c := 1000
		l := qn_log.New()
		s := "@1234567890#"
		f := "test.log"
		p := qn_file.TempDir(qn_time.TimestampNanoStr())
		t.Assert(l.SetPath(p), nil)
		defer qn_file.Remove(p)
		wg := sync.WaitGroup{}
		ch := make(chan struct{})
		for i := 0; i < c; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-ch
				l.File(f).Stdout(false).Print(s)
			}()
		}
		close(ch)
		wg.Wait()
		content := qn_file.GetContents(qn_file.Join(p, f))
		t.Assert(gstr.Count(content, s), c)
	})
}
