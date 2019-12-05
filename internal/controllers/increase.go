package controllers

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/maykonlf/semver-cli/internal/enum/phases"
)

var (
	increaseCommandMap = map[string]func(version *entities.Version) *entities.Version{
		"alpha":   IncreaseVersionAlpha,
		"beta":    IncreaseVersionBeta,
		"rc":      IncreaseReleaseCandidate,
		"release": IncreaseRelease,
	}
)

func IncreaseCommandController(version *entities.Version, args ...string) (*entities.Version, error) {
	if len(args) < 1 {
		return nil, errors.New("insufficient args for increase operation. expected 1, given 0")
	}

	increaseFunction := increaseCommandMap[args[0]]
	if increaseFunction == nil {
		return nil, fmt.Errorf("unknown version increase type '%s'", args[0])
	}
	return increaseFunction(version), nil
}

func IncreaseVersionAlpha(version *entities.Version) *entities.Version {
	if version.Phase == phases.Alpha {
		version.PatchNumber++
		return version
	}

	version.Minor++
	version.Phase = phases.Alpha
	version.PatchNumber = 1

	return version
}

func IncreaseVersionBeta(version *entities.Version) *entities.Version {
	if version.Phase == phases.Beta {
		version.PatchNumber++
		return version
	}

	if version.Phase != phases.Alpha {
		version.Minor++
	}

	version.Phase = phases.Beta
	version.PatchNumber = 1
	return version
}

func IncreaseReleaseCandidate(version *entities.Version) *entities.Version {
	if version.Phase == phases.Alpha || version.Phase == phases.Beta || version.Phase == phases.Release {
		version.PatchNumber = 1
	}

	if version.Phase == phases.ReleaseCandidate {
		version.PatchNumber++
	}

	if version.Phase == phases.Release {
		version.Minor++
	}

	version.Phase = phases.ReleaseCandidate

	return version
}

func IncreaseRelease(version *entities.Version) *entities.Version {
	if version.Phase == phases.Release {
		version.Patch++
	}

	version.Phase = phases.Release
	return version
}
