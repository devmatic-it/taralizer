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

package terraform

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// TestParserCorpus runs the ANTLR import parser (including heredoc
// normalization) over every fixture in testdata/corpus. Each fixture must
// parse without syntax errors; fixtures with a <name>.tf.expected sidecar
// must yield exactly the listed resource addresses (type.name, one per line,
// in file order). Fixtures without a sidecar must contain no resource blocks.
func TestParserCorpus(t *testing.T) {
	matches, err := filepath.Glob(filepath.Join("testdata", "corpus", "*.tf"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) == 0 {
		t.Fatal("no corpus fixtures found in testdata/corpus")
	}

	for _, path := range matches {
		t.Run(filepath.Base(path), func(t *testing.T) {
			refs, err := parseResourceBlocks(path)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}

			var got []string
			for _, r := range refs {
				got = append(got, r.typ+"."+r.name)
			}

			wantPath := path + ".expected"
			if _, statErr := os.Stat(wantPath); os.IsNotExist(statErr) {
				if len(got) != 0 {
					t.Fatalf("expected no resource blocks, got %v", got)
				}
				return
			}

			wantRaw, err := os.ReadFile(wantPath)
			if err != nil {
				t.Fatalf("read %s: %v", wantPath, err)
			}
			var want []string
			for _, line := range strings.Split(strings.TrimSpace(string(wantRaw)), "\n") {
				if line != "" {
					want = append(want, line)
				}
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("resource addresses =\n  %v\nwant:\n  %v", got, want)
			}
		})
	}
}
