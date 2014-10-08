package helpers

import (
	"errors"

	"github.com/codegangsta/cli"
)

type CliCommands []cli.Command

func (commands CliCommands) Find(name string) (cli.Command, error) {
	for _, v := range commands {
		if v.Name == name {
			return v, nil
		}
	}
	return cli.Command{}, errors.New("Command not found")
}
