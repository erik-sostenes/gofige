package services

import (
	"encoding/json"

	"github.com/erik-sostenes/gofige/internal/model"
)

// flags contains the flags
type flags map[string]string

// UnmarshalFlags returns a new map with the flgas required
// example [name:erik, group:nil, tuition: XFB649J] -> [name:erik, tuition: XFB649J]
func (f *flags) UnmarshalFlags(flags model.Student) (flags, error) {
	bytes, err := json.Marshal(flags)

	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return *f, err
	}

	for k, v := range *f {
		if v == "nil" {
			delete(*f, k)
		}
	}
	return *f, err

}
