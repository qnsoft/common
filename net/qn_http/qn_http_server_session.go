// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import "github.com/qnsoft/common/os/gsession"

// Session is actually a alias of gsession.Session,
// which is bound to a single request.
type Session = gsession.Session
