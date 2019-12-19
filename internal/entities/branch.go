package entities

import (
	"fmt"
	"strings"
)

type Branch struct {
	name string
}

func (b *Branch) String() string {
	return b.name
}

func (b *Branch) Set(v string) error {
	if v == "master" || v == "develop" {
		b.name = v
		return nil
	}

	allowedPrefixes := []string{"feature/", "bugfix/", "release/", "hotfix/"}
	for _, i := range allowedPrefixes {
		if strings.HasPrefix(v, i) {
			b.name = v
			return nil
		}
	}

	return fmt.Errorf("unknown branch type: %s", v)
}

func (b *Branch) Type() string {
	return "branch"
}
