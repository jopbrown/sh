`sh` is a library to let you coding `go` like a shell script.

## Examples

```go
package main

import (
	"os"
	"strings"

	"github.com/jopbrown/sh"
)

func init() {
	sh.SetStopOnError(true)
	sh.SetXtrace(true)
}

func main() {
	// echo "$#"
	sh.Echof("%d", sh.ArgN())
	// echo "$0"
	sh.Echo(sh.Arg(0))
	// echo "$1"
	sh.Echo(sh.Arg(1))

	// for arg in $@; do
	// 	echo $arg
	// done
	sh.FromSlice(sh.Args()).Print()

	// pwd
	sh.Echo(sh.Pwd())
	// echo "$PATH"
	sh.Echo(os.Getenv("PATH")).Print()

	// rm -rf tmp*
	sh.Rm("tmp*")

	// ls
	sh.Ls()
	// ls *.go
	sh.Ls("*.go")

	// mkdir -p tmp tmp2
	sh.Mkdir("tmp", "tmp2")
	// cp *.go tmp/.
	sh.Cp("*.go", "tmp/.")

	// cp -r tmp tmp3
	sh.Cp("tmp", "tmp3")
	// mv tmp/*.go tmp2/.
	sh.Mv("tmp/*.go", "tmp2/.")

	// pushd tmp2
	sh.Pushd("tmp2")
	// touch 1.txt 2.txt
	sh.Touch("1.txt", "2.txt")

	// if [[ -e 1.txt ]]; then
	//     echo "exist"
	// fi
	if sh.Exists("1.txt") {
		sh.Echo("exist").Print()
	}

	// if [[ -f 1.txt ]]; then
	//     echo "file exist"
	// fi
	if sh.ExistsFile("1.txt") {
		sh.Echo("file exist").Print()
	}

	// if [[ -d 1.txt ]]; then
	//     echo "dir exist"
	// fi
	if sh.ExistsDir("1.txt") {
		sh.Echo("dir exist").Print()
	}

	sh.Ls("*.txt")
	// ls *.txt
	sh.Rm("*.txt", "*.go")
	// rm *.txt *.go

	sh.Popd()
	// popd

	// pushd tmp3
	sh.Pushd("tmp3")

	// grep 'func' main.go
	sh.Grep("func", "main.go").Print()
	// cat *.go | grep 'func'
	sh.Cat("*.go").Grep("func").Print()

	// sed -e 's/func/xxxx/g' main.go > main2.txt
	sh.Sed(func(s string) string {
		return strings.ReplaceAll(s, "func", "xxxx")
	}, "main.go").PrintToFile("main2.txt", false)

	// cat main.go | sed -e 's/func/ssss/g' >> main2.txt
	sh.Cat("main.go").Sed(func(s string) string {
		return strings.ReplaceAll(s, "func", "ssss")
	}).PrintToFile("main2.txt", true)

	// tar zcvf xxx.tar.gz *
	sh.Exec("tar", "zcvf", "xxx.tar.gz", "*").Print()
	if sh.Err() != nil {
		sh.Exit(-1)
	}

	// echo "xxxxx" | base64
	sh.Echo("xxxxx").Exec("base64").Print()

	// popd
	sh.Popd()
}

```

## sgo

sgo is a interpreter to run go without go installed.

Embed `sh` and subset of [stdlib](https://github.com/jopbrown/sh/blob/main/cmd/sgo/extract.txt).

### Install
```
go install github.com/jopbrown/sh/cmd/sgo
```

### Usage
```bash
# sgo [-vendor <VENDOR_DIR>] <ENTRY_GO_FILE> [args...]
sgo main.go arg1 arg2
```

## License
Apache 2.0.
