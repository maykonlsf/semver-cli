package commands

import "strings"

type Command string

const (
	Increase Command = "increase"

	Unknown Command = ""
)

func (c Command) String() string {
	return string(c)
}

func Values() []Command {
	return []Command{
		Increase,
	}
}

func ValueOf(s string) Command {
	for _, value := range Values() {
		if strings.EqualFold(value.String(), s) {
			return value
		}
	}

	return Unknown
}