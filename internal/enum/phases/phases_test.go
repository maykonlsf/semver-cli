package phases_test

import (
	"github.com/maykonlf/semver-cli/internal/enum/phases"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValueOf(t *testing.T) {
	testCases := []struct {
		strValue string
		expected phases.Phase
	}{
		{strValue: "", expected: phases.Release},
		{strValue: "release", expected: phases.Release},
		{strValue: "alpha", expected: phases.Alpha},
		{strValue: "beta", expected: phases.Beta},
		{strValue: "rc", expected: phases.ReleaseCandidate},
		{strValue: "invalid", expected: phases.Unknown},
		{strValue: "unknown", expected: phases.Unknown},
	}

	Convey("Given a string form of the release phase", t, func() {
		Convey("Should identify the valid values and return the Phase value", func() {
			for _, testCase := range testCases {
				actual := phases.ValueOf(testCase.strValue)
				So(actual, ShouldEqual, testCase.expected)
			}
		})
	})
}

func TestIsRelease(t *testing.T) {
	Convey("Given a list of phases", t, func() {
		testCases := []struct {
			value    phases.Phase
			expected bool
		}{
			{value: phases.Release, expected: true},
			{value: phases.Alpha, expected: false},
			{value: phases.Beta, expected: false},
			{value: phases.ReleaseCandidate, expected: false},
			{value: phases.Unknown, expected: false},
		}
		Convey("Should only return true when is a phase.Release", func() {
			for _, testCase := range testCases {
				So(testCase.value.IsRelease(), ShouldEqual, testCase.expected)
			}
		})
	})
}
