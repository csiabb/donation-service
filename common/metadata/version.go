/*
Copyright Arxan Chain Ltd. 2020 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
