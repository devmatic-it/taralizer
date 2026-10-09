package taralizer_test

import (
	"testing"

	"github.com/devmatic-it/taralizer/pkg/taralizer"
)

func TestCWEFromMetadata(t *testing.T) {
	// Create a Taralizer instance
	svc := taralizer.NewTaralizer("asvs")

	// Evaluate a model with risks
	report := svc.Evaluate("../../examples/gcp/bank_of_anthos.yaml")

	// Check that risks have CWE numbers populated from METADATA blocks
	for _, risk := range report.Risks {
		if risk.Title == "Missing Injection Protection" || risk.Id == "injection-protection@frontend" {
			if risk.Cwe == 0 {
				t.Errorf("Expected CWE to be populated from METADATA block, but got 0 for risk: %s", risk.Id)
			} else {
				t.Logf("✓ Risk %s has CWE: %d", risk.Id, risk.Cwe)
			}
		}
	}
}
