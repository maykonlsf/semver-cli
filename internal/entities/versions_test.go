package entities

import (
	"github.com/maykonlf/semver-cli/internal/enum/phases"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestVersions_Less(t *testing.T) {
	Convey("Given a versions list", t, func() {
		Convey("When is alpha versions list", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Alpha, PatchNumber: 1},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Alpha, PatchNumber: 2},
			}
			Convey("Then should compare both", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is an alpha and beta version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Alpha, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Beta, PatchNumber: 1},
			}
			Convey("Then alpha should be lesser than beta", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is an alpha and rc version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Alpha, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.ReleaseCandidate, PatchNumber: 1},
			}
			Convey("Then alpha should be lesser than rc", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is an alpha and release version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Alpha, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Release},
			}
			Convey("Then alpha should be lesser than release", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is beta versions list", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Beta, PatchNumber: 1},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Beta, PatchNumber: 2},
			}
			Convey("Then should compare both", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is beta and rc version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Beta, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.ReleaseCandidate, PatchNumber: 1},
			}
			Convey("Then should compare both", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is rc and rc version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.ReleaseCandidate, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.ReleaseCandidate, PatchNumber: 3},
			}
			Convey("Then should compare both", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When is rc and release version", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.ReleaseCandidate, PatchNumber: 2},
				Version{Major:1, Minor: 0, Patch: 0, Phase: phases.Release},
			}
			Convey("Then should compare both", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When When both is release with same major and minor", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 1, Phase: phases.Release},
				Version{Major:1, Minor: 0, Patch: 3, Phase: phases.Release},
			}
			Convey("Then should compare based on patch", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When both is release and minor differs", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 9, Phase: phases.Release},
				Version{Major:1, Minor: 2, Patch: 1, Phase: phases.Release},
			}
			Convey("Then should compare based on minor value", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})

		Convey("When both is release and major differs", func() {
			versions := Versions{
				Version{Major:1, Minor: 0, Patch: 9, Phase: phases.Release},
				Version{Major:2, Minor: 1, Patch: 0, Phase: phases.Release},
			}
			Convey("Then should compare based on major value", func() {
				So(versions.Less(0, 1), ShouldBeTrue)
				So(versions.Less(1, 0), ShouldBeFalse)
			})
		})
	})
}
