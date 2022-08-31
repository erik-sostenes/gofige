package model

type (
	// Student domain object
	Student struct {
		Name    string `json:"name" bson:"name"`
		Tuition string `json:"tuition" bson:"tuition"`
		Grade   string `json:"grade" bson:"grade"`
		Group   string `json:"group" bson:"group"`
		Career  string `json:"carreer" bson:"career"`
	}
	// Students slice of students
	Students []Student
)
