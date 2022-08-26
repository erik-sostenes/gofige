package cli

import (
	"errors"
	"fmt"
	"os"
)

// Runner will define all subcommands and will initialize
type Runner interface {
	// Parse parses flag definitions from the argument list,
	// does not include the command name, only the subcommands
	Parse([]string) error
	// Run execute the process with the values of the subcommands  
	Run() error
	// Name returns the name of the subcommand to be worked with
	Name() string
}


// Execute the subcommand to be worked with
// will validate that the subcommand exists
func Execute(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{}

	subcommand := os.Args[1]	

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Parse(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}
