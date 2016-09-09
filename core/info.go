package core

import (
	"fmt"
	"strings"
)

// The git commit that was compiled. This will be filled in by the compiler.
var (
	GitCommit   string
	GitDescribe string
)

// ProductName is the name of the product
const ProductName = "gochlog"

// Version is main version number that is being run at the moment.
const Version = "0.1.0"

// VersionPrerelease is a pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = "dev"

// GetVersionDisplay composes the parts of the version in a way that's suitable
// for displaying to humans.
func GetVersionDisplay() string {
	return fmt.Sprintf("%s version %s\n", ProductName, getHumanVersion())
}

func getHumanVersion() string {
	version := Version
	if GitDescribe != "" {
		version = GitDescribe
	}

	release := VersionPrerelease
	if GitDescribe == "" && release == "" {
		release = "dev"
	}
	if release != "" {
		version += fmt.Sprintf("-%s", release)
		if GitCommit != "" {
			version += fmt.Sprintf(" (%s)", GitCommit)
		}
	}

	// Strip off any single quotes added by the git information.
	return strings.Replace(version, "'", "", -1)
}
