package cli

import (
	"context"
	"flag"
	"fmt"
)

type (
	// Creator contains methods to create data in a collection
	Creator interface {
		// Create create data
		Create(context.Context)
	}
	creator struct {
		fs   *flag.FlagSet
		path string
	}
)

func NewCreator() *creator {
	c := &creator{
		fs: flag.NewFlagSet("insert", flag.ContinueOnError),
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
	fmt.Println("Creator", ", path: ", c.path)
	return nil
}
