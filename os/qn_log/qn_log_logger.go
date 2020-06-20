// Copyright 2017 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/os/gfpool"
	"github.com/qnsoft/common/os/gmlock"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/os/qn_timer"
	qn_conv "github.com/qnsoft/common/util/qn_conv"

	"github.com/qnsoft/common/debug/qn_debug"

	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/text/qn_regex"
)

// Logger is the struct for logging management.
type Logger struct {
	rmu    sync.Mutex      // Mutex for rotation feature.
	ctx    context.Context // Context for logging.
	parent *Logger         // Parent logger, if it is not empty, it means the logger is used in chaining function.
	config Config          // Logger configuration.
}

const (
	gDEFAULT_FILE_FORMAT = `{Y-m-d}.log`
	gDEFAULT_FILE_FLAGS  = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	gDEFAULT_FILE_PERM   = os.FileMode(0666)
	gDEFAULT_FILE_EXPIRE = time.Minute
	gPATH_FILTER_KEY     = "/os/qn_log/qn_log"
)

const (
	F_ASYNC      = 1 << iota // Print logging content asynchronously。
	F_FILE_LONG              // Print full file name and line number: /a/b/c/d.go:23.
	F_FILE_SHORT             // Print final file name element and line number: d.go:23. overrides F_FILE_LONG.
	F_TIME_DATE              // Print the date in the local time zone: 2009-01-23.
	F_TIME_TIME              // Print the time in the local time zone: 01:23:23.
	F_TIME_MILLI             // Print the time with milliseconds in the local time zone: 01:23:23.675.
	F_TIME_STD   = F_TIME_DATE | F_TIME_MILLI
)

// New creates and returns a custom logger.
func New() *Logger {
	logger := &Logger{
		config: DefaultConfig(),
	}
	// Initialize the internal handler after some delay.
	qn_timer.AddOnce(time.Second, func() {
		qn_timer.AddOnce(logger.config.RotateCheckInterval, logger.rotateChecksTimely)
	})
	return logger
}

// NewWithWriter creates and returns a custom logger with io.Writer.
func NewWithWriter(writer io.Writer) *Logger {
	l := New()
	l.SetWriter(writer)
	return l
}

// Clone returns a new logger, which is the clone the current logger.
// It's commonly used for chaining operations.
func (l *Logger) Clone() *Logger {
	logger := Logger{}
	logger = *l
	logger.parent = l
	return &logger
}

// getFilePath returns the logging file path.
// The logging file name must have extension name of "log".
func (l *Logger) getFilePath(now time.Time) string {
	// Content containing "{}" in the file name is formatted using qn_time.
	file, _ := qn_regex.ReplaceStringFunc(`{.+?}`, l.config.File, func(s string) string {
		return qn_time.New(now).Format(strings.Trim(s, "{}"))
	})
	file = qn_file.Join(l.config.Path, file)
	if qn_file.ExtName(file) != "log" {
		file += ".log"
	}
	return file
}

// print prints <s> to defined writer, logging file or passed <std>.
func (l *Logger) print(std io.Writer, lead string, values ...interface{}) {
	var (
		now    = time.Now()
		buffer = bytes.NewBuffer(nil)
	)
	if l.config.HeaderPrint {
		// Time.
		timeFormat := ""
		if l.config.Flags&F_TIME_DATE > 0 {
			timeFormat += "2006-01-02 "
		}
		if l.config.Flags&F_TIME_TIME > 0 {
			timeFormat += "15:04:05 "
		}
		if l.config.Flags&F_TIME_MILLI > 0 {
			timeFormat += "15:04:05.000 "
		}
		if len(timeFormat) > 0 {
			buffer.WriteString(now.Format(timeFormat))
		}
		// Lead string.
		if len(lead) > 0 {
			buffer.WriteString(lead)
			if len(values) > 0 {
				buffer.WriteByte(' ')
			}
		}
		// Caller path.
		callerPath := ""
		if l.config.Flags&F_FILE_LONG > 0 {
			_, path, line := qn_debug.CallerWithFilter(gPATH_FILTER_KEY, l.config.StSkip)
			callerPath = fmt.Sprintf(`%s:%d: `, path, line)
		}
		if l.config.Flags&F_FILE_SHORT > 0 {
			_, path, line := qn_debug.CallerWithFilter(gPATH_FILTER_KEY, l.config.StSkip)
			callerPath = fmt.Sprintf(`%s:%d: `, qn_file.Basename(path), line)
		}
		if len(callerPath) > 0 {
			buffer.WriteString(callerPath)
		}
		// Prefix.
		if len(l.config.Prefix) > 0 {
			buffer.WriteString(l.config.Prefix + " ")
		}
	}
	// Convert value to string.
	var (
		tempStr  = ""
		valueStr = ""
	)
	// Context values.
	if l.ctx != nil && len(l.config.CtxKeys) > 0 {
		ctxStr := ""
		for _, key := range l.config.CtxKeys {
			if v := l.ctx.Value(key); v != nil {
				if ctxStr != "" {
					ctxStr += ", "
				}
				ctxStr += fmt.Sprintf("%s: %+v", key, v)
			}
		}
		if ctxStr != "" {
			buffer.WriteString(fmt.Sprintf("{%s} ", ctxStr))
		}
	}
	for _, v := range values {
		if err, ok := v.(error); ok {
			tempStr = fmt.Sprintf("%+v", err)
		} else {
			tempStr = qn_conv.String(v)
		}
		if len(valueStr) > 0 {
			if valueStr[len(valueStr)-1] == '\n' {
				// Remove one blank line(\n\n).
				if tempStr[0] == '\n' {
					valueStr += tempStr[1:]
				} else {
					valueStr += tempStr
				}
			} else {
				valueStr += " " + tempStr
			}
		} else {
			valueStr = tempStr
		}
	}
	buffer.WriteString(valueStr + "\n")
	if l.config.Flags&F_ASYNC > 0 {
		err := asyncPool.Add(func() {
			l.printToWriter(now, std, buffer)
		})
		if err != nil {
			intlog.Error(err)
		}
	} else {
		l.printToWriter(now, std, buffer)
	}
}

