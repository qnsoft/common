// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// 静态文件搜索优先级: Resource > ServerPaths > ServerRoot > SearchPath

package qn_http

import (
	"fmt"
	"strings"

	"github.com/qnsoft/common/os/gres"
	"github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/os/qn_file"
)

// staticPathItem is the item struct for static path configuration.
type staticPathItem struct {
	prefix string // The router URI.
	path   string // The static path.
}

// SetIndexFiles sets the index files for server.
func (s *Server) SetIndexFiles(indexFiles []string) {
	s.config.IndexFiles = indexFiles
}

// GetIndexFiles retrieves and returns the index files from server.
func (s *Server) GetIndexFiles() []string {
	return s.config.IndexFiles
}

// SetIndexFolder enables/disables listing the sub-files if requesting a directory.
func (s *Server) SetIndexFolder(enabled bool) {
	s.config.IndexFolder = enabled
}

// SetFileServerEnabled enables/disables the static file service.
// It's the main switch for the static file service. When static file service configuration
// functions like SetServerRoot, AddSearchPath and AddStaticPath are called, this configuration
// is automatically enabled.
func (s *Server) SetFileServerEnabled(enabled bool) {
	s.config.FileServerEnabled = enabled
}

// SetServerRoot sets the document root for static service.
func (s *Server) SetServerRoot(root string) {
	realPath := root
	if !gres.Contains(realPath) {
		if p, err := qn_file.Search(root); err != nil {
			s.Logger().Fatal(fmt.Sprintf(`[ghttp] SetServerRoot failed: %v`, err))
		} else {
			realPath = p
		}
	}
	s.Logger().Debug("[ghttp] SetServerRoot path:", realPath)
	s.config.SearchPaths = []string{strings.TrimRight(realPath, qn_file.Separator)}
	s.config.FileServerEnabled = true
}

// AddSearchPath add searching directory path for static file service.
func (s *Server) AddSearchPath(path string) {
	realPath := path
	if !gres.Contains(realPath) {
		if p, err := qn_file.Search(path); err != nil {
			s.Logger().Fatal(fmt.Sprintf(`[ghttp] AddSearchPath failed: %v`, err))
		} else {
			realPath = p
		}
	}
	s.config.SearchPaths = append(s.config.SearchPaths, realPath)
	s.config.FileServerEnabled = true
}

// AddStaticPath sets the uri to static directory path mapping for static file service.
func (s *Server) AddStaticPath(prefix string, path string) {
	realPath := path
	if !gres.Contains(realPath) {
		if p, err := qn_file.Search(path); err != nil {
			s.Logger().Fatal(fmt.Sprintf(`[ghttp] AddStaticPath failed: %v`, err))
		} else {
			realPath = p
		}
	}
	addItem := staticPathItem{
		prefix: prefix,
		path:   realPath,
	}
	if len(s.config.StaticPaths) > 0 {
		s.config.StaticPaths = append(s.config.StaticPaths, addItem)
		// Sort the array by length of prefix from short to long.
		array := qn_array.NewSortedArray(func(v1, v2 interface{}) int {
			s1 := qn_conv.String(v1)
			s2 := qn_conv.String(v2)
			r := len(s2) - len(s1)
			if r == 0 {
				r = strings.Compare(s1, s2)
			}
			return r
		})
		for _, v := range s.config.StaticPaths {
			array.Add(v.prefix)
		}
		// Add the items to paths by previous sorted slice.
		paths := make([]staticPathItem, 0)
		for _, v := range array.Slice() {
			for _, item := range s.config.StaticPaths {
				if strings.EqualFold(qn_conv.String(v), item.prefix) {
					paths = append(paths, item)
					break
				}
			}
		}
		s.config.StaticPaths = paths
	} else {
		s.config.StaticPaths = []staticPathItem{addItem}
	}
	s.config.FileServerEnabled = true
}
