package cmd

import (
	//"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv("GO_ENV", "testing")
}

func Test_Diagram_Cmd(t *testing.T) {
	diagramCmd.Run(diagramCmd, []string{"../example/model.yaml"})
	//assert.Empty(t, err)
}
