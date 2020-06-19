// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// package qn_log implements powerful and easy-to-use levelled logging functionality.
package qn_log

import (
	"github.com/qnsoft/common/internal/cmdenv"
	"github.com/qnsoft/common/os/grpool"
)

var (
	// Default logger object, for package method usage.
	logger = New()

	// Goroutine pool for async logging output.
	// It uses only one asynchronize worker to ensure log sequence.
	asyncPool = grpool.New(1)

	// defaultDebug enables debug level or not in default,
	// which can be configured using command option or system environment.
	defaultDebug = true
)

func init() {
	defaultDebug = cmdenv.Get("qn.qn_log.debug", true).Bool()
	SetDebug(defaultDebug)
}

// Default returns the default logger.
func DefaultLogger() *Logger {
	return logger
}

// SetDefaultLogger sets the default logger for package qn_log.
// Note that there might be concurrent safety issue if calls this function
// in different goroutines.
func SetDefaultLogger(l *Logger) {
	logger = l
}
