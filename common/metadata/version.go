/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metadata

import (
	"fmt"
	"runtime"
)

// Version indicates the program version information
type Version struct {
	ProgramName string
	Release     uint8
	Fixpack     uint8
	Hotfix      uint8

	BuildNumber string
}

// ShortVersion returns a version string as 1.2.0
func (v *Version) ShortVersion() string {
	return fmt.Sprintf("%d.%d.%d", v.Release, v.Fixpack, v.Hotfix)
}

// FullVersion ...
func (v *Version) FullVersion() string {
	return fmt.Sprintf("%s:\n Version: %s\n BuildNumber: %s\n Go version: %s\n OS/Arch: %s",
		v.ProgramName, v.ShortVersion(), v.BuildNumber, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}

// GoVersion go version
func (v *Version) GoVersion() string {
	return fmt.Sprintf(runtime.Version())
}

// Platform Get Platform
func (v *Version) Platform() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

// DefaultVersion ...
func DefaultVersion(progName string, buildNumber string) *Version {
	return &Version{
		ProgramName: progName,
		BuildNumber: buildNumber,
		Release:     0,
		Fixpack:     1,
		Hotfix:      0,
	}
}
