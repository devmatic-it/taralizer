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
	"strings"
	"testing"
)

func TestNormalizeHeredocs(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "arbitrary_tag_replaced",
			src: `resource "google_compute_instance" "vm" {
  script = <<SCRIPT
#!/bin/bash
echo hello
SCRIPT
}`,
			want: `resource "google_compute_instance" "vm" {
  script = "__taralizer_heredoc__"



}`,
		},
		{
			name: "indented_tag_and_closing",
			src: `resource "google_compute_instance" "vm" {
  script = <<-EOT
    line one
    line two
  EOT
}`,
			want: `resource "google_compute_instance" "vm" {
  script = "__taralizer_heredoc__"



}`,
		},
		{
			name: "native_eof_tag_also_replaced",
			src: `locals {
  x = <<EOF
body
EOF
}`,
			want: `locals {
  x = "__taralizer_heredoc__"


}`,
		},
		{
			name: "multiple_heredocs",
			src: `resource "a" "b" {
  x = <<ONE
one body
ONE
  y = <<TWO
two body
TWO
}`,
			want: `resource "a" "b" {
  x = "__taralizer_heredoc__"


  y = "__taralizer_heredoc__"


}`,
		},
		{
			name: "arrow_in_string_not_replaced",
			src:  `locals { note = "use <<SCRIPT here" }`,
			want: `locals { note = "use <<SCRIPT here" }`,
		},
		{
			name: "arrow_in_line_comment_not_replaced",
			src: `# see <<SCRIPT docs
locals { a = 1 }`,
			want: `# see <<SCRIPT docs
locals { a = 1 }`,
		},
		{
			name: "arrow_in_block_comment_not_replaced",
			src: `/* <<SCRIPT */
locals { a = 1 }`,
			want: `/* <<SCRIPT */
locals { a = 1 }`,
		},
		{
			name: "unterminated_heredoc_returns_original",
			src: `resource "a" "b" {
  x = <<SCRIPT
body without closing line
}`,
			want: `resource "a" "b" {
  x = <<SCRIPT
body without closing line
}`,
		},
		{
			name: "arrow_with_space_not_a_heredoc",
			src:  `locals { a = 1 << 2 }`,
			want: `locals { a = 1 << 2 }`,
		},
		{
			name: "escaped_quote_in_string",
			src:  `locals { a = "say \"hi\" <<SCRIPT" }`,
			want: `locals { a = "say \"hi\" <<SCRIPT" }`,
		},
		{
			name: "heredoc_in_list",
			src:  "locals {\n  scripts = [<<A\nx\n" + "A\n,\n    <<B\ny\nB\n] }\n",
			want: "locals {\n  scripts = [\"__taralizer_heredoc__\"\n\n\n,\n    \"__taralizer_heredoc__\"\n\n\n] }\n",
		},
		{
			name: "no_heredocs_unchanged",
			src:  `resource "a" "b" { name = "n" }`,
			want: `resource "a" "b" { name = "n" }`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := normalizeHeredocs(tc.src); got != tc.want {
				t.Errorf("normalizeHeredocs() =\n%q\nwant:\n%q", got, tc.want)
			}
		})
	}
}

// TestNormalizeHeredocsPreservesLineCount checks that dropped heredoc bodies
// keep their newlines so later parse errors report accurate line numbers.
func TestNormalizeHeredocsPreservesLineCount(t *testing.T) {
	src := "resource \"a\" \"b\" {\n  x = <<SCRIPT\nl1\nl2\nl3\nSCRIPT\n}\n"
	got := normalizeHeredocs(src)

	if gotLines, wantLines := strings.Count(got, "\n"), strings.Count(src, "\n"); gotLines != wantLines {
		t.Errorf("newline count = %d, want %d:\n%s", gotLines, wantLines, got)
	}
}
