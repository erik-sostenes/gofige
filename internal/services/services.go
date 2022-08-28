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
		// Find find all data and returns a .csv file
		Find(context.Context) (model.Students, error)
		// FindByFlags find any collection that matches the flags
		// returns a .csv file
		FindByFlags(context.Context, model.Student) (model.Students, error)
	}
	// studentService implements StudentService interface
	studentService struct {
		fileService
		studentStorer repository.StudentStorer
	}
)

// NewStudentService returns a instance of StudentStorer interface
func NewStudentService(studentStorer repository.StudentStorer) StudentService {
	return &studentService{
		studentStorer: studentStorer,
	}
}

// Create creates an slice of students
func (s *studentService) Create(ctx context.Context, path string) (err error) {
	data, err := s.Read(path)
	if err != nil {
		err = fmt.Errorf("an error occurred while reading the file, check the directory %s", path)
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

func (s *studentService) Find(ctx context.Context) (model.Students, error) {
	return nil, nil
}
func (s *studentService) FindByFlags(ctx context.Context, studentFlags model.Student) (model.Students, error) {
	return nil, nil
}
