package main

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/commands"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "semver",
		Short: "Semantic version cli",
		Long:  "Semantic version tool helper to validate and increase versions semantically",
	}
)

func init() {
	rootCmd.AddCommand(commands.NewValidateCommand().Cmd())
	rootCmd.AddCommand(commands.NewIncreaseCommand().Cmd())
	rootCmd.AddCommand(commands.NewSortCommand().Cmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
