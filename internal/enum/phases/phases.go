package phases

import "strings"

type Phase string

func (tagType Phase) String() string {
	return string(tagType)
}

func (tagType Phase) IsRelease() bool {
	return  tagType == Release
}

const (
	Alpha            Phase = "alpha"
	Beta             Phase = "beta"
	ReleaseCandidate Phase = "rc"
	Release          Phase = ""

	Unknown Phase = "unknown"
)

func Values() []Phase {
	return []Phase{
		Alpha,
		Beta,
		ReleaseCandidate,
		Release,
	}
}

func ValueOf(value string) Phase {
	for _, valid := range Values() {
		if IsEqual(value, valid.String()) {
			return valid
		}
	}

	return Unknown
}

func IsEqual(value string, valid string) bool {
	return strings.EqualFold(value, valid)
}