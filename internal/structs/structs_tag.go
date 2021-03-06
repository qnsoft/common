// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package structs

import (
	"reflect"

	"github.com/gqcn/structs"
)

// TagFields retrieves struct tags as []*Field from <pointer>, and returns it.
//
// The parameter <pointer> should be type of struct/*struct.
//
// The parameter <recursive> specifies whether retrieving the struct field recursively.
//
// Note that it only retrieves the exported attributes with first letter up-case from struct.
func TagFields(pointer interface{}, priority []string, recursive bool) []*Field {
	return doTagFields(pointer, priority, recursive, map[string]struct{}{})
}

// doTagFields retrieves the tag and corresponding attribute name from <pointer>. It also filters repeated
// tag internally.
// The parameter <pointer> should be type of struct/*struct.
func doTagFields(pointer interface{}, priority []string, recursive bool, taqn_map map[string]struct{}) []*Field {
	var fields []*structs.Field
	if v, ok := pointer.(reflect.Value); ok {
		fields = structs.Fields(v.Interface())
	} else {
		var (
			rv   = reflect.ValueOf(pointer)
			kind = rv.Kind()
		)
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		// If pointer is type of **struct and nil, then automatically create a temporary struct,
		// which is used for structs.Fields.
		if kind == reflect.Ptr && (!rv.IsValid() || rv.IsNil()) {
			fields = structs.Fields(reflect.New(rv.Type().Elem()).Elem().Interface())
		} else {
			fields = structs.Fields(pointer)
		}
	}
	var (
		tag  = ""
		name = ""
	)
	tagFields := make([]*Field, 0)
	for _, field := range fields {
		name = field.Name()
		// Only retrieve exported attributes.
		if name[0] < byte('A') || name[0] > byte('Z') {
			continue
		}
		tag = ""
		for _, p := range priority {
			tag = field.Tag(p)
			if tag != "" {
				break
			}
		}
		if tag != "" {
			// Filter repeated tag.
			if _, ok := taqn_map[tag]; ok {
				continue
			}
			tagFields = append(tagFields, &Field{
				Field: field,
				Tag:   tag,
			})
		}
		if recursive {
			var (
				rv   = reflect.ValueOf(field.Value())
				kind = rv.Kind()
			)
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			if kind == reflect.Struct {
				tagFields = append(tagFields, doTagFields(rv, priority, recursive, taqn_map)...)
			}
		}
	}
	return tagFields
}

// Taqn_mapName retrieves struct tags as map[tag]attribute from <pointer>, and returns it.
//
// The parameter <pointer> should be type of struct/*struct.
//
// The parameter <recursive> specifies whether retrieving the struct field recursively.
//
// Note that it only retrieves the exported attributes with first letter up-case from struct.
func Taqn_mapName(pointer interface{}, priority []string, recursive bool) map[string]string {
	fields := TagFields(pointer, priority, recursive)
	taqn_map := make(map[string]string, len(fields))
	for _, v := range fields {
		taqn_map[v.Tag] = v.Name()
	}
	return taqn_map
}

// Taqn_mapField retrieves struct tags as map[tag]*Field from <pointer>, and returns it.
//
// The parameter <pointer> should be type of struct/*struct.
//
// The parameter <recursive> specifies whether retrieving the struct field recursively.
//
// Note that it only retrieves the exported attributes with first letter up-case from struct.
func Taqn_mapField(pointer interface{}, priority []string, recursive bool) map[string]*Field {
	fields := TagFields(pointer, priority, recursive)
	taqn_map := make(map[string]*Field, len(fields))
	for _, v := range fields {
		taqn_map[v.Tag] = v
	}
	return taqn_map
}
