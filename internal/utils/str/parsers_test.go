package str_test

import (
	"github.com/maykonlf/semver-cli/internal/utils/str"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParseUIntOrDefault(t *testing.T) {
	Convey("Given a set of numeric strings", t, func() {
		Convey("It should parte and return the int value", func() {
			testCases := []struct {
				input    string
				expected int
			}{
				{input: "0", expected: 0},
				{input: "1", expected: 1},
				{input: "345", expected: 345},
			}

			for _, testCase := range testCases {
				So(str.ParseUIntOrDefault(testCase.input), ShouldEqual, testCase.expected)
			}
		})
	})

	Convey("Given an non numeric string", t, func() {
		Convey("It should return zer", func() {
			testCases := []string{
				"AAA",
				"",
				"1.2",
				"-%63",
				"$#@-",
			}

			for _, testCase := range testCases {
				So(str.ParseUIntOrDefault(testCase), ShouldEqual, 0)
			}
		})
	})
}
