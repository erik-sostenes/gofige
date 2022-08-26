package cli

import (
	"context"
	"flag"
	"fmt"
)

type (
	// Deleter contains methods to delete data in a collection
	Deleter interface {
		// Delete delete data by an identifier
		Delete(context.Context)
	}
	deleter struct {
		fs *flag.FlagSet
	}
)

func NewDeleter() *deleter {
	d := &deleter{
		fs: flag.NewFlagSet("delete", flag.ContinueOnError),
	}
	return d
}

func (d *deleter) Name() string {
	return d.fs.Name()
}

func (d *deleter) Parse(args []string) error {
	return d.fs.Parse(args)
}

func (d *deleter) Run() error {
	fmt.Println("Delete")
	return nil
}
