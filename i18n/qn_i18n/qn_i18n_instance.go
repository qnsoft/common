// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_i18n

import "github.com/qnsoft/common/container/qn_map"

const (
	// Default group name for instance usage.
	DEFAULT_NAME = "default"
)

var (
	// instances is the instances map for management
	// for multiple i18n instance by name.
	instances = qn_map.NewStrAnyMap(true)
)

// Instance returns an instance of Resource.
// The parameter <name> is the name for the instance.
func Instance(name ...string) *Manager {
	key := DEFAULT_NAME
	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*Manager)
}
