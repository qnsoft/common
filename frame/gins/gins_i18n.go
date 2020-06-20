// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"github.com/qnsoft/common/i18n/qn_i18n"
)

// I18n returns an instance of qn_i18n.Manager.
// The parameter <name> is the name for the instance.
func I18n(name ...string) *qn_i18n.Manager {
	return qn_i18n.Instance(name...)
}
