#!/bin/sh
antlr4 -Dlanguage=Go -no-visitor -package terraform Terraform.g4
