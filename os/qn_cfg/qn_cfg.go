// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// package qn_cfg provides reading, caching and managing for configuration.
package qn_cfg

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gogf/gf/text/gstr"
	"github.com/qnsoft/common/os/qn_res"

	"github.com/qnsoft/common/container/gmap"
	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/encoding/qn_json"
	"github.com/qnsoft/common/internal/cmdenv"
	"github.com/qnsoft/common/os/gfsnotify"
	"github.com/qnsoft/common/os/gspath"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_log"
)

const (
	// DEFAULT_CONFIG_FILE is the default configuration file name.
	DEFAULT_CONFIG_FILE = "config.toml"
)

// Configuration struct.
type Config struct {
	name  string             // Default configuration file name.
	paths *qn_array.StrArray // Searching path array.
	jsons *gmap.StrAnyMap    // The pared JSON objects for configuration files.
	vc    bool               // Whether do violence check in value index searching. It affects the performance when set true(false in default).
}

var (
	resourceTryFiles = []string{"", "/", "config/", "config", "/config", "/config/"}
)

// New returns a new configuration management object.
// The parameter <file> specifies the default configuration file name for reading.
func New(file ...string) *Config {
	name := DEFAULT_CONFIG_FILE
	if len(file) > 0 {
		name = file[0]
	}
	c := &Config{
		name:  name,
		paths: qn_array.NewStrArray(true),
		jsons: gmap.NewStrAnyMap(true),
	}
	// Customized dir path from env/cmd.
	if envPath := cmdenv.Get("gf.qn_cfg.path").String(); envPath != "" {
		if qn_file.Exists(envPath) {
			_ = c.SetPath(envPath)
		} else {
			if errorPrint() {
				qn_log.Errorf("Configuration directory path does not exist: %s", envPath)
			}
		}
	} else {
		// Dir path of working dir.
		_ = c.SetPath(qn_file.Pwd())
		// Dir path of binary.
		if selfPath := qn_file.SelfDir(); selfPath != "" && qn_file.Exists(selfPath) {
			_ = c.AddPath(selfPath)
		}
		// Dir path of main package.
		if mainPath := qn_file.MainPkgPath(); mainPath != "" && qn_file.Exists(mainPath) {
			_ = c.AddPath(mainPath)
		}
	}
	return c
}

// filePath returns the absolute configuration file path for the given filename by <file>.
func (c *Config) filePath(file ...string) (path string) {
	name := c.name
	if len(file) > 0 {
		name = file[0]
	}
	path = c.FilePath(name)
	if path == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[qn_cfg] cannot find config file \"%s\" in following paths:", name))
			c.paths.RLockFunc(func(array []string) {
				index := 1
				for _, v := range array {
					v = gstr.TrimRight(v, `\/`)
					buffer.WriteString(fmt.Sprintf("\n%d. %s", index, v))
					index++
					buffer.WriteString(fmt.Sprintf("\n%d. %s", index, v+qn_file.Separator+"config"))
					index++
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf("[qn_cfg] cannot find config file \"%s\" with no path set/add", name))
		}
		if errorPrint() {
			qn_log.Error(buffer.String())
		}
	}
	return path
}

