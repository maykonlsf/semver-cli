package commands

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type InitCommandI interface {
	Handle() error
	Execute(cmd *cobra.Command, args []string) error
	Cmd() *cobra.Command
	Init()
}

func NewInitCommand() InitCommandI {
	cmd := &InitCommand{}
	cmd.Init()
	return cmd
}

type InitCommand struct {
	version *entities.Version
	branch  *entities.Branch
	cmd     *cobra.Command
}

func (i *InitCommand) Cmd() *cobra.Command {
	return i.cmd
}

func (i *InitCommand) Execute(cmd *cobra.Command, args []string) error {
	err := i.Handle()
	if err == nil {
		fmt.Println("version:", viper.GetString("version"))
		fmt.Println("branch:", viper.GetString("branch"))
	}
	return err
}

func (i *InitCommand) Handle() error {
	return viper.SafeWriteConfig()
}

func (i *InitCommand) Init() {
	i.cmd = &cobra.Command{
		Use:     "init",
		Short:   "Initialize semver config",
		Long:    "Initialize required semver config file",
		Example: "semver init [--version v0.1.0] [--branch develop]",
		RunE:    i.Execute,
	}
}
