package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec_Cat(t *testing.T) {
	input := `aaa
bbb
ccc
ddd
`

	output := Echo(input).Exec("cat").String()
	assert.Equal(t, input, output)
}

func TestExec_Cat_RoundTrip(t *testing.T) {
	moduleLine := Exec("cat", "go.mod").Grep("module").Exec("cat").String()
	assert.Equal(t, "module github.com/jopbrown/sh\n", moduleLine)
}
