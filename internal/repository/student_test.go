package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erik-sostenes/gofige/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type arguments map[string]string

func TestStudentStorer_Insert(t *testing.T) {
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
		studentStorer StudentStorer
		mockStudents  model.Students
		expectedError error
	}{
		"Given a successful insertion in mongodb, the data is stored in the collection students": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			mockStudents:  mockStudents,
			expectedError: nil,
		},
		"Given a successful insertion in mongodb, the data is stored in the collection students_two": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students_two")),
			mockStudents:  mockStudents,
			expectedError: nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			err := ts.studentStorer.Insert(context.TODO(), ts.mockStudents)
			if err != ts.expectedError {
				t.Errorf("expected error %s, got error %s", ts.expectedError, err)
			}
		})
	}
}

func TestStudentStorer_Find(t *testing.T) {
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
	)

	tsc := map[string]struct {
		studentStorer   StudentStorer
		expetedStudents model.Students
		arguments       map[string]string
		expectedError   error
	}{
		"Given a successful search, it returns a slice of students": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			expetedStudents: model.Students{
				mockStudentOne,
				mockStudentTwo,
			},
		},
		"Given a successful search with arguments, it returns a slice of students": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			expetedStudents: model.Students{
				mockStudentOne,
			},
			arguments: arguments{"tuition": mockStudentOne.Tuition},
		},
		"Given a search that the collection does not contains data, returns a mongo.ErrNoDocuments error": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			expectedError: *&mongo.ErrNoDocuments,
		},
	}

	for name, ts := range tsc {
		ts := ts
		t.Run(name, func(t *testing.T) {
			_ = ts.studentStorer.Insert(context.TODO(), ts.expetedStudents)

			t.Cleanup(func() {
				_ = ts.studentStorer.Delete(context.TODO(), ts.arguments)
			})

			mock, err := ts.studentStorer.Find(context.TODO(), ts.arguments)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("expected error %s, got error %s", ts.expectedError, err)
				t.SkipNow()
			}

			if !reflect.DeepEqual(mock, ts.expetedStudents) {
				t.Errorf("expeted %v,\n got %v", ts.expetedStudents, mock)
			}
		})
	}
}
