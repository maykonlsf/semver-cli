package handlers

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/maykonlf/semver-cli/internal/enum/types"
)

func IncreaseCommandHandler(version *entities.Version, args ...string) (*entities.Version, error) {
	if len(args) < 1 {
		return nil, errors.New("insufficient args for increase operation. expected 1, given 0")
	}

	phase := types.ValueOf(args[0])
	if phase == types.Unknown {
		return nil, fmt.Errorf("unknown version increase type '%s'", args[0])
	}

	return version, nil
}

func IncreaseVersionAlpha(version *entities.Version) *entities.Version {
	if version.Phase == types.Alpha {
		version.PatchNumber += 1
		return version
	}

	version.Minor += 1
	version.Phase = types.Alpha
	version.PatchNumber = 1

	return version
}

func IncreaseVersionBeta(version *entities.Version) *entities.Version {
	if version.Phase == types.Beta {
		version.PatchNumber += 1
		return version
	}

	if version.Phase != types.Alpha {
		version.Minor += 1
	}

	version.Phase = types.Beta
	version.PatchNumber = 1
	return version
}

func IncreaseReleaseCandidate(version *entities.Version) *entities.Version {
	if version.Phase == types.Alpha || version.Phase == types.Beta || version.Phase == types.Release {
		version.PatchNumber = 1
	}

	if version.Phase == types.ReleaseCandidate {
		version.PatchNumber += 1
	}

	if version.Phase == types.Release {
		version.Minor += 1
	}

	version.Phase = types.ReleaseCandidate

	return version
}