// printToWriter writes buffer to writer.
func (l *Logger) printToWriter(now time.Time, std io.Writer, buffer *bytes.Buffer) {
	if l.config.Writer == nil {
		// Output content to disk file.
		if l.config.Path != "" {
			l.printToFile(now, buffer)
		}
		// Allow output to stdout?
		if l.config.StdoutPrint {
			if _, err := std.Write(buffer.Bytes()); err != nil {
				intlog.Error(err)
			}
		}
	} else {
		if _, err := l.config.Writer.Write(buffer.Bytes()); err != nil {
			panic(err)
		}
	}
}

// printToFile outputs logging content to disk file.
func (l *Logger) printToFile(now time.Time, buffer *bytes.Buffer) {
	var (
		loqn_filePath = l.getFilePath(now)
		memoryLockKey = "qn_log.file.lock:" + loqn_filePath
	)
	gmlock.Lock(memoryLockKey)
	defer gmlock.Unlock(memoryLockKey)
	file := l.getFilePointer(loqn_filePath)
	defer file.Close()
	// Rotation file size checks.
	if l.config.RotateSize > 0 {
		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if stat.Size() > l.config.RotateSize {
			l.rotateFileBySize(now)
			file = l.getFilePointer(loqn_filePath)
			defer file.Close()
		}
	}
	if _, err := file.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}

// getFilePointer retrieves and returns a file pointer from file pool.
func (l *Logger) getFilePointer(path string) *gfpool.File {
	file, err := gfpool.Open(
		path,
		gDEFAULT_FILE_FLAGS,
		gDEFAULT_FILE_PERM,
		gDEFAULT_FILE_EXPIRE,
	)
	if err != nil {
		panic(err)
	}
	return file
}

// printStd prints content <s> without stack.
func (l *Logger) printStd(lead string, value ...interface{}) {
	l.print(os.Stdout, lead, value...)
}

// printStd prints content <s> with stack check.
func (l *Logger) printErr(lead string, value ...interface{}) {
	if l.config.StStatus == 1 {
		if s := l.GetStack(); s != "" {
			value = append(value, "\nStack:\n"+s)
		}
	}
	// In matter of sequence, do not use stderr here, but use the same stdout.
	l.print(os.Stdout, lead, value...)
}

// format formats <values> using fmt.Sprintf.
func (l *Logger) format(format string, value ...interface{}) string {
	return fmt.Sprintf(format, value...)
}

// PrintStack prints the caller stack,
// the optional parameter <skip> specify the skipped stack offset from the end point.
func (l *Logger) PrintStack(skip ...int) {
	if s := l.GetStack(skip...); s != "" {
		l.Println("Stack:\n" + s)
	} else {
		l.Println()
	}
}

// GetStack returns the caller stack content,
// the optional parameter <skip> specify the skipped stack offset from the end point.
func (l *Logger) GetStack(skip ...int) string {
	stackSkip := l.config.StSkip
	if len(skip) > 0 {
		stackSkip += skip[0]
	}
	filters := []string{gPATH_FILTER_KEY}
	if l.config.StFilter != "" {
		filters = append(filters, l.config.StFilter)
	}
	return qn_debug.StackWithFilters(filters, stackSkip)
}
