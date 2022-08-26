package services

import (
	"io/fs"
	"io/ioutil"
)

// File contains methods for working with files
type File struct {
	Path string
}

// Read reads a file through the path
func (f *File) Read() ([]byte, error) {
	return ioutil.ReadFile(f.Path)
}

// Write writes a file with the assigned data
func (f *File) Write(data []byte) error {
	var fileMode fs.FileMode
	return ioutil.WriteFile(f.Path, data, fileMode)
}
