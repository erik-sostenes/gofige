package services

import (
	"encoding/json"

	"github.com/erik-sostenes/gofige/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// unmarshalFlags returns a new map with the flgas required
// example [name:erik, group:nil, tuition: XFB649J] -> [name:erik, tuition: XFB649J]
func unmarshalFlags(flags model.Student) (bson.M, error) {
	b := make(bson.M)
	bytes, err := json.Marshal(flags)

	err = json.Unmarshal(bytes, &b)
	if err != nil {
		return b, err
	}

	for k, v := range b {
		if v == "nil" {
			delete(b, k)
		}
	}
	return b, err
}
