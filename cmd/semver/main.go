package main

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/commands"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".semver-ci")
	_ = viper.ReadInConfig()

	var version entities.Version
	var branch entities.Branch

	rootCmd.PersistentFlags().Var(&version, "version", "current version")
	rootCmd.PersistentFlags().Var(&branch, "branch", "current branch name")
	_ = viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))
	_ = viper.BindPFlag("branch", rootCmd.PersistentFlags().Lookup("branch"))
	viper.SetDefault("version", "v0.0.0")
	viper.SetDefault("branch", "develop")

	rootCmd.AddCommand(commands.NewInitCommand().Cmd())
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
