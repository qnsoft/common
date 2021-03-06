// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_ins

import (
	"fmt"

	"github.com/qnsoft/common/internal/intlog"
	qn_util "github.com/qnsoft/common/util/qn_util"

	"github.com/qnsoft/common/database/gdb"
	"github.com/qnsoft/common/text/qn_regex"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

const (
	gFRAME_CORE_COMPONENT_NAME_DATABASE = "gf.core.component.database"
)

// Database returns an instance of database ORM object
// with specified configuration group name.
func Database(name ...string) gdb.DB {
	config := Config()
	group := gdb.DEFAULT_GROUP_NAME
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", gFRAME_CORE_COMPONENT_NAME_DATABASE, group)
	db := instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		// Configuration already exists.
		if gdb.GetConfig(group) != nil {
			db, err := gdb.Instance(group)
			if err != nil {
				panic(err)
			}
			return db
		}
		m := config.GetMap("database")
		if m == nil {
			panic(`database init failed: "database" node not found, is config file or configuration missing?`)
		}
		// Parse <m> as map-slice and adds it to gdb's global configurations.
		for group, groupConfig := range m {
			cg := gdb.ConfigGroup{}
			switch value := groupConfig.(type) {
			case []interface{}:
				for _, v := range value {
					if node := parseDBConfigNode(v); node != nil {
						cg = append(cg, *node)
					}
				}
			case map[string]interface{}:
				if node := parseDBConfigNode(value); node != nil {
					cg = append(cg, *node)
				}
			}
			if len(cg) > 0 {
				intlog.Printf("%s, %#v", group, cg)
				gdb.SetConfigGroup(group, cg)
			}
		}
		// Parse <m> as a single node configuration,
		// which is the default group configuration.
		if node := parseDBConfigNode(m); node != nil {
			cg := gdb.ConfigGroup{}
			if node.LinkInfo != "" || node.Host != "" {
				cg = append(cg, *node)
			}
			if len(cg) > 0 {
				intlog.Printf("%s, %#v", gdb.DEFAULT_GROUP_NAME, cg)
				gdb.SetConfigGroup(gdb.DEFAULT_GROUP_NAME, cg)
			}
		}

		if db, err := gdb.New(name...); err == nil {
			// Initialize logger for ORM.
			m := config.GetMap(fmt.Sprintf("database.%s", qn_logGER_NODE_NAME))
			if m == nil {
				m = config.GetMap(qn_logGER_NODE_NAME)
			}
			if m != nil {
				if err := db.GetLogger().SetConfigWithMap(m); err != nil {
					panic(err)
				}
			}
			return db
		} else {
			panic(err)
		}
		return nil
	})
	if db != nil {
		return db.(gdb.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *gdb.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	node := &gdb.ConfigNode{}
	err := qn_conv.Struct(nodeMap, node)
	if err != nil {
		panic(err)
	}
	if _, v := qn_util.MapPossibleItemByKey(nodeMap, "link"); v != nil {
		node.LinkInfo = qn_conv.String(v)
	}
	// Parse link syntax.
	if node.LinkInfo != "" && node.Type == "" {
		match, _ := qn_regex.MatchString(`([a-z]+):(.+)`, node.LinkInfo)
		if len(match) == 3 {
			node.Type = qn.str.Trim(match[1])
			node.LinkInfo = qn.str.Trim(match[2])
		}
	}
	return node
}
