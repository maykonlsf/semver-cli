package controllers

import (
	"github.com/maykonlf/semver-cli/internal/entities"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIncreaseVersionAlpha(t *testing.T) {
	Convey("Given an previous version and increase to next alpha", t, func() {
		Convey("When previous version is already and alpha", func() {
			version, _ := entities.NewVersion("1.0.1-alpha.1")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.0.1-alpha.2")
			})
		})

		Convey("When previous version is already and alpha with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-alpha.1")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-alpha.2")
			})
		})

		Convey("When previous version is an beta", func() {
			version, _ := entities.NewVersion("1.0.0-beta.3")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})

		Convey("When previous version is an beta with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-beta.3")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})

		Convey("When previous version is an release candidate", func() {
			version, _ := entities.NewVersion("1.0.0-rc.1")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})

		Convey("When previous version is an release candidate with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-rc.1")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})

		Convey("When previous version is an release", func() {
			version, _ := entities.NewVersion("1.0.1")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})

		Convey("When previous version is an release with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0")
			Convey("Then should increase patch count version", func() {
				newVersion := IncreaseVersionAlpha(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-alpha.1")
			})
		})
	})
}

func TestIncreaseVersionBeta(t *testing.T) {
	Convey("Given an previous version and increase to next beta", t, func() {
		Convey("When previous version is an alpha", func() {
			version, _ := entities.NewVersion("1.0.0-alpha.7")
			Convey("Then should promote to beta and set patch to number 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-beta.1")
			})
		})

		Convey("When previous version is an alpha with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-alpha.7")
			Convey("Then should promote to beta and set patch to number 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-beta.1")
			})
		})

		Convey("When previous version is an beta", func() {
			version, _ := entities.NewVersion("1.0.0-beta.2")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-beta.3")
			})
		})

		Convey("When previous version is an beta with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-beta.2")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-beta.3")
			})
		})

		Convey("When previous version is an release candidate", func() {
			version, _ := entities.NewVersion("1.0.0-rc.2")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-beta.1")
			})
		})

		Convey("When previous version is an release candidate with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-rc.2")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-beta.1")
			})
		})

		Convey("When previous version is an final release", func() {
			version, _ := entities.NewVersion("1.0.0")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-beta.1")
			})
		})

		Convey("When previous version is an final release with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0")
			Convey("Then should increase minor and set patch number to 1", func() {
				newVersion := IncreaseVersionBeta(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-beta.1")
			})
		})
	})
}

func TestIncreaseReleaseCandidate(t *testing.T) {
	Convey("Given an previous version and increase to next release Candidate", t, func() {
		Convey("When previous version is alpha", func() {
			version, _ := entities.NewVersion("1.0.0-alpha.3")
			Convey("Then should promote patch to release candidate and set patch number to 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.1")
			})
		})

		Convey("When previous version is alpha with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-alpha.3")
			Convey("Then should promote patch to release candidate and set patch number to 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.1")
			})
		})

		Convey("When previous version is beta", func() {
			version, _ := entities.NewVersion("1.0.0-beta.7")
			Convey("Then should promote patch to release candidate and set patch number to 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.1")
			})
		})

		Convey("When previous version is beta with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-beta.7")
			Convey("Then should promote patch to release candidate and set patch number to 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.1")
			})
		})

		Convey("When previous version is a release candidate", func() {
			version, _ := entities.NewVersion("1.0.0-rc.7")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.8")
			})
		})

		Convey("When previous version is a release candidate with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-rc.7")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0-rc.8")
			})
		})

		Convey("When previous version is a final release", func() {
			version, _ := entities.NewVersion("1.0.0")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-rc.1")
			})
		})

		Convey("When previous version is a final release with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0")
			Convey("Then should increase patch number in 1", func() {
				newVersion := IncreaseReleaseCandidate(version)
				So(newVersion.String(), ShouldEqual, "v1.1.0-rc.1")
			})
		})
	})
}

func TestIncreaseRelease(t *testing.T) {
	Convey("Given an previous version and increase to the next release version", t, func() {
		Convey("When previous version is alpha", func() {
			version, _ := entities.NewVersion("1.0.0-alpha.3")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is alpha with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-alpha.3")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is beta", func() {
			version, _ := entities.NewVersion("1.0.0-beta.2")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is beta with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-beta.2")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is a release candidate", func() {
			version, _ := entities.NewVersion("1.0.0-rc.2")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is a release candidate with prefix", func() {
			version, _ := entities.NewVersion("v1.0.0-rc.2")
			Convey("Then should promote version to final release", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.0")
			})
		})

		Convey("When previous version is already a final release", func() {
			version, _ := entities.NewVersion("1.0.3")
			Convey("Then should increase patch version", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.4")
			})
		})

		Convey("When previous version is already a final release with prefix", func() {
			version, _ := entities.NewVersion("v1.0.3")
			Convey("Then should increase patch version", func() {
				newVersion := IncreaseRelease(version)
				So(newVersion.String(), ShouldEqual, "v1.0.4")
			})
		})
	})
}
