// Package terraform attribute-level drift (Phase A4): extract static literal
// values from parsed Terraform resource bodies so model-declared attributes
// can be compared against the infrastructure.
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
	"strings"
)

// extractAttributeValue resolves a dot-separated Terraform argument path
// (e.g. "settings.ip_configuration.ipv4_enabled") within a resource body and
// returns its value when it is a static literal. ok is false in two cases:
// the path does not exist, or the value at that path is not a static literal
// (a reference such as var.x, a function call, a collection, or a compound
// expression). In both cases the value cannot be determined statically, so
// attribute drift is reported as a warning rather than a finding.
//
// Intermediate path segments must be nested blocks; the final segment must be
// an argument. When a segment matches several blocks, the first occurrence
// wins (consistent with ImportFromDir's duplicate handling).
func extractAttributeValue(body *BlockbodyContext, path string) (value string, ok bool) {
	if body == nil || path == "" {
		return "", false
	}

	segments := strings.Split(path, ".")
	current := body
	for i, seg := range segments {
		if i == len(segments)-1 {
			for _, arg := range current.AllArgument() {
				if arg.Identifier().GetText() == seg {
					return literalValue(arg.Expression())
				}
			}
			return "", false // no argument with this name at this level
		}

		descended := false
		for _, blk := range current.AllBlock() {
			if blk.Blocktype().GetText() != seg {
				continue
			}
			next, isBody := blk.Blockbody().(*BlockbodyContext)
			if !isBody {
				return "", false
			}
			current = next
			descended = true
			break // first matching block wins
		}
		if !descended {
			return "", false // no nested block with this name at this level
		}
	}

	return "", false // unreachable: the loop returns on the final segment
}

// literalValue returns the static scalar value of an expression, or ok=false
// when the expression is not a plain literal. Only string, boolean, number and
// null literals are comparable; references (var.x), function calls, file(),
// collections and compound expressions are not statically determinable. A
// normalized heredoc (see normalizeHeredocs) is also non-literal: its real
// body was dropped, so the placeholder must not be compared as a value.
func literalValue(expr IExpressionContext) (value string, ok bool) {
	section := expr.Section()
	if section == nil {
		return "", false // binary/unary/index/forloop/parenthesized expression
	}

	val := section.Val()
	if val == nil {
		return "", false // list_ or map_ collection
	}

	switch {
	case val.String_() != nil:
		if val.String_().MULTILINESTRING() != nil {
			return "", false // heredoc body is not a comparable scalar
		}
		s := unquote(val.String_().STRING().GetText())
		if s == unquote(heredocPlaceholder) {
			return "", false // normalized heredoc: real value unknown
		}
		return s, true
	case val.BOOL() != nil:
		return val.BOOL().GetText(), true // "true" or "false"
	case val.Signed_number() != nil:
		return val.Signed_number().GetText(), true
	case val.NULL_() != nil:
		return "null", true
	default:
		return "", false // identifier reference, function call, file(), ...
	}
}
