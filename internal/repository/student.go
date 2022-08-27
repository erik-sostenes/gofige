package repository

import (
	"context"

	"github.com/erik-sostenes/gofige/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// StudentStorer interface containing the methods to insert documents into mongodb
	StudentStorer interface {
		// Insert method that inserts one document a mongodb
		Insert(context.Context, model.Students) error
	}
	// studentStorer studentStorer will connect and work with the mongo client
	studentStorer struct {
		*mongo.Collection
	}
)

// NewStudentStorer returns an instance of the StudentStorer interface
func NewStudentStorer(db *mongo.Collection) StudentStorer {
	return &studentStorer{
		db,
	}
}

func (s *studentStorer) Insert(ctx context.Context, students model.Students) (err error) {
	var docs []interface{}

	for _, v := range students {
		docs = append(docs, v)
	}
	_, err = s.InsertMany(ctx, docs)
	return
}
