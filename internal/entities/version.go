package entities

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/enum/phases"
	"github.com/maykonlf/semver-cli/internal/utils/str"
	"regexp"
)

type Version struct {
	Prefix      string
	Major       uint
	Minor       uint
	Patch       uint
	Phase       phases.Phase
	PatchNumber uint
}

func NewVersion(v string) (*Version, error) {
	tagPattern, _ := regexp.Compile(`^(v?)(\d+)\.(\d+)\.(\d+)((-(alpha|beta|rc))\.(\d+))?$`)
	if !tagPattern.MatchString(v) {
		return nil, errors.New("invalid version format")
	}

	parts := tagPattern.FindStringSubmatch(v)
	version := &Version{
		Prefix:      parts[1],
		Major:       str.ParseUIntOrDefault(parts[2]),
		Minor:       str.ParseUIntOrDefault(parts[3]),
		Patch:       str.ParseUIntOrDefault(parts[4]),
		Phase:       phases.ValueOf(parts[7]),
		PatchNumber: str.ParseUIntOrDefault(parts[8]),
	}
	return version, nil
}

func (v *Version) String() string {
	if v.Phase.IsRelease() {
		return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	}
	return fmt.Sprintf("v%d.%d.%d-%s.%d", v.Major, v.Minor, v.Patch, v.Phase, v.PatchNumber)
}
