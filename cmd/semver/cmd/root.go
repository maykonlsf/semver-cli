package cmd

import (
	"fmt"
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
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(increaseCmd)
	rootCmd.AddCommand(sortCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
