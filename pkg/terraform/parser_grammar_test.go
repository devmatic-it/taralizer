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
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// syntaxErrorCounter counts parse errors reported through the ANTLR error
// listener channel.
type syntaxErrorCounter struct {
	*antlr.DefaultErrorListener
	count int
}

func (l *syntaxErrorCounter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.count++
}

// ruleCountingListener counts the rule contexts visited during a tree walk.
type ruleCountingListener struct {
	*antlr.BaseParseTreeListener
	rules int
}

func (l *ruleCountingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.rules++
}

// newParserFor builds an ANTLR Terraform parser over src.
func newParserFor(src string) *TerraformParser {
	input := antlr.NewInputStream(src)
	lexer := NewTerraformLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return NewTerraformParser(stream)
}

// parseSource parses src as a full Terraform file, fails the test on any
// syntax error, and walks the tree to confirm rule contexts were produced.
func parseSource(t *testing.T, src string) *File_Context {
	t.Helper()
	parser := newParserFor(src)
	counter := &syntaxErrorCounter{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	parser.AddErrorListener(counter)
	tree := parser.File_()
	if counter.count != 0 {
		t.Fatalf("expected no syntax errors, got %d for:\n%s", counter.count, src)
	}
	fileCtx, ok := tree.(*File_Context)
	if !ok {
		t.Fatalf("expected *File_Context, got %T", tree)
	}
	walker := &ruleCountingListener{}
	antlr.ParseTreeWalkerDefault.Walk(walker, tree)
	if walker.rules == 0 {
		t.Fatal("listener visited no rule contexts")
	}
	return fileCtx
}

// TestParserGrammarCoverage feeds the ANTLR parser one case per grammar rule
// or alternative so every production in Terraform.g4 is exercised. Each case
// must parse with zero syntax errors.
func TestParserGrammarCoverage(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{
			name: "all_filedecl_types",
			src: `
locals { env = "prod" }
module "deploy" { source = "./modules/deploy" }
output "ip" { value = 1.5 }
provider "google" { region = "us-central1" }
variable "name" { default = "x" }
data "google_compute_network" "lookup" { name = "n" }
resource "google_compute_instance" "vm" { name = "i" }
terraform { required_version = ">= 1.0" }
`,
		},
		{
			name: "comments",
			src: `# line comment
// slash comment
/* block
comment */
resource "google_compute_instance" "vm" { name = "i" } // trailing
`,
		},
		{
			name: "nested_blocks_with_labels",
			src: `resource "google_container_cluster" "c" {
  name = "c"
  node_pool { size = 1 }
  network_config {
    subnetwork "a" { name = "s1" }
    subnetwork "b" "extra" { name = "s2" }
  }
}
`,
		},
		{
			name: "all_value_types",
			src: `resource "google_compute_instance" "vm" {
  a = nul
  b = +42
  c = -7
  d = 3.14
  e = "hello"
  f = <<-EOF
multiline string
EOF
  g = true
  h = false
  i = var.region
  j = local.env
  k = data.google_compute_network.n.name
  l = module.m.source
  m = plain_id
  n = <<DESCRIPTION
some description text
DESCRIPTION
  o = file("config.json")
  p = length([1, 2])
  q = concat("a", "b")
  r = jsonencode({ a = 1 })
  s = <<EOF
eof heredoc
EOF
}
`,
		},
		{
			name: "lists_and_maps",
			src: `resource "google_compute_instance" "vm" {
  empty_list = []
  one = [1]
  many = ["a", "b", "c"]
  trailing = [1, 2, 3,]
  nested = [[1], ["x"]]
  empty_map = {}
  flat_map = { a = 1, b = "two" }
  no_comma = { c = 3 d = 4 }
}
`,
		},
		{
			name: "identifier_chains",
			src: `resource "google_compute_instance" "vm" {
  a = foo
  b = foo.bar.baz
  c = list[0]
  d = list[1].attr
  e = *
  f = *.name
  g = provider.google.project
  h = variable.name
  i = in.list
  j = 2.value
}
`,
		},
		{
			name: "all_operators",
			src: `resource "google_compute_instance" "vm" {
  a = 1 + 2
  b = 3 - 4
  c = 6 / 2
  d = 3 * 4
  e = 7 % 3
  f = 1 > 2
  g = 1 >= 2
  h = 1 < 2
  i = 1 <= 2
  j = 1 == 2
  k = 1 != 2
  l = true && false
  m = true || false
}
`,
		},
		{
			name: "parens_ternary_forloop",
			src: `resource "google_compute_instance" "vm" {
  a = (1 + 2) * 3
  b = true ? "yes" : "no"
  c = for x in [1, 2] : x * 2
}
`,
		},
		{
			name: "empty_bodies",
			src: `resource "google_compute_instance" "vm" {}
terraform {}
provider "google" {}
`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			parseSource(t, tc.src)
		})
	}
}

// TestParseFileTreeShape verifies that a file containing one data, one
// resource and one terraform block parses into the expected tree structure.
func TestParseFileTreeShape(t *testing.T) {
	src := `data "google_x" "d" {} resource "google_y" "r" {} terraform { a = 1 }`
	fileCtx := parseSource(t, src)

	if got := len(fileCtx.AllData()); got != 1 {
		t.Errorf("AllData() = %d entries, want 1", got)
	} else if fileCtx.Data(0) == nil {
		t.Error("Data(0) = nil, want context")
	}

	if got := len(fileCtx.AllResource()); got != 1 {
		t.Errorf("AllResource() = %d entries, want 1", got)
	} else if fileCtx.Resource(0) == nil {
		t.Error("Resource(0) = nil, want context")
	}

	if got := len(fileCtx.AllTerraform()); got != 1 {
		t.Errorf("AllTerraform() = %d entries, want 1", got)
	} else if fileCtx.Terraform(0) == nil {
		t.Error("Terraform(0) = nil, want context")
	}
}
