package commands

import (
	"github.com/maykonlf/semver-cli/internal/controllers"
	"github.com/maykonlf/semver-cli/internal/entities"
)

func IncreaseVersionCommand(versionStr string, phaseStr string) (*entities.Version, error) {
	version, err := entities.NewVersion(versionStr)
	if err != nil {
		return nil, err
	}

	return controllers.IncreaseCommandController(version, phaseStr)
}
