package services

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

// file contains methods for working with files
type fileService struct{}

// Read reads a file through the path
func (f *fileService) Read(path string) (data []byte, err error) {
	if err = f.checkFilePath(path); err != nil {
		return
	}
	return ioutil.ReadFile(path)
}

// Write writes a file with the assigned data
func (f *fileService) Write(path string, data []byte) (err error) {
	if err = f.checkFilePath(path); err != nil {
		return
	}
	var fileMode fs.FileMode
	return ioutil.WriteFile(path, data, fileMode)

}

// checkFilePath verify that the path is valid
func (f *fileService) checkFilePath(path string) (err error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.Mode().IsDir() {
		err = fmt.Errorf("The path does not exist %s", path)
		return
	}
	return
}
