package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/jopbrown/sh/cmd/sgo/internal/sgolib"
	"github.com/jopbrown/sh/cmd/sgo/internal/vfs"
	"github.com/traefik/yaegi/interp"
)

const (
	_GOPATH     = "_pkg"
	_VENDORPATH = _GOPATH + "/src/"
)

var context struct {
	entry string
	// ventry    string
	args      []string
	vendorDir string
	fsys      fs.FS
}

func init() {
	flag.StringVar(&context.vendorDir, "vendor", "", "vendor folder")
	parseArgs()
}

func main() {
	inter := interp.New(interp.Options{
		GoPath:               _GOPATH,
		SourcecodeFilesystem: context.fsys,
	})

	inter.Use(sgolib.Symbols)

	os.Args = context.args
	_, err := inter.EvalPath(context.entry)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func parseArgs() {
	flag.Parse()
	context.args = flag.Args()
	if len(context.args) < 1 {
		log.Fatalf("usage: %s [-vendor <VENDOR_DIR>] <ENTRY_GO_FILE> [args...]", os.Args[0])
	}

	context.entry = context.args[0]

	fsys := vfs.New()
	fsys.Add(context.entry, context.entry)

	if len(context.vendorDir) == 0 {
		context.vendorDir = filepath.Join(filepath.Dir(context.entry), "vendor")
	}

	err := filepath.Walk(context.vendorDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		rname := path
		relPath, err := filepath.Rel(context.vendorDir, rname)
		if err != nil {
			return err
		}
		vname := _VENDORPATH + relPath
		fsys.Add(rname, vname)

		return nil
	})

	if err != nil {
		log.Fatalf("%+v", err)
	}

	context.fsys = fsys
}
