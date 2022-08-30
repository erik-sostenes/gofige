package repository

import (
	"context"

	"github.com/erik-sostenes/gofige/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// StudentStorer interface containing methods for working with documents in mongodb
	StudentStorer interface {
		// Insert method that inserts one document with collections to mongodb
		Insert(context.Context, model.Students) error
		// Find method that find one document with collections to mongodb
		Find(context.Context, map[string]string) (model.Students, error)
		// Delete
		Delete(context.Context, map[string]string) error
	}
	// studentStorer studentStorer will connect and work with the mongo client
	studentStorer struct {
		db *mongo.Collection
	}
)

// NewStudentStorer returns an instance of the StudentStorer interface
func NewStudentStorer(db *mongo.Collection) StudentStorer {
	return &studentStorer{
		db,
	}
}

// Insert inserts ine document with collections of students
func (s *studentStorer) Insert(ctx context.Context, students model.Students) (err error) {
	var docs []interface{}

	for _, v := range students {
		docs = append(docs, v)
	}
	_, err = s.db.InsertMany(ctx, docs)
	return
}

// Find gets a document depending on the filter conditions
// the filter will be through the flags
func (s *studentStorer) Find(ctx context.Context, flags map[string]string) (students model.Students, err error) {
	cur, err := s.db.Find(ctx, flags)
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var student model.Student
		if err = cur.Decode(&student); err != nil {
			return
		}
		students = append(students, student)
	}
	if err = cur.Err(); err != nil {
		return
	}
	if len(students) == 0 {
		err = mongo.ErrNoDocuments
		return
	}
	return
}

func (s *studentStorer) Delete(ctx context.Context, _ map[string]string) (err error) {
	_, err = s.db.DeleteMany(ctx, bson.M{})
	return
}

// _ "implement" constraint for *MockStudentStorer
var _ StudentStorer = (*MockStudentStorer)(nil)

//MockStudentStorer store od model.Students
type MockStudentStorer struct {
	MockStudents model.Students
	Error        error
}

// Insert save a model.Students
func (m *MockStudentStorer) Insert(ctx context.Context, students model.Students) (err error) {
	m.MockStudents = append(m.MockStudents, students...)
	return m.Error
}

// Find search a model.Students by arguments
func (m *MockStudentStorer) Find(context.Context, map[string]string) (model.Students, error) {
	return m.MockStudents, m.Error
}

// Delete delete a model.Students by arguments
func (m *MockStudentStorer) Delete(context.Context, map[string]string) (err error) {
	return
}
