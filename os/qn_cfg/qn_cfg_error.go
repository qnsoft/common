// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cfg

import (
	"github.com/qnsoft/common/internal/cmdenv"
)

const (
	// qn_error_PRINT_KEY is used to specify the key controlling error printing to stdout.
	// This error is designed not to be returned by functions.
	qn_error_PRINT_KEY = "gf.qn_cfg.errorprint"
)

// errorPrint checks whether printing error to stdout.
func errorPrint() bool {
	return cmdenv.Get(qn_error_PRINT_KEY, true).Bool()
}
