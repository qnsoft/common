// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import (
	"time"

	"github.com/qnsoft/common/os/gsession"
)

// SetSessionMaxAge sets the SessionMaxAge for server.
func (s *Server) SetSessionMaxAge(ttl time.Duration) {
	s.config.SessionMaxAge = ttl
}

// SetSessionIdName sets the SessionIdName for server.
func (s *Server) SetSessionIdName(name string) {
	s.config.SessionIdName = name
}

// SetSessionStorage sets the SessionStorage for server.
func (s *Server) SetSessionStorage(storage gsession.Storage) {
	s.config.SessionStorage = storage
}

// GetSessionMaxAge returns the SessionMaxAge of server.
func (s *Server) GetSessionMaxAge() time.Duration {
	return s.config.SessionMaxAge
}

// GetSessionIdName returns the SessionIdName of server.
func (s *Server) GetSessionIdName() string {
	return s.config.SessionIdName
}
