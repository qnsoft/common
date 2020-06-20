// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"fmt"

	"github.com/qnsoft/common/os/qn_view"
)

const (
	gFRAME_CORE_COMPONENT_NAME_VIEWER = "gf.core.component.viewer"
	qn_viewER_NODE_NAME               = "viewer"
)

// View returns an instance of View with default settings.
// The parameter <name> is the name for the instance.
func View(name ...string) *qn_view.View {
	instanceName := qn_view.DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", gFRAME_CORE_COMPONENT_NAME_VIEWER, instanceName)
	return instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		return getViewInstance(instanceName)
	}).(*qn_view.View)
}

func getViewInstance(name ...string) *qn_view.View {
	instanceName := qn_view.DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		instanceName = name[0]
	}
	view := qn_view.Instance(instanceName)
	// To avoid file no found error while it's not necessary.
	if Config().Available() {
		var m map[string]interface{}
		// It firstly searches the configuration of the instance name.
		if m = Config().GetMap(fmt.Sprintf(`%s.%s`, qn_viewER_NODE_NAME, instanceName)); m == nil {
			// If the configuration for the instance does not exist,
			// it uses the default view configuration.
			m = Config().GetMap(qn_viewER_NODE_NAME)
		}
		if m != nil {
			if err := view.SetConfigWithMap(m); err != nil {
				panic(err)
			}
		}
	}
	return view
}
