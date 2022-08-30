package services

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/erik-sostenes/gofige/internal/model"
	"github.com/erik-sostenes/gofige/internal/repository"
)

type (
	// StudentService contains the methods that are responsible for verifying that the business logic is correct
	StudentService interface {
		Create(context.Context, string) error
		// Find find any collection that matches the flags
		// returns a .csv file
		Find(context.Context, string, model.Student) (model.Students, error)
	}
	// studentService implements StudentService interface
	studentService struct {
		fileService
		flags
		studentStorer repository.StudentStorer
	}
)

// NewStudentService returns a instance of StudentStorer interface
func NewStudentService(studentStorer repository.StudentStorer) StudentService {
	f := make(flags)
	return &studentService{
		flags:         f,
		studentStorer: studentStorer,
	}
}

// Create creates an slice of students
func (s *studentService) Create(ctx context.Context, path string) (err error) {
	data, err := s.Read(path)
	if err != nil {
		return
	}

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
	err = s.studentStorer.Insert(ctx, students[1:])
	return
}

func (s *studentService) Find(ctx context.Context, path string, flags model.Student) (students model.Students, err error) {
	m, err := s.flags.UnmarshalFlags(flags)
	if err != nil {
		return
	}

	students, err = s.studentStorer.Find(ctx, m)
	if err != nil {
		return
	}

	csvFile, err := s.Write(path)
	if err != nil {
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, v := range students {
		row := []string{
			v.Tuition,
			v.Name,
			v.Grade,
			v.Group,
			v.Carrer,
		}
		if err = csvWriter.Write(row); err != nil {
			err = fmt.Errorf("error writing record to file %s", err)
		}
	}
	return
}
