#!/bin/sh
# Run go vet and fail on any finding outside generated code.
#
# Generated files are excluded (same header detection as check-coverage.sh
# and staticcheck.sh, per CONSTRAINTS.md §3): ANTLR 4.13.x's Go target emits
# a trailing `goto errorExit` after `return localctx` that trips go vet's
# unreachable analyzer. Every finding in hand-written code or any build
# failure still fails the build.
set -eu

GENERATED=$(grep -rlE '^// Code generated .* DO NOT EDIT\.' --include='*.go' . 2>/dev/null | sed 's|^\./||' || true)

go vet ./... > .govet.raw 2>&1 || true

if [ -n "$GENERATED" ]; then
  PATTERN=$(printf '%s\n' $GENERATED | sed 's/\./\\./g' | awk 'NR > 1 { printf "|" } { printf "%s", $0 }')
  awk -v pat="^(\\./)?(${PATTERN}):[0-9]+:" '
    /^# / && !/\[/ { hdr = $0; next }
    $0 ~ pat { next }
    {
      if (hdr != "") { print hdr; hdr = "" }
      print
    }
  ' .govet.raw > .govet.filtered
else
  cp .govet.raw .govet.filtered
fi

if [ -s .govet.filtered ]; then
  echo "go vet findings (generated code excluded):" >&2
  cat .govet.filtered >&2
  rm -f .govet.raw .govet.filtered
  exit 1
fi

rm -f .govet.raw .govet.filtered
echo "go vet OK (generated code excluded)."
