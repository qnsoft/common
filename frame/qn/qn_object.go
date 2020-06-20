// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn

import (
	"github.com/qnsoft/common/database/gdb"
	"github.com/qnsoft/common/database/gredis"
	"github.com/qnsoft/common/frame/qn_ins"
	"github.com/qnsoft/common/i18n/qn_i18n"
	"github.com/qnsoft/common/net/qn_http"
	"github.com/qnsoft/common/net/qn_tcp"
	"github.com/qnsoft/common/net/qn_udp"
	"github.com/qnsoft/common/os/qn_cfg"
	"github.com/qnsoft/common/os/qn_log"
	"github.com/qnsoft/common/os/qn_res"
	"github.com/qnsoft/common/os/qn_view"
)

// Client is a convenience function, that creates and returns a new HTTP client.
func Client() *qn_http.Client {
	return qn_http.NewClient()
}

// Server returns an instance of http server with specified name.
func Server(name ...interface{}) *qn_http.Server {
	return qn_ins.Server(name...)
}

// TCPServer returns an instance of tcp server with specified name.
func TCPServer(name ...interface{}) *qn_tcp.Server {
	return qn_tcp.GetServer(name...)
}

// UDPServer returns an instance of udp server with specified name.
func UDPServer(name ...interface{}) *qn_udp.Server {
	return qn_udp.GetServer(name...)
}

// View returns an instance of template engine object with specified name.
func View(name ...string) *qn_view.View {
	return qn_ins.View(name...)
}

// Config returns an instance of config object with specified name.
func Config(name ...string) *qn_cfg.Config {
	return qn_ins.Config(name...)
}

// Cfg is alias of Config.
// See Config.
func Cfg(name ...string) *qn_cfg.Config {
	return Config(name...)
}

// Resource returns an instance of Resource.
// The parameter <name> is the name for the instance.
func Resource(name ...string) *qn_res.Resource {
	return qn_ins.Resource(name...)
}

// I18n returns an instance of qn_i18n.Manager.
// The parameter <name> is the name for the instance.
func I18n(name ...string) *qn_i18n.Manager {
	return qn_ins.I18n(name...)
}

// Res is alias of Resource.
// See Resource.
func Res(name ...string) *qn_res.Resource {
	return Resource(name...)
}

// Log returns an instance of qn_log.Logger.
// The parameter <name> is the name for the instance.
func Log(name ...string) *qn_log.Logger {
	return qn_ins.Log(name...)
}

// Database returns an instance of database ORM object with specified configuration group name.
func Database(name ...string) gdb.DB {
	return qn_ins.Database(name...)
}

// DB is alias of Database.
// See Database.
func DB(name ...string) gdb.DB {
	return qn_ins.Database(name...)
}

// Table creates and returns a model from specified database or default database configuration.
// The optional parameter <db> specifies the configuration group name of the database,
// which is "default" in default.
func Table(tables string, db ...string) *gdb.Model {
	return DB(db...).Table(tables)
}

// Redis returns an instance of redis client with specified configuration group name.
func Redis(name ...string) *gredis.Redis {
	return qn_ins.Redis(name...)
}
