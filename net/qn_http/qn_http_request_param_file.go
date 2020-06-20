// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_http

import (
	"errors"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/qnsoft/common/internal/intlog"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/util/qn_rand"
)

// UploadFile wraps the multipart uploading file with more and convenient features.
type UploadFile struct {
	*multipart.FileHeader
}

// UploadFiles is array type for *UploadFile.
type UploadFiles []*UploadFile

// Save saves the single uploading file to directory path and returns the saved file name.
//
// The parameter <dirPath> should be a directory path or it returns error.
//
// Note that it will OVERWRITE the target file if there's already a same name file exist.
func (f *UploadFile) Save(dirPath string, randomlyRename ...bool) (filename string, err error) {
	if f == nil {
		return "", errors.New("file is empty, maybe you retrieve it from invalid field name or form enctype")
	}
	if !qn_file.Exists(dirPath) {
		if err = qn_file.Mkdir(dirPath); err != nil {
			return
		}
	} else if !qn_file.IsDir(dirPath) {
		return "", errors.New(`parameter "dirPath" should be a directory path`)
	}

	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	name := qn_file.Basename(f.Filename)
	if len(randomlyRename) > 0 && randomlyRename[0] {
		name = strings.ToLower(strconv.FormatInt(qn_time.TimestampNano(), 36) + qn_rand.S(6))
		name = name + qn_file.Ext(f.Filename)
	}
	filePath := qn_file.Join(dirPath, name)
	newFile, err := qn_file.Create(filePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	intlog.Printf(`save upload file: %s`, filePath)
	if _, err := io.Copy(newFile, file); err != nil {
		return "", err
	}
	return qn_file.Basename(filePath), nil
}

// Save saves all uploading files to specified directory path and returns the saved file names.
//
// The parameter <dirPath> should be a directory path or it returns error.
//
// The parameter <randomlyRename> specifies whether randomly renames all the file names.
func (fs UploadFiles) Save(dirPath string, randomlyRename ...bool) (filenames []string, err error) {
	if len(fs) == 0 {
		return nil, errors.New("file array is empty, maybe you retrieve it from invalid field name or form enctype")
	}
	for _, f := range fs {
		if filename, err := f.Save(dirPath, randomlyRename...); err != nil {
			return filenames, err
		} else {
			filenames = append(filenames, filename)
		}
	}
	return
}

// GetUploadFile retrieves and returns the uploading file with specified form name.
// This function is used for retrieving single uploading file object, which is
// uploaded using multipart form content type.
//
// It returns nil if retrieving failed or no form file with given name posted.
//
// Note that the <name> is the file field name of the multipart form from client.
func (r *Request) GetUploadFile(name string) *UploadFile {
	uploadFiles := r.GetUploadFiles(name)
	if len(uploadFiles) > 0 {
		return uploadFiles[0]
	}
	return nil
}

// GetUploadFiles retrieves and returns multiple uploading files with specified form name.
// This function is used for retrieving multiple uploading file objects, which are
// uploaded using multipart form content type.
//
// It returns nil if retrieving failed or no form file with given name posted.
//
// Note that the <name> is the file field name of the multipart form from client.
func (r *Request) GetUploadFiles(name string) UploadFiles {
	multipartFiles := r.GetMultipartFiles(name)
	if len(multipartFiles) > 0 {
		uploadFiles := make(UploadFiles, len(multipartFiles))
		for k, v := range multipartFiles {
			uploadFiles[k] = &UploadFile{
				FileHeader: v,
			}
		}
		return uploadFiles
	}
	return nil
}
