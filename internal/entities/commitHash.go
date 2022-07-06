package entities

import (
	"errors"
	"regexp"
)

type CommitHash struct {
	Value string
}

func (e *CommitHash) String() string {
	return e.Value
}

func (e *CommitHash) Set(value string) error {
	commitHashPattern, _ := regexp.Compile(`\b[0-9a-f]{5,40}\b`)
	if !commitHashPattern.MatchString(value) {
		return errors.New("invalid commit-hash format")
	}

	e.Value = value
	return nil
}

func (e *CommitHash) Type() string {
	return ("CommitHash")
}
