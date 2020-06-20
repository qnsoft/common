// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_udp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qnsoft/common/net/gudp"
	"github.com/qnsoft/common/os/glog"
	"github.com/qnsoft/common/test/qn_test"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

func Test_Basic(t *testing.T) {
	p, _ := ports.PopRand()
	s := gudp.NewServer(fmt.Sprintf("127.0.0.1:%d", p), func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					glog.Error(err)
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	// gudp.Conn.Send
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(fmt.Sprintf("127.0.0.1:%d", p))
			t.Assert(err, nil)
			t.Assert(conn.Send([]byte(qn_conv.String(i))), nil)
			conn.Close()
		}
	})
	// gudp.Conn.SendRecv
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewConn(fmt.Sprintf("127.0.0.1:%d", p))
			t.Assert(err, nil)
			result, err := conn.SendRecv([]byte(qn_conv.String(i)), -1)
			t.Assert(err, nil)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})
	// gudp.Send
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			err := gudp.Send(fmt.Sprintf("127.0.0.1:%d", p), []byte(qn_conv.String(i)))
			t.Assert(err, nil)
		}
	})
	// gudp.SendRecv
	qn_test.C(t, func(t *qn_test.T) {
		for i := 0; i < 100; i++ {
			result, err := gudp.SendRecv(fmt.Sprintf("127.0.0.1:%d", p), []byte(qn_conv.String(i)), -1)
			t.Assert(err, nil)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
		}
	})
}

// If the read buffer size is less than the sent package size,
// the rest data would be dropped.
func Test_Buffer(t *testing.T) {
	p, _ := ports.PopRand()
	s := gudp.NewServer(fmt.Sprintf("127.0.0.1:%d", p), func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(1)
			if len(data) > 0 {
				if err := conn.Send(data); err != nil {
					glog.Error(err)
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	qn_test.C(t, func(t *qn_test.T) {
		result, err := gudp.SendRecv(fmt.Sprintf("127.0.0.1:%d", p), []byte("123"), -1)
		t.Assert(err, nil)
		t.Assert(string(result), "1")
	})
	qn_test.C(t, func(t *qn_test.T) {
		result, err := gudp.SendRecv(fmt.Sprintf("127.0.0.1:%d", p), []byte("456"), -1)
		t.Assert(err, nil)
		t.Assert(string(result), "4")
	})
}
