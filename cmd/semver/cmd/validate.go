package cmd

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/spf13/cobra"
)

var (
	validateCmd = &cobra.Command{
		Use:     "validate",
		Short:   "Validate given version",
		Long:    "Validate the given version according to semantic version rules",
		Example: "semver validate 1.0.0",
		RunE:    validate,
	}
)

func validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("needs version arg")
	}

	version, err := entities.NewVersion(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("valid version: %s\n", version.String())
	return nil
}
