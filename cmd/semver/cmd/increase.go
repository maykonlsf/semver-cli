package cmd

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/commands"
	"github.com/spf13/cobra"
)

var (
	increaseCmd = &cobra.Command{
		Use:     "increase",
		Short:   "Increment given version based on provided software phase",
		Long:    "Increment the given version according to software phase and semantic version rules",
		Example: "semver increase 1.0.0 alpha",
		RunE:    increase,
	}
)

func increase(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected 2 args: semver increase <version> <phase>")
	}

	newVersion, err := commands.IncreaseVersionCommand(args[0], args[1])
	if err != nil {
		return err
	}

	fmt.Println(newVersion)
	return nil
}
