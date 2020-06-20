// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_session

import (
	"errors"
	"time"

	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/os/qn_time"

	"github.com/qnsoft/common/container/qn_map"
	"github.com/qnsoft/common/container/qn_var"
	qn_conv "github.com/qnsoft/common/util/qn_conv"
)

// Session struct for storing single session data,
// which is bound to a single request.
type Session struct {
	id      string            // Session id.
	data    *qn_map.StrAnyMap // Session data.
	dirty   bool              // Used to mark session is modified.
	start   bool              // Used to mark session is started.
	manager *Manager          // Parent manager.

	// idFunc is a callback function used for creating custom session id.
	// This is called if session id is empty ever when session starts.
	idFunc func(ttl time.Duration) (id string)
}

// init does the lazy initialization for session.
// It here initializes real session if necessary.
func (s *Session) init() {
	if s.start {
		return
	}
	if s.id != "" {
		var err error
		// Retrieve memory session data from manager.
		if r := s.manager.sessionData.Get(s.id); r != nil {
			s.data = r.(*qn_map.StrAnyMap)
			intlog.Print("session init data:", s.data)
		}
		// Retrieve stored session data from storage.
		if s.manager.storage != nil {
			if s.data, err = s.manager.storage.GetSession(s.id, s.manager.ttl, s.data); err != nil {
				intlog.Errorf("session restoring failed for id '%s': %v", s.id, err)
			}
		}
		// If it's an invalid or expired session id,
		// it should create a new session id.
		if s.data == nil {
			s.id = ""
		}
	}
	// Use custom session id creating function.
	if s.id == "" && s.idFunc != nil {
		s.id = s.idFunc(s.manager.ttl)
	}
	// Use default session id creating function of storage.
	if s.id == "" {
		s.id = s.manager.storage.New(s.manager.ttl)
	}
	// Use default session id creating function.
	if s.id == "" {
		s.id = NewSessionId()
	}
	if s.data == nil {
		s.data = qn_map.NewStrAnyMap(true)
	}
	s.start = true
}

// Close closes current session and updates its ttl in the session manager.
// If this session is dirty, it also exports it to storage.
//
// NOTE that this function must be called ever after a session request done.
func (s *Session) Close() {
	if s.start && s.id != "" {
		size := s.data.Size()
		if s.manager.storage != nil {
			if s.dirty {
				if err := s.manager.storage.SetSession(s.id, s.data, s.manager.ttl); err != nil {
					panic(err)
				}
			} else if size > 0 {
				if err := s.manager.storage.UpdateTTL(s.id, s.manager.ttl); err != nil {
					panic(err)
				}
			}
		}
		if s.dirty || size > 0 {
			s.manager.UpdateSessionTTL(s.id, s.data)
		}
	}
}

