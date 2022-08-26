package services

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/erik-sostenes/gofige/internal/model"
)
// Create contains the methods that are responsible for verifying
// that the business logic is correct
type Create struct {
	File
}

// Create creates an slice of students
func (c *Create) Create() error {
	data, err := c.File.Read()

	r := csv.NewReader(strings.NewReader(string(data)))

	var students model.Students

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		student := model.Student{
			Name:    record[0],
			Tuition: record[1],
			Grade:   record[2],
			Group:   record[3],
			Carrer:  record[4],
		}
		students = append(students, student)
	}
	return err
}
