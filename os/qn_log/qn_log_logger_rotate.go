// Copyright 2020 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_log

import (
	"fmt"
	"time"

	"github.com/qnsoft/common/container/qn_array"
	"github.com/qnsoft/common/encoding/qn_compress"
	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/os/qn_timer"
	"github.com/qnsoft/common/text/qn_regex"
)

// rotateFileBySize rotates the current logging file according to the
// configured rotation size.
func (l *Logger) rotateFileBySize(now time.Time) {
	if l.config.RotateSize <= 0 {
		return
	}
	l.rmu.Lock()
	defer l.rmu.Unlock()
	if err := l.doRotateFile(l.getFilePath(now)); err != nil {
		panic(err)
	}
}

// doRotateFile rotates the given logging file.
func (l *Logger) doRotateFile(filePath string) error {
	// No backups, it then just removes the current logging file.
	if l.config.RotateBackupLimit == 0 {
		if err := qn_file.Remove(filePath); err != nil {
			return err
		}
		intlog.Printf(`%d size exceeds, no backups set, remove original logging file: %s`, l.config.RotateSize, filePath)
		return nil
	}
	// Else it creates new backup files.
	var (
		dirPath     = qn_file.Dir(filePath)
		fileName    = qn_file.Name(filePath)
		fileExtName = qn_file.ExtName(filePath)
		newFilePath = ""
	)
	// Rename the logging file by adding extra datetime information to microseconds, like:
	// access.log          -> access.20200326101301899002.log
	// access.20200326.log -> access.20200326.20200326101301899002.log
	for {
		var (
			now   = qn_time.Now()
			micro = now.Microsecond() % 1000
		)
		for micro < 100 {
			micro *= 10
		}
		newFilePath = qn_file.Join(
			dirPath,
			fmt.Sprintf(
				`%s.%s%d.%s`,
				fileName, now.Format("YmdHisu"), micro, fileExtName,
			),
		)
		if !qn_file.Exists(newFilePath) {
			break
		}
	}
	if err := qn_file.Rename(filePath, newFilePath); err != nil {
		return err
	}
	return nil
}

// rotateChecksTimely timely checks the backups expiration and the compression.
func (l *Logger) rotateChecksTimely() {
	defer qn_timer.AddOnce(l.config.RotateCheckInterval, l.rotateChecksTimely)
	// Checks whether file rotation not enabled.
	if l.config.RotateSize <= 0 && l.config.RotateExpire == 0 {
		return
	}
	var (
		now      = time.Now()
		pattern  = "*.log, *.gz"
		files, _ = qn_file.ScanDirFile(l.config.Path, pattern, true)
	)
	intlog.Printf("logging rotation start checks: %+v", files)
	// =============================================================
	// Rotation expire file checks.
	// =============================================================
	if l.config.RotateExpire > 0 {
		var (
			mtime         time.Time
			subDuration   time.Duration
			expireRotated bool
		)
		for _, file := range files {
			if qn_file.ExtName(file) == "gz" {
				continue
			}
			mtime = qn_file.MTime(file)
			subDuration = now.Sub(mtime)
			if subDuration > l.config.RotateExpire {
				expireRotated = true
				intlog.Printf(
					`%v - %v = %v > %v, rotation expire logging file: %s`,
					now, mtime, subDuration, l.config.RotateExpire, file,
				)
				if err := l.doRotateFile(file); err != nil {
					intlog.Error(err)
				}
			}
		}
		if expireRotated {
			// Update the files array.
			files, _ = qn_file.ScanDirFile(l.config.Path, pattern, true)
		}
	}

	// =============================================================
	// Rotated file compression.
	// =============================================================
	needCompressFileArray := qn_array.NewStrArray()
	if l.config.RotateBackupCompress > 0 {
		for _, file := range files {
			// Eg: access.20200326101301899002.log.gz
			if qn_file.ExtName(file) == "gz" {
				continue
			}
			// Eg:
			// access.20200326101301899002.log
			if qn_regex.IsMatchString(`.+\.\d{20}\.log`, qn_file.Basename(file)) {
				needCompressFileArray.Append(file)
			}
		}
		if needCompressFileArray.Len() > 0 {
			needCompressFileArray.Iterator(func(_ int, path string) bool {
				err := qn_compress.GzipFile(path, path+".gz")
				if err == nil {
					intlog.Printf(`compressed done, remove original logging file: %s`, path)
					if err = qn_file.Remove(path); err != nil {
						intlog.Print(err)
					}
				} else {
					intlog.Print(err)
				}
				return true
			})
			// Update the files array.
			files, _ = qn_file.ScanDirFile(l.config.Path, pattern, true)
		}
	}

	// =============================================================
	// Backups count limit and expiration checks.
	// =============================================================
	var (
		backupFilesMap            = make(map[string]*qn_array.SortedArray)
		originalLogginqn_filePath = ""
	)
	if l.config.RotateBackupLimit > 0 || l.config.RotateBackupExpire > 0 {
		for _, file := range files {
			originalLogginqn_filePath, _ = qn_regex.ReplaceString(`\.\d{20}`, "", file)
			if backupFilesMap[originalLogginqn_filePath] == nil {
				backupFilesMap[originalLogginqn_filePath] = qn_array.NewSortedArray(func(a, b interface{}) int {
					// Sorted by rotated/backup file mtime.
					// The old rotated/backup file is put in the head of array.
					file1 := a.(string)
					file2 := b.(string)
					result := qn_file.MTimestampMilli(file1) - qn_file.MTimestampMilli(file2)
					if result <= 0 {
						return -1
					}
					return 1
				})
			}
			// Check if this file a rotated/backup file.
			if qn_regex.IsMatchString(`.+\.\d{20}\.log`, qn_file.Basename(file)) {
				backupFilesMap[originalLogginqn_filePath].Add(file)
			}
		}
		intlog.Printf(`calculated backup files map: %+v`, backupFilesMap)
		for _, array := range backupFilesMap {
			diff := array.Len() - l.config.RotateBackupLimit
			for i := 0; i < diff; i++ {
				path, _ := array.PopLeft()
				intlog.Printf(`remove exceeded backup limit file: %s`, path)
				if err := qn_file.Remove(path.(string)); err != nil {
					intlog.Print(err)
				}
			}
		}
		// Backup expiration checks.
		if l.config.RotateBackupExpire > 0 {
			var (
				mtime       time.Time
				subDuration time.Duration
			)
			for _, array := range backupFilesMap {
				array.Iterator(func(_ int, v interface{}) bool {
					path := v.(string)
					mtime = qn_file.MTime(path)
					subDuration = now.Sub(mtime)
					if subDuration > l.config.RotateBackupExpire {
						intlog.Printf(
							`%v - %v = %v > %v, remove expired backup file: %s`,
							now, mtime, subDuration, l.config.RotateBackupExpire, path,
						)
						if err := qn_file.Remove(path); err != nil {
							intlog.Print(err)
						}
						return true
					} else {
						return false
					}
				})
			}
		}
	}
}
