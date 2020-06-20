// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_i18n

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/qnsoft/common/internal/intlog"

	"github.com/qnsoft/common/os/glog"

	"github.com/qnsoft/common/os/gfsnotify"

	"github.com/qnsoft/common/text/qn_regex"

	qn_conv "github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/encoding/qn_json"

	"github.com/qnsoft/common/os/gres"
	"github.com/qnsoft/common/os/qn_file"
)

// Manager, it is concurrent safe, supporting hot reload.
type Manager struct {
	mu      sync.RWMutex
	data    map[string]map[string]string // Translating map.
	pattern string                       // Pattern for regex parsing.
	options Options                      // configuration options.
}

// Options is used for i18n object configuration.
type Options struct {
	Path       string   // I18n files storage path.
	Language   string   // Local language.
	Delimiters []string // Delimiters for variable parsing.
}

var (
	// defaultDelimiters defines the default key variable delimiters.
	defaultDelimiters = []string{"{#", "}"}
)

// New creates and returns a new i18n manager.
// The optional parameter <option> specifies the custom options for i18n manager.
// It uses a default one if it's not passed.
func New(options ...Options) *Manager {
	var opts Options
	if len(options) > 0 {
		opts = options[0]
	} else {
		opts = DefaultOptions()
	}
	if len(opts.Delimiters) == 0 {
		opts.Delimiters = defaultDelimiters
	}
	m := &Manager{
		options: opts,
		pattern: fmt.Sprintf(
			`%s(\w+)%s`,
			qn_regex.Quote(opts.Delimiters[0]),
			qn_regex.Quote(opts.Delimiters[1]),
		),
	}
	intlog.Printf(`New: %#v`, m)
	return m
}

// DefaultOptions creates and returns a default options for i18n manager.
func DefaultOptions() Options {
	var (
		path        = "i18n"
		realPath, _ = qn_file.Search(path)
	)
	if realPath != "" {
		path = realPath
		// To avoid of the source path of GF: github.com/gogf/i18n/gi18n
		if qn_file.Exists(path + qn_file.Separator + "gi18n") {
			path = ""
		}
	}
	return Options{
		Path:       path,
		Language:   "en",
		Delimiters: defaultDelimiters,
	}
}

// SetPath sets the directory path storing i18n files.
func (m *Manager) SetPath(path string) error {
	if gres.Contains(path) {
		m.options.Path = path
	} else {
		realPath, _ := qn_file.Search(path)
		if realPath == "" {
			return errors.New(fmt.Sprintf(`%s does not exist`, path))
		}
		m.options.Path = realPath
	}
	intlog.Printf(`SetPath: %s`, m.options.Path)
	return nil
}

// SetLanguage sets the language for translator.
func (m *Manager) SetLanguage(language string) {
	m.options.Language = language
	intlog.Printf(`SetLanguage: %s`, m.options.Language)
}

// SetDelimiters sets the delimiters for translator.
func (m *Manager) SetDelimiters(left, right string) {
	m.pattern = fmt.Sprintf(`%s(\w+)%s`, qn_regex.Quote(left), qn_regex.Quote(right))
	intlog.Printf(`SetDelimiters: %v`, m.pattern)
}

// T is alias of Translate for convenience.
func (m *Manager) T(content string, language ...string) string {
	return m.Translate(content, language...)
}

// Translate translates <content> with configured language.
// The parameter <language> specifies custom translation language ignoring configured language.
func (m *Manager) Translate(content string, language ...string) string {
	m.init()
	m.mu.RLock()
	defer m.mu.RUnlock()
	transLang := m.options.Language
	if len(language) > 0 && language[0] != "" {
		transLang = language[0]
	} else {
		transLang = m.options.Language
	}
	data := m.data[transLang]
	if data == nil {
		return content
	}
	// Parse content as name.
	if v, ok := data[content]; ok {
		return v
	}
	// Parse content as variables container.
	result, _ := qn_regex.ReplaceStringFuncMatch(
		m.pattern, content,
		func(match []string) string {
			if v, ok := data[match[1]]; ok {
				return v
			}
			return match[0]
		})
	intlog.Printf(`Translate for language: %s`, transLang)
	return result
}

// GetValue retrieves and returns the configured content for given key and specified language.
// It returns an empty string if not found.
func (m *Manager) GetContent(key string, language ...string) string {
	m.init()
	m.mu.RLock()
	defer m.mu.RUnlock()
	transLang := m.options.Language
	if len(language) > 0 && language[0] != "" {
		transLang = language[0]
	} else {
		transLang = m.options.Language
	}
	if data, ok := m.data[transLang]; ok {
		return data[key]
	}
	return ""
}

// init initializes the manager for lazy initialization design.
// The i18n manager is only initialized once.
func (m *Manager) init() {
	m.mu.RLock()
	// If the data is not nil, means it's already initialized.
	if m.data != nil {
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()

	m.mu.Lock()
	defer m.mu.Unlock()
	if gres.Contains(m.options.Path) {
		files := gres.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) > 0 {
			var (
				path  string
				name  string
				lang  string
				array []string
			)
			m.data = make(map[string]map[string]string)
			for _, file := range files {
				name = file.Name()
				path = name[len(m.options.Path)+1:]
				array = strings.Split(path, "/")
				if len(array) > 1 {
					lang = array[0]
				} else {
					lang = qn_file.Name(array[0])
				}
				if m.data[lang] == nil {
					m.data[lang] = make(map[string]string)
				}
				if j, err := qn_json.LoadContent(file.Content()); err == nil {
					for k, v := range j.ToMap() {
						m.data[lang][k] = qn_conv.String(v)
					}
				} else {
					glog.Errorf("load i18n file '%s' failed: %v", name, err)
				}
			}
		}
	} else if m.options.Path != "" {
		files, _ := qn_file.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) == 0 {
			intlog.Printf(
				"no i18n files found in configured directory: %s",
				m.options.Path,
			)
			return
		}
		var (
			path  string
			lang  string
			array []string
		)
		m.data = make(map[string]map[string]string)
		for _, file := range files {
			path = file[len(m.options.Path)+1:]
			array = strings.Split(path, qn_file.Separator)
			if len(array) > 1 {
				lang = array[0]
			} else {
				lang = qn_file.Name(array[0])
			}
			if m.data[lang] == nil {
				m.data[lang] = make(map[string]string)
			}
			if j, err := qn_json.LoadContent(qn_file.GetBytes(file)); err == nil {
				for k, v := range j.ToMap() {
					m.data[lang][k] = qn_conv.String(v)
				}
			} else {
				glog.Errorf("load i18n file '%s' failed: %v", file, err)
			}
		}
		// Monitor changes of i18n files for hot reload feature.
		_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
			// Any changes of i18n files, clear the data.
			m.mu.Lock()
			m.data = nil
			m.mu.Unlock()
			gfsnotify.Exit()
		})
	}
}
