package commands

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/spf13/cobra"
	"sort"
)

type SortCommandI interface {
	Handle(args ...string) error
	Execute(cmd *cobra.Command, args []string) error
	Cmd() *cobra.Command
	Init()
}

func NewSortCommand() SortCommandI {
	cmd := &SortCommand{}
	cmd.Init()
	return cmd
}

type SortCommand struct {
	cmd *cobra.Command
}

func (s *SortCommand) Init() {
	s.cmd = &cobra.Command{
		Use:     "sort",
		Short:   "Sort the given versions",
		Long:    "Sort the given versions according to software phase and semantic version rules",
		Example: "semver sort 1.0.0 alpha 1.2.0-rc.2 1.4.0-beta.3 1.4.0-alpha.1",
		RunE:    s.Execute,
	}
}

func (s *SortCommand) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "sort",
		Short:   "Sort the given versions",
		Long:    "Sort the given versions according to software phase and semantic version rules",
		Example: "semver sort 1.0.0 alpha 1.2.0-rc.2 1.4.0-beta.3 1.4.0-alpha.1",
		RunE:    s.Execute,
	}
}

func (s *SortCommand) Execute(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 args: semver sort <version 1> <version 2> <version 3>")
	}

	return s.Handle(args...)
}

func (s *SortCommand)Handle(versionsStr ...string) error {
	var versions entities.Versions
	versions = make([]entities.Version, 0)
	for _, value := range versionsStr {
		version, err := entities.NewVersion(value)
		if err != nil {
			return err
		}
		versions = append(versions, *version)
	}

	sort.Sort(versions)
	s.printVersions(&versions)
	return nil
}

func (s *SortCommand) printVersions(versions *entities.Versions) {
	for _, value := range *versions {
		fmt.Println(value.String())
	}
}