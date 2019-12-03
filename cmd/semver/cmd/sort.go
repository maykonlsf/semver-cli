package cmd

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/commands"
	"github.com/spf13/cobra"
)

var (
	sortCmd = &cobra.Command{
		Use:     "sort",
		Short:   "Sort the given versions",
		Long:    "Sort the given versions according to software phase and semantic version rules",
		Example: "semver sort 1.0.0 alpha 1.2.0-rc.2 1.4.0-beta.3 1.4.0-alpha.1",
		RunE:    sort,
	}
)

func sort(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args: semver sort <version 1> <version 2> <version 3>")
	}

	err := commands.SortCommand(args...)
	return err
}
