package metadata

import "fmt"

//Version represents a SemVer version for this package
type Version struct {
	Major            int
	Minor            int
	Patch            int
	ReleaseCandidate int
	Beta             bool
}

//VERSION is the current version for this package
//Update this with "yo gop:minor", "yo gop:minor" and "yo gop:patch".
//To enable beta "yo gop:beta" and RC number with "yo gop:rc 2".
var VERSION = &Version{
	Major:            0,
	Minor:            1,
	Patch:            0,
	ReleaseCandidate: 0,
	Beta:             false,
}

//GetVersion in string format of Major.Minor.Patch(-beta|-rc.1)?
func GetVersion(v ...*Version) string {
	var version *Version
	if v == nil || len(v) != 1 {
		version = VERSION
	} else {
		version = v[0]
	}

	rc := ""

	if version.Beta {
		rc = "-beta"
	} else {
		if version.ReleaseCandidate > 0 {
			rc = fmt.Sprintf("-rc.%d", version.ReleaseCandidate)
		}
	}
	return fmt.Sprintf("%d.%d.%d%s", version.Major, version.Minor, version.Patch, rc)
}
