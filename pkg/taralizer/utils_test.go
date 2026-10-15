// Package taralizer Threat and Risk Analyzer
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
package taralizer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMapIntValue(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]interface{}
		key     string
		want    int64
		wantErr bool
	}{
		{"float64", map[string]interface{}{"likelihood": float64(3)}, "likelihood", 3, false},
		{"int", map[string]interface{}{"impact": int(2)}, "impact", 2, false},
		{"int64", map[string]interface{}{"severity": int64(9)}, "severity", 9, false},
		{"json.Number", map[string]interface{}{"cwe": json.Number("79")}, "cwe", 79, false},
		{"json.Number invalid", map[string]interface{}{"cwe": json.Number("abc")}, "cwe", 0, true},
		{"non-numeric value", map[string]interface{}{"likelihood": "high"}, "likelihood", 0, true},
		{"missing key", map[string]interface{}{"other": float64(1)}, "likelihood", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMapIntValue(tt.data, tt.key, "test")
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetMapStringValue(t *testing.T) {
	t.Run("string value", func(t *testing.T) {
		got, err := GetMapStringValue(map[string]interface{}{"id": "risk-1"}, "id", "test")
		require.NoError(t, err)
		assert.Equal(t, "risk-1", got)
	})

	t.Run("non-string value", func(t *testing.T) {
		_, err := GetMapStringValue(map[string]interface{}{"id": float64(1)}, "id", "test")
		require.Error(t, err)
	})

	t.Run("missing key", func(t *testing.T) {
		_, err := GetMapStringValue(map[string]interface{}{"other": "x"}, "id", "test")
		require.Error(t, err)
	})
}

func TestStringWriter(t *testing.T) {
	var buf string
	sw := NewStringWriter(&buf)

	n, err := sw.Write([]byte("hello "))
	require.NoError(t, err)
	assert.Equal(t, 6, n)

	n, err = sw.Write([]byte("world"))
	require.NoError(t, err)
	assert.Equal(t, 5, n)

	assert.Equal(t, "hello world", sw.String())
	assert.Equal(t, "hello world", buf)
}
