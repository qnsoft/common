// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package gmvc

import "github.com/qnsoft/common/database/gdb"

type (
	M     = Model      // M is alias for Model, just for short write purpose.
	Model = *gdb.Model // Model is alias for *gdb.Model.
)
