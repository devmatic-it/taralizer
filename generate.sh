#!/bin/sh
# Regenerate the ANTLR parser from pkg/terraform/Terraform.g4.
# Requires antlr 4.13.x (e.g. `brew install antlr`) and a matching
# github.com/antlr4-go/antlr/v4 runtime in go.mod.
#
# Output is generated into a temp dir and copied into pkg/terraform because
# on some macOS setups the Java process is blocked from writing directly
# under ~/Documents (TCC), which makes `antlr -o <project dir>` fail with
# CANNOT_WRITE_FILE.
set -e

cd "$(dirname "$0")/pkg/terraform"
OUT=$(mktemp -d)
trap 'rm -rf "$OUT"' EXIT

antlr -Dlanguage=Go -no-visitor -package terraform -o "$OUT" Terraform.g4

cd ../..

cp "$OUT"/terraform_lexer.go \
   "$OUT"/terraform_parser.go \
   "$OUT"/terraform_listener.go \
   "$OUT"/terraform_base_listener.go \
   pkg/terraform/

cp "$OUT"/Terraform.interp \
   "$OUT"/Terraform.tokens \
   "$OUT"/TerraformLexer.interp \
   "$OUT"/TerraformLexer.tokens \
   pkg/terraform/

echo "Regenerated ANTLR parser in pkg/terraform/"
