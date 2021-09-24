package cmd

import (
	//"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv("GO_ENV", "testing")
}

func TestDiagramCmd(t *testing.T) {
	diagramCmd.Run(diagramCmd, []string{"../examples/gcp/bank_of_anthos.yaml"})
}
