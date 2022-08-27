package repository

import (
	"context"
	"testing"

	"github.com/erik-sostenes/gofige/internal/model"
)

func TestStudentStorer_Insert(t *testing.T) {
	tsc := map[string]struct {
		studentStorer StudentStorer
		students      model.Students
		expectedError error
	}{
		"one": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			students: model.Students{
				{
					Name:    "some_name",
					Tuition: "some_tuition",
					Grade:   "some_grade",
					Group:   "some_grouop",
					Carrer:  "some_carreer",
				},
				{
					Name:    "some_name",
					Tuition: "some_tuition",
					Grade:   "some_grade",
					Group:   "some_grouop",
					Carrer:  "some_carreer",
				},
			},
			expectedError: nil,
		},
		"two": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			students: model.Students{
				{
					Name:    "some_name",
					Tuition: "some_tuition",
					Grade:   "some_grade",
					Group:   "some_grouop",
					Carrer:  "some_carreer",
				},
				{
					Name:    "some_name",
					Tuition: "some_tuition",
					Grade:   "some_grade",
					Group:   "some_grouop",
					Carrer:  "some_carreer",
				},
			},
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