// SetPath sets the configuration directory path for file search.
// The parameter <path> can be absolute or relative path,
// but absolute path is strongly recommended.
func (c *Config) SetPath(path string) error {
	isDir := false
	realPath := ""
	if file := qn_res.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = qn_file.RealPath(path)
		if realPath == "" {
			// Relative path.
			c.paths.RLockFunc(func(array []string) {
				for _, v := range array {
					if path, _ := gspath.Search(v, path); path != "" {
						realPath = path
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = qn_file.IsDir(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[qn_cfg] SetPath failed: cannot find directory \"%s\" in following paths:", path))
			c.paths.RLockFunc(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`[qn_cfg] SetPath failed: path "%s" does not exist`, path))
		}
		err := errors.New(buffer.String())
		if errorPrint() {
			qn_log.Error(err)
		}
		return err
	}
	// Should be a directory.
	if !isDir {
		err := fmt.Errorf(`[qn_cfg] SetPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			qn_log.Error(err)
		}
		return err
	}
	// Repeated path check.
	if c.paths.Search(realPath) != -1 {
		return nil
	}
	c.jsons.Clear()
	c.paths.Clear()
	c.paths.Append(realPath)
	return nil
}

// SetViolenceCheck sets whether to perform hierarchical conflict checking.
// This feature needs to be enabled when there is a level symbol in the key name.
// It is off in default.
//
// Note that, turning on this feature is quite expensive, and it is not recommended
// to allow separators in the key names. It is best to avoid this on the application side.
func (c *Config) SetViolenceCheck(check bool) {
	c.vc = check
	c.Clear()
}

// AddPath adds a absolute or relative path to the search paths.
func (c *Config) AddPath(path string) error {
	var (
		isDir    = false
		realPath = ""
	)
	// It firstly checks the resource manager,
	// and then checks the filesystem for the path.
	if file := qn_res.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = qn_file.RealPath(path)
		if realPath == "" {
			// Relative path.
			c.paths.RLockFunc(func(array []string) {
				for _, v := range array {
					if path, _ := gspath.Search(v, path); path != "" {
						realPath = path
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = qn_file.IsDir(realPath)
		}
	}
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[qn_cfg] AddPath failed: cannot find directory \"%s\" in following paths:", path))
			c.paths.RLockFunc(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`[qn_cfg] AddPath failed: path "%s" does not exist`, path))
		}
		err := errors.New(buffer.String())
		if errorPrint() {
			qn_log.Error(err)
		}
		return err
	}
	if !isDir {
		err := fmt.Errorf(`[qn_cfg] AddPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			qn_log.Error(err)
		}
		return err
	}
	// Repeated path check.
	if c.paths.Search(realPath) != -1 {
		return nil
	}
	c.paths.Append(realPath)
	//qn_log.Debug("[qn_cfg] AddPath:", realPath)
	return nil
}

// GetFilePath returns the absolute path of the specified configuration file.
// If <file> is not passed, it returns the configuration file path of the default name.
// If the specified configuration file does not exist,
// an empty string is returned.
func (c *Config) FilePath(file ...string) (path string) {
	name := c.name
	if len(file) > 0 {
		name = file[0]
	}
	// Searching resource manager.
	if !qn_res.IsEmpty() {
		for _, v := range resourceTryFiles {
			if file := qn_res.Get(v + name); file != nil {
				path = file.Name()
				return
			}
		}
		c.paths.RLockFunc(func(array []string) {
			for _, prefix := range array {
				for _, v := range resourceTryFiles {
					if file := qn_res.Get(prefix + v + name); file != nil {
						path = file.Name()
						return
					}
				}
			}
		})
	}
	// Searching the file system.
	c.paths.RLockFunc(func(array []string) {
		for _, prefix := range array {
			prefix = gstr.TrimRight(prefix, `\/`)
			if path, _ = gspath.Search(prefix, name); path != "" {
				return
			}
			if path, _ = gspath.Search(prefix+qn_file.Separator+"config", name); path != "" {
				return
			}
		}
	})
	return
}

// SetFileName sets the default configuration file name.
func (c *Config) SetFileName(name string) *Config {
	c.name = name
	return c
}

// GetFileName returns the default configuration file name.
func (c *Config) GetFileName() string {
	return c.name
}

// Available checks and returns whether configuration of given <file> is available.
func (c *Config) Available(file ...string) bool {
	var name string
	if len(file) > 0 && file[0] != "" {
		name = file[0]
	} else {
		name = c.name
	}
	if c.FilePath(name) != "" {
		return true
	}
	if GetContent(name) != "" {
		return true
	}
	return false
}

// getJson returns a *qn_json.Json object for the specified <file> content.
// It would print error if file reading fails. It return nil if any error occurs.
func (c *Config) getJson(file ...string) *qn_json.Json {
	var name string
	if len(file) > 0 && file[0] != "" {
		name = file[0]
	} else {
		name = c.name
	}
	r := c.jsons.GetOrSetFuncLock(name, func() interface{} {
		content := ""
		filePath := ""
		if content = GetContent(name); content == "" {
			filePath = c.filePath(name)
			if filePath == "" {
				return nil
			}
			if file := qn_res.Get(filePath); file != nil {
				content = string(file.Content())
			} else {
				content = qn_file.GetContents(filePath)
			}
		}
		if j, err := qn_json.LoadContent(content, true); err == nil {
			j.SetViolenceCheck(c.vc)
			// Add monitor for this configuration file,
			// any changes of this file will refresh its cache in Config object.
			if filePath != "" && !qn_res.Contains(filePath) {
				_, err = gfsnotify.Add(filePath, func(event *gfsnotify.Event) {
					c.jsons.Remove(name)
				})
				if err != nil && errorPrint() {
					qn_log.Error(err)
				}
			}
			return j
		} else {
			if errorPrint() {
				if filePath != "" {
					qn_log.Criticalf(`[qn_cfg] Load config file "%s" failed: %s`, filePath, err.Error())
				} else {
					qn_log.Criticalf(`[qn_cfg] Load configuration failed: %s`, err.Error())
				}
			}
		}
		return nil
	})
	if r != nil {
		return r.(*qn_json.Json)
	}
	return nil
}
