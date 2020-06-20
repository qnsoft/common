// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_tcp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/net/qn_tcp"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_Pool_Basic1(t *testing.T) {
	p, _ := ports.PopRand()
	s := qn_tcp.NewServer(fmt.Sprintf(`:%d`, p), func(conn *qn_tcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				break
			}
			conn.SendPkg(data)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := []byte("9999")
		err = conn.SendPkg(data)
		t.Assert(err, nil)
		err = conn.SendPkgWithTimeout(data, time.Second)
		t.Assert(err, nil)
	})
}

func Test_Pool_Basic2(t *testing.T) {
	p, _ := ports.PopRand()
	s := qn_tcp.NewServer(fmt.Sprintf(`:%d`, p), func(conn *qn_tcp.Conn) {
		conn.Close()
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := []byte("9999")
		err = conn.SendPkg(data)
		t.Assert(err, nil)
		//err = conn.SendPkgWithTimeout(data, time.Second)
		//t.Assert(err, nil)

		_, err = conn.SendRecv(data, -1)
		t.AssertNE(err, nil)
	})
}
