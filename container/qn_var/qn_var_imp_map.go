// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_var

import gconv "github.com/qnsoft/common/util/qn_conv"

// Map converts and returns <v> as map[string]interface{}.
func (v *VarImp) Map(tags ...string) map[string]interface{} {
	return gconv.Map(v.Val(), tags...)
}

// MapStrStr converts and returns <v> as map[string]string.
func (v *VarImp) MapStrStr(tags ...string) map[string]string {
	return gconv.MapStrStr(v.Val(), tags...)
}

// MapStrVar converts and returns <v> as map[string]Var.
func (v *VarImp) MapStrVar(tags ...string) map[string]Var {
	m := v.Map(tags...)
	if len(m) > 0 {
		vMap := make(map[string]Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// MapDeep converts and returns <v> as map[string]interface{} recursively.
func (v *VarImp) MapDeep(tags ...string) map[string]interface{} {
	return gconv.MapDeep(v.Val(), tags...)
}

// MapDeep converts and returns <v> as map[string]string recursively.
func (v *VarImp) MapStrStrDeep(tags ...string) map[string]string {
	return gconv.MapStrStrDeep(v.Val(), tags...)
}

// MapStrVarDeep converts and returns <v> as map[string]*VarImp recursively.
func (v *VarImp) MapStrVarDeep(tags ...string) map[string]Var {
	m := v.MapDeep(tags...)
	if len(m) > 0 {
		vMap := make(map[string]Var, len(m))
		for k, v := range m {
			vMap[k] = New(v)
		}
		return vMap
	}
	return nil
}

// Maps converts and returns <v> as map[string]string.
// See gconv.Maps.
func (v *VarImp) Maps(tags ...string) []map[string]interface{} {
	return gconv.Maps(v.Val(), tags...)
}

// MapToMap converts any map type variable <params> to another map type variable <pointer>.
// See gconv.MapToMap.
func (v *VarImp) MapToMap(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMap(v.Val(), pointer, mapping...)
}

// MapToMapDeep converts any map type variable <params> to another map type variable
// <pointer> recursively.
// See gconv.MapToMapDeep.
func (v *VarImp) MapToMapDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMapDeep(v.Val(), pointer, mapping...)
}

// MapToMaps converts any map type variable <params> to another map type variable <pointer>.
// See gconv.MapToMaps.
func (v *VarImp) MapToMaps(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMaps(v.Val(), pointer, mapping...)
}

// MapToMapsDeep converts any map type variable <params> to another map type variable
// <pointer> recursively.
// See gconv.MapToMapsDeep.
func (v *VarImp) MapToMapsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.MapToMapsDeep(v.Val(), pointer, mapping...)
}
