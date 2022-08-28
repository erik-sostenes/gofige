package cli

import (
	"context"
	"flag"

	"github.com/erik-sostenes/gofige/internal/model"
	"github.com/erik-sostenes/gofige/internal/services"
)

type (
	// Finder contains methods to search data in a collection
	Finder interface {
		// Find find all data and returns a .csv file
		Find(context.Context) (model.Students, error)
		// FindByFlags find any collection that matches the flags
		// returns a .csv file
		FindByFlags(context.Context, model.Student) (model.Students, error)
	}
	finder struct {
		fs             *flag.FlagSet
		studentService services.StudentService
		path           string
		all            bool
		studentFlags   model.Student
	}
)

func NewFinder(studentService services.StudentService) *finder {
	f := &finder{
		fs:             flag.NewFlagSet("find", flag.ContinueOnError),
		studentService: studentService,
	}
	f.fs.BoolVar(&f.all, "all", false, "find all the documents of the collection")

	f.fs.StringVar(&f.path, "path", "./", "addres when you want to save the .csv file with the data.")
	f.fs.StringVar(&f.studentFlags.Tuition, "tuition", "nil", "find the collection that matches with tuition")
	f.fs.StringVar(&f.studentFlags.Name, "name", "nil", "find the collection that matches with name")
	f.fs.StringVar(&f.studentFlags.Carrer, "carreer", "nil", "find the collection that matches with carreer")
	f.fs.StringVar(&f.studentFlags.Grade, "grade", "nil", "find the collection that matches with grade")
	f.fs.StringVar(&f.studentFlags.Group, "group", "nil", "find the collection that matches with group")
	return f
}

func (f *finder) Name() string {
	return f.fs.Name()
}

func (f *finder) Parse(args []string) error {
	return f.fs.Parse(args)
}

func (f *finder) Run() (err error) {
	if f.all {
		_, err = f.Find(context.TODO())
		return
	}
	_, err = f.FindByFlags(context.TODO(), f.studentFlags)
	return
}

func (f *finder) Find(ctx context.Context) (model.Students, error) {
	return f.studentService.Find(ctx)
}

func (f *finder) FindByFlags(ctx context.Context, studentFlags model.Student) (model.Students, error) {
	return f.studentService.FindByFlags(ctx, studentFlags)
}
