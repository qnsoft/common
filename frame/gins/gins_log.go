// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"fmt"

	"github.com/qnsoft/common/os/qn_log"
)

const (
	gFRAME_CORE_COMPONENT_NAME_LOGGER = "gf.core.component.logger"
	qn_logGER_NODE_NAME               = "logger"
)

// Log returns an instance of qn_log.Logger.
// The parameter <name> is the name for the instance.
func Log(name ...string) *qn_log.Logger {
	instanceName := qn_log.DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", gFRAME_CORE_COMPONENT_NAME_LOGGER, instanceName)
	return instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		logger := qn_log.Instance(instanceName)
		// To avoid file no found error while it's not necessary.
		if Config().Available() {
			var m map[string]interface{}
			// It firstly searches the configuration of the instance name.
			if m = Config().GetMap(fmt.Sprintf(`%s.%s`, qn_logGER_NODE_NAME, instanceName)); m == nil {
				// If the configuration for the instance does not exist,
				// it uses the default logging configuration.
				m = Config().GetMap(qn_logGER_NODE_NAME)
			}
			if m != nil {
				if err := logger.SetConfigWithMap(m); err != nil {
					panic(err)
				}
			}
		}
		return logger
	}).(*qn_log.Logger)
}
