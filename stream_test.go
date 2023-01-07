package sh

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetExitOnError(true)
	SetXtrace(true)
	os.Exit(m.Run())
}

func TestEcho(t *testing.T) {
	var input, output string

	input = "hello world"
	output = Echo(input).String()
	assert.Equal(t, input+"\n", output)

	input = "hello\nworld"
	output = Echo(input).String()
	assert.Equal(t, input+"\n", output)

	input = "hello\nworld\n"
	output = Echo(input).String()
	assert.Equal(t, input, output)

	output = Echof("hello %s", "world").String()
	assert.Equal(t, "hello world\n", output)
}

func TestFromSliceAndPrintToFile(t *testing.T) {
	fname := "tmp2/from_slice_test"
	Rm(fname)

	Echo("START").PrintToFile(fname, false)

	ss := []string{"hello", "world"}
	FromSlice(ss).PrintToFile(fname, true)

	golden := `START
hello
world
`

	assert.Equal(t, golden, Cat(fname).String())

}
