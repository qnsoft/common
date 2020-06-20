// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn

import (
	"github.com/qnsoft/common/os/qn_log"
)

// SetDebug disables/enables debug level for logging component globally.
func SetDebug(debug bool) {
	qn_log.SetDebug(debug)
}

// SetLogLevel sets the logging level globally.
func SetLogLevel(level int) {
	qn_log.SetLevel(level)
}

// GetLogLevel returns the global logging level.
func GetLogLevel() int {
	return qn_log.GetLevel()
}
