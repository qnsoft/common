// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"github.com/qnsoft/common/os/qn_cfg"
)

// Config returns an instance of View with default settings.
// The parameter <name> is the name for the instance.
func Config(name ...string) *qn_cfg.Config {
	return qn_cfg.Instance(name...)
}
