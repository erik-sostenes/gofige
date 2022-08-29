package services

import (
	"fmt"
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

// Write writes a .csv file
func (f *fileService) Write(path string) (csvFile *os.File, err error) {
	if err = f.checkFilePath(path); err != nil {
		return
	}
	csvFile, err = os.Create(path)
	return
}

// checkFilePath verify that the path is valid
func (f *fileService) checkFilePath(path string) (err error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.IsDir() {
		err = fmt.Errorf("an error occurred while reading the file, check the directory %s", path)
		return
	}
	return
}
