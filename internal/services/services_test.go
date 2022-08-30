package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/erik-sostenes/gofige/internal/model"
	"github.com/erik-sostenes/gofige/internal/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStudentService_Find(t *testing.T) {
	var (
		mockStudentOne = model.Student{
			Tuition: "X189HR",
			Name:    "Erik Sostenes",
			Grade:   "5",
			Group:   "A",
			Carrer:  "Ingenieria",
		}
		mockStudentTwo = model.Student{
			Tuition: "EVS765",
			Name:    "Erik Simon",
			Grade:   "7",
			Group:   "B",
			Carrer:  "Sistemas",
		}
		mockStudents = model.Students{
			mockStudentOne,
			mockStudentTwo,
		}
	)

	tsc := map[string]struct {
		studentService  StudentService
		path            string
		mockStudent     model.Student
		expetedStudents model.Students
		expetedError    error
	}{
		"Given a successful search, it returns a slice of students": {
			studentService: NewStudentService(
				&repository.MockStudentStorer{MockStudents: mockStudents},
			),
			path:            "C:/Users/eriks/OneDrive/Escritorio/students.csv",
			mockStudent:     mockStudentOne,
			expetedStudents: mockStudents,
		},
		"Given an incorrect search, returns a mongo.ErrNoDocuments error": {
			studentService: NewStudentService(
				&repository.MockStudentStorer{Error: mongo.ErrNoDocuments},
			),
			path:         "C:/Users/eriks/OneDrive/Escritorio/students.csv",
			mockStudent:  mockStudentOne,
			expetedError: mongo.ErrNoDocuments,
		},
		"Given an incorrect path, returns a WrongFile error": {
			studentService: NewStudentService(
				&repository.MockStudentStorer{},
			),
			path: "C:/Users/eriks/OneDrive/Escritorio/some.csv",
			expetedError: model.WrongFile(
				fmt.Sprintf("an error occurred while reading the file, check the directory %s",
					"C:/Users/eriks/OneDrive/Escritorio/some.csv",
				),
			),
		},
	}

	for name, ts := range tsc {
		ts := ts
		t.Run(name, func(t *testing.T) {
			mock, err := ts.studentService.Find(context.Background(), ts.path, ts.mockStudent)
			if !errors.Is(err, ts.expetedError) {
				t.Errorf("expeted error %T, got error %T", ts.expetedError, err)
				t.SkipNow()
			}

			if !reflect.DeepEqual(mock, ts.expetedStudents) {
				t.Errorf("expeted %v,\n got %v", ts.expetedStudents, mock)
			}
		})
	}
}
