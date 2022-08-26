package cli

import (
	"context"
	"flag"
	"fmt"
)

type (
	// Finder contains methods to search data in a collection
	Finder interface {
		// Find find all data
		Find(context.Context)
	}
	finder struct {
		fs *flag.FlagSet
	}
)

func NewFinder() *finder {
	f := &finder{
		fs: flag.NewFlagSet("find", flag.ContinueOnError),
	}
	return f
}

func (f *finder) Name() string {
	return f.fs.Name()
}

func (f *finder) Parse(args []string) error {
	return f.fs.Parse(args)
}

func (f *finder) Run() error {
	fmt.Println("Finder")
	return nil
}
