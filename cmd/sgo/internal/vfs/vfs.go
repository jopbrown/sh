package vfs

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing/fstest"
)

// GOVFS wrap fstest.MapFS to only add go source files
// force to path normalize to support both windows and unix-like OS
type GOVFS fstest.MapFS

func New() GOVFS {
	return make(GOVFS)
}

func (fsys GOVFS) Add(rname, vname string) error {
	if !isGoFile(rname) {
		return fmt.Errorf("%s is not go source file", rname)
	}

	if !isGoFile(vname) {
		return fmt.Errorf("%s is invalid virtual go source file", rname)
	}

	data, err := os.ReadFile(rname)
	if err != nil {
		return err
	}

	fsys[normalize(vname)] = &fstest.MapFile{Data: data}
	return nil
}

func (fsys GOVFS) Open(name string) (fs.File, error) {
	return fstest.MapFS(fsys).Open(normalize(name))
}

func normalize(vname string) string {
	return path.Clean(filepath.ToSlash(vname))
}

func isGoFile(name string) bool {
	return strings.HasSuffix(name, ".go") && filepath.Base(name) != ".go"
}
