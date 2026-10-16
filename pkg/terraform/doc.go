// Package terraform imports and exports Terraform configurations as
// taralizer models. ImportFromDir parses a directory of .tf files with the
// ANTLR-based parser and maps resource blocks to technical assets and trust
// boundaries via a technology profile; ExportToTerraform renders a taralizer
// model back into Terraform scaffolds with traceability markers.
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
