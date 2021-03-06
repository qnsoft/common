// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_proc

import (
	"errors"
	"fmt"

	"github.com/qnsoft/common/container/qn_map"
	"github.com/qnsoft/common/net/qn_tcp"
	"github.com/qnsoft/common/os/gfcache"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/util/qn_conv"
)

// MsgRequest is the request structure for process communication.
type MsgRequest struct {
	SendPid int    // Sender PID.
	RecvPid int    // Receiver PID.
	Group   string // Message group name.
	Data    []byte // Request data.
}

// Msqn_response is the response structure for process communication.
type Msqn_response struct {
	Code    int    // 1: OK; Other: Error.
	Message string // Response message.
	Data    []byte // Response data.
}

const (
	gPROC_COMM_DEFAULT_GRUOP_NAME = ""    // Default group name.
	gPROC_DEFAULT_TCP_PORT        = 10000 // Starting port number for receiver listening.
	gPROC_MSG_QUEUE_MAX_LENGTH    = 10000 // Max size for each message queue of the group.
)

var (
	// commReceiveQueues is the group name to queue map for storing received data.
	// The value of the map is type of *gqueue.Queue.
	commReceiveQueues = qn_map.NewStrAnyMap(true)

	// commPidFolderPath specifies the folder path storing pid to port mapping files.
	commPidFolderPath = qn_file.TempDir("gproc")
)

func init() {
	// Automatically create the storage folder.
	if !qn_file.Exists(commPidFolderPath) {
		err := qn_file.Mkdir(commPidFolderPath)
		if err != nil {
			panic(fmt.Errorf(`create gproc folder failed: %v`, err))
		}
	}
}

// getConnByPid creates and returns a TCP connection for specified pid.
func getConnByPid(pid int) (*qn_tcp.PoolConn, error) {
	port := getPortByPid(pid)
	if port > 0 {
		if conn, err := qn_tcp.NewPoolConn(fmt.Sprintf("127.0.0.1:%d", port)); err == nil {
			return conn, nil
		} else {
			return nil, err
		}
	}
	return nil, errors.New(fmt.Sprintf("could not find port for pid: %d", pid))
}

// getPortByPid returns the listening port for specified pid.
// It returns 0 if no port found for the specified pid.
func getPortByPid(pid int) int {
	path := getCommFilePath(pid)
	content := gfcache.GetContents(path)
	return qn_conv.Int(content)
}

// getCommFilePath returns the pid to port mapping file path for given pid.
func getCommFilePath(pid int) string {
	return qn_file.Join(commPidFolderPath, qn_conv.String(pid))
}
