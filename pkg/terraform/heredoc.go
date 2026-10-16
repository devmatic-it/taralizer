// Package terraform heredoc normalization: rewrites HCL2 heredocs with
// arbitrary delimiters into string literals so the ANTLR grammar (which only
// knows the fixed <<EOF / <<-EOF / <<DESCRIPTION tags) can parse real-world
// Terraform files.
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

import "strings"

// heredocPlaceholder replaces each normalized heredoc. It is a quoted string
// literal, which is a valid value in every position where an HCL2 heredoc can
// appear (arguments, list elements, map values, function arguments).
const heredocPlaceholder = `"__taralizer_heredoc__"`

// normalizeHeredocs rewrites every HCL2 heredoc (<<TAG or <<-TAG) in src to
// the placeholder string literal. The ANTLR grammar only knows the fixed tags
// <<EOF, <<-EOF and <<DESCRIPTION, but real-world Terraform uses arbitrary
// delimiters (<<SCRIPT, <<-EOT, ...); replacing the whole heredoc with a
// string keeps the block structure intact so resource blocks can still be
// extracted from such files.
//
// The scan is string- and comment-aware: << inside quoted strings or comments
// is left alone. If a heredoc is never closed, src is returned unchanged so
// the parser reports its usual syntax error instead of silently accepting a
// broken file. Newlines inside dropped heredoc bodies are preserved so line
// numbers in later parse errors stay accurate.
func normalizeHeredocs(src string) string {
	var b strings.Builder
	i, n := 0, len(src)

	inString := false
	inLineComment := false
	inBlockComment := false
	heredocDelim := ""

	for i < n {
		if heredocDelim != "" {
			// Inside a heredoc body: drop lines until the closing delimiter.
			lineEnd := i
			for lineEnd < n && src[lineEnd] != '\n' {
				lineEnd++
			}
			if strings.TrimSpace(src[i:lineEnd]) == heredocDelim {
				heredocDelim = "" // closing line consumed; resume normal scan
				i = lineEnd
				continue
			}
			b.WriteByte('\n') // preserve line count for error positions
			i = lineEnd + 1   // skip the dropped line and its newline (if any)
			continue
		}

		c := src[i]
		switch {
		case inLineComment:
			b.WriteByte(c)
			if c == '\n' {
				inLineComment = false
			}
			i++
		case inBlockComment:
			if c == '*' && i+1 < n && src[i+1] == '/' {
				b.WriteString("*/")
				i += 2
				inBlockComment = false
			} else {
				b.WriteByte(c)
				i++
			}
		case inString:
			if c == '\\' && i+1 < n {
				b.WriteString(src[i : i+2])
				i += 2
				continue
			}
			b.WriteByte(c)
			if c == '"' {
				inString = false
			}
			i++
		default:
			switch {
			case c == '"':
				inString = true
				b.WriteByte(c)
				i++
			case c == '#' || (c == '/' && i+1 < n && src[i+1] == '/'):
				inLineComment = true
				b.WriteByte(c)
				i++
			case c == '/' && i+1 < n && src[i+1] == '*':
				inBlockComment = true
				b.WriteString("/*")
				i += 2
			case c == '<' && i+1 < n && src[i+1] == '<':
				if tag, end := heredocTag(src[i:]); tag != "" {
					b.WriteString(heredocPlaceholder)
					// Copy the remainder of the opening line verbatim, then
					// enter heredoc-body mode. end is relative to src[i:].
					j := i + end
					for j < n && src[j] != '\n' {
						b.WriteByte(src[j])
						j++
					}
					i = j
					heredocDelim = tag
				} else {
					b.WriteByte(c)
					i++
				}
			default:
				b.WriteByte(c)
				i++
			}
		}
	}

	if heredocDelim != "" {
		return src // unterminated heredoc: let the parser report it
	}
	return b.String()
}

// heredocTag reports the delimiter tag of a heredoc opening at src[0:2]
// ("<<" or "<<-") and the index just past it. The tag must start immediately
// after the opening (HCL2 allows no whitespace) and is an identifier. It
// returns ("", 0) when src does not start a heredoc (e.g. "<< " or "<<<").
func heredocTag(src string) (tag string, end int) {
	i := 2
	if i < len(src) && src[i] == '-' {
		i++
	}
	if i >= len(src) || !isTagStart(src[i]) {
		return "", 0
	}
	start := i
	i++
	for i < len(src) && isTagChar(src[i]) {
		i++
	}
	return src[start:i], i
}

func isTagStart(c byte) bool {
	return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isTagChar(c byte) bool {
	return isTagStart(c) || (c >= '0' && c <= '9')
}
