package commands

import (
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"sort"
)

func SortCommand(versionsStr ...string) error {
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
	printVersions(&versions)
	return nil
}

func printVersions(versions *entities.Versions) {
	for _, value := range *versions {
		fmt.Println(value.String())
	}
}