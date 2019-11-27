package controllers

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/maykonlf/semver-cli/internal/enum/commands"
)

type CommandController struct {
	commandMap map[commands.Command] func(version *entities.Version, args ...string) error
}

func (c *CommandController) DispatchCommand(version *entities.Version, commandStr string, args ...string) error {
	command := commands.ValueOf(commandStr)
	if command == commands.Unknown {
		return fmt.Errorf("unknown command '%s'", commandStr)
	}

	return c.commandMap[command](version, args...)
}
