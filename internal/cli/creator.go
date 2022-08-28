package cli

import (
	"context"
	"flag"

	"github.com/erik-sostenes/gofige/internal/services"
)

type (
	// Creator contains methods to create data in a collection
	Creator interface {
		// Create creates data
		Create(context.Context, string) error
	}
	creator struct {
		fs   *flag.FlagSet
		studentService services.StudentService
		path string
	}
)

// NewCreator cretes the subcommands and assigns the data to the variables of the flag
func NewCreator(studentService services.StudentService) *creator {
	c := &creator{
		fs: flag.NewFlagSet("insert", flag.ContinueOnError),
		studentService: studentService,
	}
	c.fs.StringVar(&c.path, "path", "", "address of the file from which the data is to be extracted")
	return c
}

func (c *creator) Name() string {
	return c.fs.Name()
}

func (c *creator) Parse(args []string) error {
	return c.fs.Parse(args)
}

func (c *creator) Run() error {
	return c.Create(context.TODO(), c.path)
}

func (c *creator) Create(ctx context.Context, path string) error {
	return c.studentService.Create(path)
}
