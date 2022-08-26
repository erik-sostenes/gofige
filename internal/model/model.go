package model

type (
	// Student domain object
	Student struct {
		Name    string `json:"name"`
		Tuition string `json:"tuition"`
		Grade   string `json:"grade"`
		Group   string `json:"group"`
		Carrer  string `json:"carreer"`
	}
	// Students slice of students
	Students []Student
)