// Set sets key-value pair to this session.
func (s *Session) Set(key string, value interface{}) error {
	s.init()
	if err := s.manager.storage.Set(s.id, key, value, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.Set(key, value)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Sets batch sets the session using map.
func (s *Session) Sets(data map[string]interface{}) error {
	s.init()
	if err := s.manager.storage.SetMap(s.id, data, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.Sets(data)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Remove removes key along with its value from this session.
func (s *Session) Remove(key string) error {
	if s.id == "" {
		return nil
	}
	s.init()
	if err := s.manager.storage.Remove(s.id, key); err != nil {
		if err == ErrorDisabled {
			s.data.Remove(key)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Clear is alias of RemoveAll.
func (s *Session) Clear() error {
	return s.RemoveAll()
}

// RemoveAll deletes all key-value pairs from this session.
func (s *Session) RemoveAll() error {
	if s.id == "" {
		return nil
	}
	s.init()
	if err := s.manager.storage.RemoveAll(s.id); err != nil {
		if err == ErrorDisabled {
			s.data.Clear()
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Id returns the session id for this session.
// It create and returns a new session id if the session id is not passed in initialization.
func (s *Session) Id() string {
	s.init()
	return s.id
}

// SetId sets custom session before session starts.
// It returns error if it is called after session starts.
func (s *Session) SetId(id string) error {
	if s.start {
		return errors.New("session already started")
	}
	s.id = id
	return nil
}

// SetIdFunc sets custom session id creating function before session starts.
// It returns error if it is called after session starts.
func (s *Session) SetIdFunc(f func(ttl time.Duration) string) error {
	if s.start {
		return errors.New("session already started")
	}
	s.idFunc = f
	return nil
}

// Map returns all data as map.
// Note that it's using value copy internally for concurrent-safe purpose.
func (s *Session) Map() map[string]interface{} {
	if s.id != "" {
		s.init()
		if data := s.manager.storage.GetMap(s.id); data != nil {
			return data
		}
		return s.data.Map()
	}
	return nil
}

// Size returns the size of the session.
func (s *Session) Size() int {
	if s.id != "" {
		s.init()
		if size := s.manager.storage.GetSize(s.id); size >= 0 {
			return size
		}
		return s.data.Size()
	}
	return 0
}

// Contains checks whether key exist in the session.
func (s *Session) Contains(key string) bool {
	s.init()
	return s.Get(key) != nil
}

// IsDirty checks whether there's any data changes in the session.
func (s *Session) IsDirty() bool {
	return s.dirty
}

// Get retrieves session value with given key.
// It returns <def> if the key does not exist in the session if <def> is given,
// or else it return nil.
func (s *Session) Get(key string, def ...interface{}) interface{} {
	if s.id == "" {
		return nil
	}
	s.init()
	if v := s.manager.storage.Get(s.id, key); v != nil {
		return v
	}
	if v := s.data.Get(key); v != nil {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

func (s *Session) GetVar(key string, def ...interface{}) qn_var.Var {
	return qn_var.New(s.Get(key, def...), true)
}

func (s *Session) GetString(key string, def ...interface{}) string {
	return qn_conv.String(s.Get(key, def...))
}

func (s *Session) GetBool(key string, def ...interface{}) bool {
	return qn_conv.Bool(s.Get(key, def...))
}

func (s *Session) GetInt(key string, def ...interface{}) int {
	return qn_conv.Int(s.Get(key, def...))
}

func (s *Session) GetInt8(key string, def ...interface{}) int8 {
	return qn_conv.Int8(s.Get(key, def...))
}

func (s *Session) GetInt16(key string, def ...interface{}) int16 {
	return qn_conv.Int16(s.Get(key, def...))
}

func (s *Session) GetInt32(key string, def ...interface{}) int32 {
	return qn_conv.Int32(s.Get(key, def...))
}

func (s *Session) GetInt64(key string, def ...interface{}) int64 {
	return qn_conv.Int64(s.Get(key, def...))
}

func (s *Session) GetUint(key string, def ...interface{}) uint {
	return qn_conv.Uint(s.Get(key, def...))
}

func (s *Session) GetUint8(key string, def ...interface{}) uint8 {
	return qn_conv.Uint8(s.Get(key, def...))
}

func (s *Session) GetUint16(key string, def ...interface{}) uint16 {
	return qn_conv.Uint16(s.Get(key, def...))
}

func (s *Session) GetUint32(key string, def ...interface{}) uint32 {
	return qn_conv.Uint32(s.Get(key, def...))
}

func (s *Session) GetUint64(key string, def ...interface{}) uint64 {
	return qn_conv.Uint64(s.Get(key, def...))
}

func (s *Session) GetFloat32(key string, def ...interface{}) float32 {
	return qn_conv.Float32(s.Get(key, def...))
}

func (s *Session) GetFloat64(key string, def ...interface{}) float64 {
	return qn_conv.Float64(s.Get(key, def...))
}

func (s *Session) GetBytes(key string, def ...interface{}) []byte {
	return qn_conv.Bytes(s.Get(key, def...))
}

func (s *Session) GetInts(key string, def ...interface{}) []int {
	return qn_conv.Ints(s.Get(key, def...))
}

func (s *Session) GetFloats(key string, def ...interface{}) []float64 {
	return qn_conv.Floats(s.Get(key, def...))
}

func (s *Session) GetStrings(key string, def ...interface{}) []string {
	return qn_conv.Strings(s.Get(key, def...))
}

func (s *Session) GetInterfaces(key string, def ...interface{}) []interface{} {
	return qn_conv.Interfaces(s.Get(key, def...))
}

func (s *Session) GetTime(key string, format ...string) time.Time {
	return qn_conv.Time(s.Get(key), format...)
}

func (s *Session) Getqn_time(key string, format ...string) *qn_time.Time {
	return qn_conv.qn_time(s.Get(key), format...)
}

func (s *Session) GetDuration(key string, def ...interface{}) time.Duration {
	return qn_conv.Duration(s.Get(key, def...))
}

func (s *Session) GetMap(key string, tags ...string) map[string]interface{} {
	return qn_conv.Map(s.Get(key), tags...)
}

func (s *Session) GetMapDeep(key string, tags ...string) map[string]interface{} {
	return qn_conv.MapDeep(s.Get(key), tags...)
}

func (s *Session) GetMaps(key string, tags ...string) []map[string]interface{} {
	return qn_conv.Maps(s.Get(key), tags...)
}

func (s *Session) GetMapsDeep(key string, tags ...string) []map[string]interface{} {
	return qn_conv.MapsDeep(s.Get(key), tags...)
}

func (s *Session) GetStruct(key string, pointer interface{}, mapping ...map[string]string) error {
	return qn_conv.Struct(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructDeep(key string, pointer interface{}, mapping ...map[string]string) error {
	return qn_conv.StructDeep(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructs(key string, pointer interface{}, mapping ...map[string]string) error {
	return qn_conv.Structs(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructsDeep(key string, pointer interface{}, mapping ...map[string]string) error {
	return qn_conv.StructsDeep(s.Get(key), pointer, mapping...)
}
