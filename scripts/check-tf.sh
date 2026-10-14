#!/bin/sh
# Validate exported Terraform scaffolds with the terraform CLI.
#
# Usage: sh scripts/check-tf.sh <taralizer-binary> <model.yaml> [profile]
#
# Exports the model into a temporary directory, then runs `terraform init` +
# `terraform validate` on the result so scaffold regressions (bad GCP names,
# missing required arguments, unknown resource types) fail the build.
set -eu

BIN="${1:?usage: check-tf.sh <taralizer-binary> <model.yaml> [profile]}"
MODEL="${2:?usage: check-tf.sh <taralizer-binary> <model.yaml> [profile]}"
PROFILE="${3:-./profiles/google.yaml}"

TF_EXPORT_DIR=$(mktemp -d /tmp/taralizer-tf.XXXXXX)
trap 'rm -rf "$TF_EXPORT_DIR"' EXIT

"$BIN" terraform export "$MODEL" -o "$TF_EXPORT_DIR" --profile "$PROFILE" --write-back=false

cd "$TF_EXPORT_DIR"
terraform init -backend=false -input=false >/dev/null
terraform validate
