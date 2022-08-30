package repository

import (
	"context"
	"testing"

	"github.com/erik-sostenes/gofige/internal/model"
)

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
		students      model.Students
		expectedError error
	}{
		"Given a successful insertion in mongodb, the data is stored in the collection students_one": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students_one")),
			students:      mockStudents,
			expectedError: nil,
		},
		"Given a successful insertion in mongodb, the data is stored in the collection students_two": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students_two")),
			students:      mockStudents,
			expectedError: nil,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			err := ts.studentStorer.Insert(context.TODO(), ts.students)
			if err != nil {
				t.Errorf("expected error %s, got error %s", ts.expectedError, err)
			}
		})
	}
}
