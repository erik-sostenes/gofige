package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erik-sostenes/gofige/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStudentStorer_Insert(t *testing.T) {
	var (
		mockStudentOne = model.Student{
			Tuition: "X189HR",
			Name:    "Erik Sostenes",
			Grade:   "5",
			Group:   "A",
			Career:  "Ingenieria",
		}
		mockStudentTwo = model.Student{
			Tuition: "EVS765",
			Name:    "Erik Simon",
			Grade:   "7",
			Group:   "B",
			Career:  "Sistemas",
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
			t.Cleanup(func() {
				_ = ts.studentStorer.Delete(context.TODO(), bson.M{})
			})
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
			Career:  "Ingenieria",
		}
		mockStudentTwo = model.Student{
			Tuition: "EVS765",
			Name:    "Erik Simon",
			Grade:   "7",
			Group:   "B",
			Career:  "Sistemas",
		}
	)

	tsc := map[string]struct {
		studentStorer   StudentStorer
		expetedStudents model.Students
		arguments       bson.M
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
			arguments: bson.M{"tuition": mockStudentOne.Tuition},
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

func TestStudentStorer_Delete(t *testing.T) {
	var (
		mockStudentOne = model.Student{
			Tuition: "X189HR",
			Name:    "Erik Sostenes",
			Grade:   "5",
			Group:   "A",
			Career:  "Ingenieria",
		}
		mockStudentTwo = model.Student{
			Tuition: "EVS765",
			Name:    "Erik Simon",
			Grade:   "7",
			Group:   "B",
			Career:  "Sistemas",
		}
	)

	tsc := map[string]struct {
		studentStorer   StudentStorer
		expetedStudents model.Students
		arguments       bson.M
		expectedError   error
	}{
		"Given a successful delete, deletes a collection in mongodb": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			expetedStudents: model.Students{
				mockStudentOne,
				mockStudentTwo,
			},
		},
		"Given a wrong delete, it returns error": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			expetedStudents: model.Students{
				mockStudentOne,
			},
			arguments: bson.M{"tuition": mockStudentOne.Tuition},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_ = ts.studentStorer.Insert(context.TODO(), ts.expetedStudents)
			err := ts.studentStorer.Delete(context.TODO(), ts.arguments)
			if err != nil {
				t.Errorf("expeted %v,\n got %v", ts.expectedError, err)
			}
		})
	}
}

func TestStudentStorer_Update(t *testing.T) {
	var (
		mockStudentOne = model.Student{
			Tuition: "X189HR",
			Name:    "Erik Sostenes",
			Grade:   "5",
			Group:   "A",
			Career:  "Ingenieria",
		}
		mockStudentTwo = model.Student{
			Tuition: "EVS765",
			Name:    "Erik Simon",
			Grade:   "7",
			Group:   "B",
			Career:  "Sistemas",
		}
	)

	tsc := map[string]struct {
		studentStorer   StudentStorer
		mockStudents    model.Students
		filters         bson.M
		arguments       bson.D
		expetedStudents model.Students
		expectedError   error
	}{
		"Given a successful update, the data is updated in the collection": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			mockStudents: model.Students{
				mockStudentOne,
			},
			filters: bson.M{"tuition": mockStudentOne.Tuition},
			arguments: bson.D{{
				"$set",
				bson.D{
					{"tuition", mockStudentTwo.Tuition},
					{"name", mockStudentTwo.Name},
					{"grade", mockStudentTwo.Grade},
					{"group", mockStudentTwo.Group},
					{"career", mockStudentTwo.Career},
				},
			}},
			expetedStudents: model.Students{
				mockStudentTwo,
			},
		},
		"Given a successful update, the  data is updated in the collection": {
			studentStorer: NewStudentStorer(NewMDB(Config).Collection("students")),
			mockStudents: model.Students{
				mockStudentTwo,
			},
			filters: bson.M{"tuition": mockStudentTwo.Tuition},
			arguments: bson.D{{
				"$set",
				bson.D{
					{"tuition", mockStudentOne.Tuition},
					{"name", mockStudentOne.Name},
					{"grade", mockStudentOne.Grade},
					{"group", mockStudentOne.Group},
					{"career", mockStudentOne.Career},
				},
			}},
			expetedStudents: model.Students{mockStudentOne},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_ = ts.studentStorer.Insert(context.TODO(), ts.mockStudents)

			t.Cleanup(func() {
				_ = ts.studentStorer.Delete(context.TODO(), bson.M{})
			})

			err := ts.studentStorer.Update(context.TODO(), ts.filters, ts.arguments)
			if err != ts.expectedError {
				t.Errorf("expeted error %v, got error %v", ts.expectedError, err)
				t.SkipNow()
			}

			mock, _ := ts.studentStorer.Find(context.TODO(), bson.M{})

			if !reflect.DeepEqual(ts.expetedStudents, mock) {
				t.Errorf("expeted %v,\n got %v", ts.expetedStudents, mock)
			}
		})
	}
}
