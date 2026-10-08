// Package version provides version information for Taralizer
// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package version

// Info contains version information
type Info struct {
	Version string
	Commit  string
	Date    string
	BuiltBy string
}

// Get returns version information
func Get() Info {
	return Info{
		Version: Version,
		Commit:  Commit,
		Date:    Date,
		BuiltBy: BuiltBy,
	}
}

// Version is the application version injected by the build system
var Version = "dev"

// Commit is the git commit hash injected by the build system
var Commit = "none"

// Date is the build date injected by the build system
var Date = "unknown"

// BuiltBy is who built the binary injected by the build system
var BuiltBy = "unknown"
