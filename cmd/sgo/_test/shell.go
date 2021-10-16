package main

import (
	"github.com/jopbrown/sh"
)

func init() {
	sh.SetStopOnError(true)
	sh.SetXtrace(true)
}

func main() {
	sh.FromSlice(sh.Args()).Print()
	sh.Echo("hello\nworld").Grep("world").Print()

	TestUnixTool()
	TestExec()
}

func TestUnixTool() {
	sh.Rm("tmp")

	sh.Touch("tmp/aaa.txt", "tmp/bbb.txt")

	sh.Cp("tmp/*.txt", "tmp/2/")
	sh.Cp("tmp/2", "tmp/3")

	sh.Ls("tmp/*.txt").Print()
	sh.Ls("tmp/2/").Print()
	sh.Ls("tmp/3/").Print()

	sh.Mv("tmp/3", "tmp/4")
	sh.Mv("tmp/2/*.txt", "tmp/5/")

	sh.Rm("tmp/*.txt")

	sh.Ls("tmp/").Print()
	sh.Ls("tmp/2/").Print()
	sh.Ls("tmp/4/").Print()
	sh.Ls("tmp/5/").Print()

	sh.Rm("tmp")
	sh.Mkdir("tmp/1", "tmp/2")
	sh.Ls("tmp/").Print()
	sh.Pushd("tmp")
	sh.Ls().Print()
	sh.Popd()

	sh.Echo(sh.Pwd()).Print()
}

func TestExec() {
	sh.Exec("cat", sh.Arg(1)).Grep("Pulvinar").Exec("cat").Print()

	input := `aaa
bbb
ccc
ddd`

	sh.Echo(input).Exec("cat").Print()
}
