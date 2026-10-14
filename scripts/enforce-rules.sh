#!/bin/sh
# Hard constraints that fail the build if the agent cheats.
set -u

if grep -rnw . --include='*.go' -e '//nolint'; then
  echo "FAIL: //nolint is forbidden"
  exit 1
fi

if grep -rnw . --include='*.go' -e '_ = err' -e '_ = fmt.Errorf'; then
  echo "FAIL: Blank identifier error suppression is forbidden"
  exit 1
fi

if grep -rnw . --include='*.go' -e 't.Skip()'; then
  echo "FAIL: Skipping tests is forbidden"
  exit 1
fi

exit 0
