package sh

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func slashFN(str string) string {
	return filepath.ToSlash(str)
}

func TestMain(m *testing.M) {
	SetStopOnError(true)
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
}

func TestTouchCpMvRm(t *testing.T) {

	Rm("tmp")

	Touch("tmp/aaa.txt", "tmp/bbb.txt")

	Cp("tmp/*.txt", "tmp/2/")
	Cp("tmp/2", "tmp/3")

	assert.Equal(t, []string{"tmp/aaa.txt", "tmp/bbb.txt"}, Ls("tmp/*.txt").Sed(slashFN).Slice())
	assert.Equal(t, []string{"tmp/2/aaa.txt", "tmp/2/bbb.txt"}, Ls("tmp/2/").Sed(slashFN).Slice())
	assert.Equal(t, []string{"tmp/3/aaa.txt", "tmp/3/bbb.txt"}, Ls("tmp/3/").Sed(slashFN).Slice())

	Mv("tmp/3", "tmp/4")
	Mv("tmp/2/*.txt", "tmp/5/")

	Rm("tmp/*.txt")

	assert.Equal(t, []string{"tmp/2", "tmp/4", "tmp/5"}, Ls("tmp/").Sed(slashFN).Slice())
	assert.Equal(t, []string{}, Ls("tmp/2/").Sed(slashFN).Slice())
	assert.Equal(t, []string{"tmp/4/aaa.txt", "tmp/4/bbb.txt"}, Ls("tmp/4/").Sed(slashFN).Slice())
	assert.Equal(t, []string{"tmp/5/aaa.txt", "tmp/5/bbb.txt"}, Ls("tmp/5/").Sed(slashFN).Slice())
}

func TestMkdirCd(t *testing.T) {
	Rm("tmp")

	Mkdir("tmp/1", "tmp/2")

	assert.Equal(t, []string{"tmp/1", "tmp/2"}, Ls("tmp/").Sed(slashFN).Slice())

	Pushd("tmp")

	assert.Equal(t, []string{"1", "2"}, Ls().Sed(slashFN).Slice())

	Popd()

	assert.Equal(t, []string{"tmp/1", "tmp/2"}, Ls("tmp/").Sed(slashFN).Slice())

	Cd("tmp")
	assert.Equal(t, []string{"1", "2"}, Ls().Sed(slashFN).Slice())

	Cd("..")
}
