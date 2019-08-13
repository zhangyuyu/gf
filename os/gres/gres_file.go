// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gres

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type File struct {
	zipFile *zip.File
}

// Name returns the name of the file.
func (f *File) Name() string {
	return f.zipFile.Name
}

// Open returns a ReadCloser that provides access to the File's contents.
// Multiple files may be read concurrently.
func (f *File) Open() (io.ReadCloser, error) {
	return f.zipFile.Open()
}

// Content returns the content of the file.
func (f *File) Content() ([]byte, error) {
	reader, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, reader); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// FileInfo returns an os.FileInfo for the FileHeader.
func (f *File) FileInfo() os.FileInfo {
	return f.zipFile.FileInfo()
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (f *File) MarshalJSON() ([]byte, error) {
	info := f.FileInfo()
	return json.Marshal(map[string]interface{}{
		"name": f.Name(),
		"size": info.Size(),
		"time": info.ModTime(),
	})
}