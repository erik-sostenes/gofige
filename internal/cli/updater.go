package cli

import (
	"context"
	"flag"
	"fmt"
)

type (
	// Updater contains methods to update data in a collection
	Updater interface {
		// Update update data by an identifier
		Update(context.Context, string)
	}
	updater struct {
		fs *flag.FlagSet
	}
)

func NewUpdater() *updater {
	u := &updater{
		fs: flag.NewFlagSet("update", flag.ContinueOnError),
	}
	return u
}

func (u *updater) Name() string {
	return u.fs.Name()
}

func (u *updater) Parse(args []string) error {
	return u.fs.Parse(args)
}

func (u *updater) Run() error {
	fmt.Println("Update")
	return nil
}
