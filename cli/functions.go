package cli

import (
	"errors"
	"fmt"
	"os"
)

func Execute(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	commands := []Runner{
		// use here cli commands
		MakeServerRunCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() != subcommand {
			continue
		}

		if err := cmd.Init(os.Args[2:]); err != nil {
			return err
		}
		return cmd.Run()
	}

	return errors.New(fmt.Sprintf("Unknown subcommand: %s", subcommand))
}
