// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn

import "github.com/qnsoft/common/net/qn_http"

// SetServerGraceful enables/disables graceful reload feature of http Web Server.
// This feature is disabled in default.
// Deprecated, use configuration of qn_http.Server for controlling this feature.
func SetServerGraceful(enabled bool) {
	qn_http.SetGraceful(enabled)
}
