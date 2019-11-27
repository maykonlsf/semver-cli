package handlers

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/entities"
	"github.com/maykonlf/semver-cli/internal/enum/types"
)

func IncreaseCommandHandler(version *entities.Version, args ...string) error {
	if len(args) < 1 {
		return errors.New("insufficient args for increase operation. expected 1, given 0")
	}

	phase := types.ValueOf(args[0])
	if phase == types.Unknown {
		return fmt.Errorf("unknown version increase type '%s'", args[0])
	}


}

func increase(version *entities.Version, phase types.Phase) *entities.Version {

}