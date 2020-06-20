// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"fmt"

	"github.com/qnsoft/common/net/qn_http"
)

const (
	gFRAME_CORE_COMPONENT_NAME_SERVER = "gf.core.component.server"
)

// Server returns an instance of http server with specified name.
func Server(name ...interface{}) *qn_http.Server {
	instanceKey := fmt.Sprintf("%s.%v", gFRAME_CORE_COMPONENT_NAME_SERVER, name)
	return instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		s := qn_http.GetServer(name...)
		// To avoid file no found error while it's not necessary.
		if Config().Available() {
			var m map[string]interface{}
			// It firstly searches the configuration of the instance name.
			if m = Config().GetMap(fmt.Sprintf(`server.%s`, s.GetName())); m == nil {
				// If the configuration for the instance does not exist,
				// it uses the default server configuration.
				m = Config().GetMap("server")
			}
			if m != nil {
				if err := s.SetConfigWithMap(m); err != nil {
					panic(err)
				}
			}
			// As it might use template feature,
			// it initialize the view instance as well.
			_ = getViewInstance()
		}
		return s
	}).(*qn_http.Server)
}
