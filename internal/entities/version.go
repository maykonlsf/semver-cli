package entities

import (
	"errors"
	"fmt"
	"github.com/maykonlf/semver-cli/internal/enum/types"
	"github.com/maykonlf/semver-cli/internal/utils/str"
	"regexp"
)

type Version struct {
	major       uint
	minor       uint
	patch       uint
	versionType types.Phase
	patchNumber uint
}

func NewVersion(v string) (*Version, error) {
	tagPattern, _ := regexp.Compile(`^(v?)(\d+)\.(\d+)\.(\d+)((-(alpha|beta|rc))\.(\d+))?$`)
	if !tagPattern.MatchString(v) {
		return nil, errors.New("invalid version format")
	}

	parts := tagPattern.FindStringSubmatch(v)
	version := &Version{}
	version.Major(str.ParseUIntOrDefault(parts[2])).Minor(str.ParseUIntOrDefault(parts[3])).PatchVersion(str.ParseUIntOrDefault(parts[4])).Phase(types.ValueOf(parts[7])).PatchNumber(str.ParseUIntOrDefault(parts[8]))
	return version, nil
}

func (v *Version) Major(n uint) *Version {
	v.major = n
	return v
}

func (v *Version) Minor(n uint) *Version {
	v.minor = n
	return v
}

func (v *Version) PatchVersion(patch uint) *Version {
	v.patch = patch
	return v
}

func (v *Version) Phase(p types.Phase) *Version {
	v.versionType = p
	return v
}

func (v *Version) PatchNumber(n uint) *Version {
	v.patchNumber = n
	return v
}

func (v *Version) String() string {
	if v.versionType.IsRelease() {
		return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
	} else {
		return fmt.Sprintf("%d.%d.%d-%s.%d", v.major, v.minor, v.patch, v.versionType, v.patchNumber)
	}
}
