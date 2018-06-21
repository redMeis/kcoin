package params

import (
	"fmt"
)

var (
	VersionMajor = 1          // Major version component of the current release
	VersionMinor = 0          // Minor version component of the current release
	VersionPatch = 0          // Patch version component of the current release
	VersionMeta  = "unstable" // Version metadata to append to the version string

	Version           = "" // Version holds the textual version string.
	VersionWithCommit = "" // Version holds the textual version string with commit.

	BuildTime = ""
	Commit    = ""
)

func init() {
	setVersion()
}

func SetGitTagVersion(tag string) {
	if tag == "" {
		return
	}
	_, err := fmt.Sscanf(tag, "%d.%d.%d", &VersionMajor, &VersionMinor, &VersionPatch)
	if err != nil {
		panic(err)
	}
	VersionMeta = "stable" // Assume git-tagged versions are stable
	setVersion()
}

func SetBuildTime(buildTime string) {
	BuildTime = buildTime
}

func SetCommit(commit string) {
	Commit = commit
	setVersion()
}

func setVersion() {
	v := fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	Version = v
	if len(Commit) >= 8 {
		VersionWithCommit = fmt.Sprintf("%v-%v", Version, Commit)
	} else {
		VersionWithCommit = Version
	}
}