package model

// WrongFile will return an error when something unforeseen happens in the file
type WrongFile string 

// WrongFile implements the Error interface
func (s WrongFile) Error() string {
	return string(s)
}
