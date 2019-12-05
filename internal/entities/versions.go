package entities

import "github.com/maykonlf/semver-cli/internal/enum/phases"

type Versions []Version

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	if v[i].Major < v[j].Major {
		return true
	} else if v[i].Major > v[j].Major {
		return false
	}

	if v[i].Minor < v[j].Minor {
		return true
	} else if v[i].Minor > v[j].Minor {
		return false
	}

	if v[i].Patch < v[j].Patch {
		return true
	} else if v[i].Patch > v[j].Patch {
		return false
	}

	if phases.IndexOf(v[i].Phase) < phases.IndexOf(v[j].Phase) {
		return true
	} else if phases.IndexOf(v[i].Phase) < phases.IndexOf(v[j].Phase) {
		return false
	}

	if v[i].Phase == v[j].Phase && v[i].PatchNumber < v[j].PatchNumber {
		return true
	}

	return false
}
