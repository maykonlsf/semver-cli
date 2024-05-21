package entities

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewVersion(t *testing.T) {
	Convey("Given an invalid version string format", t, func() {
		versionStr := "dumb"
		Convey("When tries to instantiate a new Version", func() {
			Convey("Then should return an error message", func() {
				version, err := NewVersion(versionStr)
				So(version, ShouldBeNil)
				So(err, ShouldBeError)
			})
		})
	})

	Convey("Given an valid version string format", t, func() {
		testCases := []string{
			"1.2.31-alpha.1",
			"1.2.32-beta.1",
			"1.2.33-rc.1",
			"1.2.34",
		}

		Convey("Should not return an error", func() {
			for _, versionStr := range testCases {
				version, err := NewVersion(versionStr)
				So(version, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(version.Type(), ShouldEqual, "version")
			}
		})
	})
}
