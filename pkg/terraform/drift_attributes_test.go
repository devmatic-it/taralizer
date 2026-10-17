// Copyright 2021 taralizer authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// parseBody writes src (a single resource block) to a temp .tf file, parses it
// and returns the first resource's body context. It is the shared fixture for
// the attribute-extraction tests.
func parseBody(t *testing.T, src string) *BlockbodyContext {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "main.tf")
	require.NoError(t, os.WriteFile(path, []byte(src), 0o644))

	refs, err := parseResourceBlocks(path)
	require.NoError(t, err)
	require.Len(t, refs, 1)
	return refs[0].body
}

const vmHeader = "resource \"google_compute_instance\" \"vm\" {"

func TestExtractAttributeValue(t *testing.T) {
	cases := []struct {
		name    string
		body    string // full resource source
		path    string
		wantVal string
		wantOK  bool
	}{
		{
			name:    "top_level_boolean",
			body:    vmHeader + "\n  can_ip_forward = true\n}\n",
			path:    "can_ip_forward",
			wantVal: "true",
			wantOK:  true,
		},
		{
			name:    "top_level_string",
			body:    vmHeader + "\n  name = \"vm\"\n}\n",
			path:    "name",
			wantVal: "vm",
			wantOK:  true,
		},
		{
			name:    "top_level_number",
			body:    vmHeader + "\n  min_cpu_platform = 7\n}\n",
			path:    "min_cpu_platform",
			wantVal: "7",
			wantOK:  true,
		},
		{
			name:    "nested_one_level",
			body:    vmHeader + "\n  settings {\n    require_ssl = true\n  }\n}\n",
			path:    "settings.require_ssl",
			wantVal: "true",
			wantOK:  true,
		},
		{
			name:    "nested_two_levels",
			body:    vmHeader + "\n  settings {\n    ip_configuration {\n      ipv4_enabled = false\n    }\n  }\n}\n",
			path:    "settings.ip_configuration.ipv4_enabled",
			wantVal: "false",
			wantOK:  true,
		},
		{
			name:   "reference_is_not_literal",
			body:   vmHeader + "\n  region = var.region\n}\n",
			path:   "region",
			wantOK: false,
		},
		{
			name:   "function_call_is_not_literal",
			body:   vmHeader + "\n  tags = length(\"abc\")\n}\n",
			path:   "tags",
			wantOK: false,
		},
		{
			name:   "list_is_not_literal",
			body:   vmHeader + "\n  tags = [\"a\", \"b\"]\n}\n",
			path:   "tags",
			wantOK: false,
		},
		{
			name:   "missing_argument",
			body:   vmHeader + "\n  name = \"vm\"\n}\n",
			path:   "does_not_exist",
			wantOK: false,
		},
		{
			name:   "intermediate_segment_is_argument_not_block",
			body:   vmHeader + "\n  can_ip_forward = true\n}\n",
			path:   "can_ip_forward.sub",
			wantOK: false,
		},
		{
			name:   "missing_intermediate_block",
			body:   vmHeader + "\n  name = \"vm\"\n}\n",
			path:   "settings.require_ssl",
			wantOK: false,
		},
		{
			name:   "heredoc_is_not_literal",
			body:   vmHeader + "\n  metadata_startup_script = <<SCRIPT\n#!/bin/bash\necho hi\nSCRIPT\n}\n",
			path:   "metadata_startup_script",
			wantOK: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := parseBody(t, tc.body)
			gotVal, gotOK := extractAttributeValue(body, tc.path)
			assert.Equal(t, tc.wantOK, gotOK, "ok mismatch")
			if tc.wantOK {
				assert.Equal(t, tc.wantVal, gotVal)
			}
		})
	}
}

func TestExtractAttributeValueNilBody(t *testing.T) {
	val, ok := extractAttributeValue(nil, "can_ip_forward")
	assert.False(t, ok)
	assert.Empty(t, val)

	_, ok = extractAttributeValue(parseBody(t, "resource \"a\" \"b\" {\n  x = true\n}\n"), "")
	assert.False(t, ok)
}
