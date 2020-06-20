// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_view

import "github.com/qnsoft/common/container/qn_map"

const (
	// Default group name for instance usage.
	DEFAULT_NAME = "default"
)

var (
	// Instances map.
	instances = qn_map.NewStrAnyMap(true)
)

// Instance returns an instance of View with default settings.
// The parameter <name> is the name for the instance.
func Instance(name ...string) *View {
	key := DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*View)
}
