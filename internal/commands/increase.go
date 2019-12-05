package commands

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/controllers"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/spf13/cobra"
)

type IncreaseCommandI interface {
	Handle(versionStr string, phaseStr string) (*entities.Version, error)
	Execute(cmd *cobra.Command, args []string) error
	Cmd() *cobra.Command
	Init()
}

func NewIncreaseCommand() IncreaseCommandI {
	cmd := &IncreaseCommand{}
	cmd.Init()
	return cmd
}

type IncreaseCommand struct {
	cmd *cobra.Command
}

func (i *IncreaseCommand) Handle(versionStr string, phaseStr string) (*entities.Version, error) {
	version, err := entities.NewVersion(versionStr)
	if err != nil {
		return nil, err
	}

	v, e := controllers.IncreaseCommandController(version, phaseStr)
	return v, e
}

func (i *IncreaseCommand) Execute(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected 2 args: semver increase <version> <phase>")
	}

	newVersion, err := i.Handle(args[0], args[1])
	if err != nil {
		return err
	}

	fmt.Println(newVersion)
	return nil

}

func (i *IncreaseCommand) Cmd() *cobra.Command {
	return i.cmd
}

func (i *IncreaseCommand) Init() {
	i.cmd = &cobra.Command{
		Use:     "increase",
		Short:   "Increment given version based on provided software phase",
		Long:    "Increment the given version according to software phase and semantic version rules",
		Example: "semver increase 1.0.0 alpha",
		RunE:    i.Execute,
	}
}

func IncreaseVersionCommand(versionStr string, phaseStr string) (*entities.Version, error) {
	version, err := entities.NewVersion(versionStr)
	if err != nil {
		return nil, err
	}

	return controllers.IncreaseCommandController(version, phaseStr)
}
