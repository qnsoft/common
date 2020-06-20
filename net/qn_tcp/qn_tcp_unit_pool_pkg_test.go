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
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

func Test_Pool_Package_Basic(t *testing.T) {
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
	// SendPkg
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		for i := 0; i < 100; i++ {
			err := conn.SendPkg([]byte(qn_conv.String(i)))
			t.Assert(err, nil)
		}
		for i := 0; i < 100; i++ {
			err := conn.SendPkgWithTimeout([]byte(qn_conv.String(i)), time.Second)
			t.Assert(err, nil)
		}
	})
	// SendPkg with big data - failure.
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := make([]byte, 65536)
		err = conn.SendPkg(data)
		t.AssertNE(err, nil)
	})
	// SendRecvPkg
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		for i := 100; i < 200; i++ {
			data := []byte(qn_conv.String(i))
			result, err := conn.SendRecvPkg(data)
			t.Assert(err, nil)
			t.Assert(result, data)
		}
		for i := 100; i < 200; i++ {
			data := []byte(qn_conv.String(i))
			result, err := conn.SendRecvPkgWithTimeout(data, time.Second)
			t.Assert(err, nil)
			t.Assert(result, data)
		}
	})
	// SendRecvPkg with big data - failure.
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := make([]byte, 65536)
		result, err := conn.SendRecvPkg(data)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg with big data - success.
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := make([]byte, 65500)
		data[100] = byte(65)
		data[65400] = byte(85)
		result, err := conn.SendRecvPkg(data)
		t.Assert(err, nil)
		t.Assert(result, data)
	})
}

func Test_Pool_Package_Timeout(t *testing.T) {
	p, _ := ports.PopRand()
	s := qn_tcp.NewServer(fmt.Sprintf(`:%d`, p), func(conn *qn_tcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				break
			}
			time.Sleep(time.Second)
			qn_test.Assert(conn.SendPkg(data), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Millisecond*500)
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := []byte("10000")
		result, err := conn.SendRecvPkgWithTimeout(data, time.Second*2)
		t.Assert(err, nil)
		t.Assert(result, data)
	})
}

func Test_Pool_Package_Option(t *testing.T) {
	p, _ := ports.PopRand()
	s := qn_tcp.NewServer(fmt.Sprintf(`:%d`, p), func(conn *qn_tcp.Conn) {
		defer conn.Close()
		option := qn_tcp.PkgOption{HeaderSize: 1}
		for {
			data, err := conn.RecvPkg(option)
			if err != nil {
				break
			}
			qn_test.Assert(conn.SendPkg(data, option), nil)
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// SendRecvPkg with big data - failure.
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := make([]byte, 0xFF+1)
		result, err := conn.SendRecvPkg(data, qn_tcp.PkgOption{HeaderSize: 1})
		t.AssertNE(err, nil)
		t.Assert(result, nil)
	})
	// SendRecvPkg with big data - success.
	qn_test.C(t, func(t *qn_test.T) {
		conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", p))
		t.Assert(err, nil)
		defer conn.Close()
		data := make([]byte, 0xFF)
		data[100] = byte(65)
		data[200] = byte(85)
		result, err := conn.SendRecvPkg(data, qn_tcp.PkgOption{HeaderSize: 1})
		t.Assert(err, nil)
		t.Assert(result, data)
	})
}
