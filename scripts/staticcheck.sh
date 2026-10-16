#!/bin/sh
# Run staticcheck and fail on any finding outside generated code.
#
# Generated files are excluded (same header detection as check-coverage.sh):
# ANTLR 4.13.x's Go target emits a dead `_la` assignment that trips SA4006,
# and generated code is tool output, not project code. Every other finding —
# including any new one in a future generated file's package that is not
# header-marked — still fails the build.
set -eu

GENERATED=$(grep -rlE '^// Code generated .* DO NOT EDIT\.' --include='*.go' . 2>/dev/null | sed 's|^\./||' || true)

# staticcheck exits non-zero when it reports findings; capture either way.
# A non-empty stderr means staticcheck itself failed (e.g. a compile error),
# which must fail the build rather than masquerade as a clean run.
staticcheck -f json ./... > .staticcheck.json 2> .staticcheck.err || true
if [ -s .staticcheck.err ]; then
  echo "staticcheck failed:" >&2
  cat .staticcheck.err >&2
  rm -f .staticcheck.json .staticcheck.err
  exit 1
fi

if [ -n "$GENERATED" ]; then
  # Match the absolute path suffix of each generated file so only exact
  # locations in generated files are filtered out.
  PATTERN=$(printf '%s\n' $GENERATED | sed 's/\./\\./g; s|^|/|' | awk 'NR > 1 { printf "|" } { printf "%s", $0 }')
  grep -vE "\"file\": ?\"[^\"]*(${PATTERN})\"" .staticcheck.json > .staticcheck.filtered || true
else
  cp .staticcheck.json .staticcheck.filtered
fi

if [ -s .staticcheck.filtered ]; then
  echo "staticcheck findings (generated code excluded):" >&2
  cat .staticcheck.filtered >&2
  rm -f .staticcheck.json .staticcheck.err .staticcheck.filtered
  exit 1
fi

rm -f .staticcheck.json .staticcheck.err .staticcheck.filtered
echo "staticcheck OK (generated code excluded)."
