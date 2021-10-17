package sh

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Glob(pathList ...string) Stream {
	return yieldStream(func(s Stream) {
		for _, path := range pathList {
			if !strings.ContainsRune(path, '*') {
				s <- path
				continue
			}
			matches, err := filepath.Glob(path)
			if err != nil {
				s <- path
				continue
			}

			for _, match := range matches {
				s <- match
			}
		}
	})
}

func Ls(args ...string) Stream {
	if len(args) != 0 {
		xtrace("ls %s", strings.Join(args, " "))
		for i, arg := range args {
			arg = filepath.ToSlash(arg)
			if len(arg) > 0 && arg[len(arg)-1] == '/' {
				arg += "*"
				args[i] = arg
			}
		}
		return Glob(args...)
	}

	xtrace("ls")
	return yieldStream(func(s Stream) {
		entries, err := os.ReadDir(".")
		if CheckErr(err) {
			return
		}

		for _, entry := range entries {
			s <- entry.Name()
		}
	})
}

func Rm(args ...string) {
	argCnt := len(args)
	if argCnt < 1 {
		CheckErr(errors.New("rm: missing args"))
	}

	for fpath := range Glob(args...) {
		xtrace("rm %s", fpath)
		err := os.RemoveAll(fpath)
		if CheckErr(err) {
			return
		}
	}
}

func Cp(args ...string) {
	argCnt := len(args)
	if argCnt < 2 {
		CheckErr(fmt.Errorf("cp: wrong number of args: %v", args))
	}

	srcList, dst := args[:argCnt-1], args[argCnt-1]

	for src := range Glob(srcList...) {
		copy(dst, src)
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func ExistsDir(name string) bool {
	finfo, err := os.Stat(name)
	if err != nil {
		return false
	}

	return finfo.IsDir()
}

func ExistsFile(name string) bool {
	finfo, err := os.Stat(name)
	if err != nil {
		return false
	}

	return finfo.Mode().IsRegular()
}

func copy(dst, src string) {
	if ExistsDir(src) {
		copyDir(dst, src)
	} else if ExistsFile(src) {
		copyFile(dst, src)
	} else {
		CheckErr(fmt.Errorf("cp: src not exist: %s", src))
	}
}

func copyFile(dst, src string) {
	srcInfo, err := os.Stat(src)
	if CheckErr(err) {
		return
	}

	dst = filepath.ToSlash(dst)
	if ExistsDir(dst) || dst[len(dst)-1] == '/' {
		dst = filepath.Join(dst, filepath.Base(src))
	}

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if CheckErr(err) {
		return
	}

	fin, err := os.Open(src)
	if CheckErr(err) {
		return
	}
	defer fin.Close()

	fout, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode().Perm())
	if CheckErr(err) {
		return
	}
	defer fout.Close()

	xtrace("cp %s %s", src, dst)

	_, err = io.Copy(fout, fin)
	if CheckErr(err) {
		return
	}
}

func copyDir(dstDir, srcDir string) {
	err := filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		if srcDir == path {
			return nil
		}
		src := path
		relsrc, err := filepath.Rel(srcDir, path)
		if CheckErr(err) {
			return nil
		}
		dst := filepath.Join(dstDir, relsrc)
		if ExistsDir(src) {
			copyDir(dst, src)
		} else {
			copyFile(dst, src)
		}
		return nil
	})

	CheckErr(err)
}

func Mv(args ...string) {
	argCnt := len(args)
	if argCnt < 2 {
		CheckErr(fmt.Errorf("mv: wrong number of args: %v", args))
	}

	srcList, dst := args[:argCnt-1], args[argCnt-1]
	dst = filepath.ToSlash(dst)

	err := os.MkdirAll(filepath.Dir(dst), 0755)
	if CheckErr(err) {
		return
	}

	manySrc := false
	isDstDir := ExistsDir(dst) || dst[len(dst)-1] == '/'

	for src := range Glob(srcList...) {
		if !Exists(src) {
			CheckErr(fmt.Errorf("mv: src not exist: %v", src))
			break
		}

		if manySrc && !isDstDir {
			CheckErr(fmt.Errorf("mv: too many src move to single dst: %s %s", src, dst))
			break
		}

		manySrc = true

		if ExistsDir(src) && dst[len(dst)-1] == '/' {
			dstDir := dst
			if dstDir[len(dstDir)-1] == '/' {
				dstDir = filepath.Join(dst, filepath.Base(src))
			}
			xtrace("mv %s %s", src, dstDir)
			err = os.Rename(src, dstDir)
			if CheckErr(err) {
				break
			}
			continue
		}

		dstFile := dst
		if isDstDir {
			dstFile = filepath.Join(dst, filepath.Base(src))
		}

		xtrace("mv %s %s", src, dstFile)
		err = os.Rename(src, dstFile)
		if CheckErr(err) {
			break
		}
	}
}

func Touch(args ...string) {
	for fname := range Glob(args...) {
		xtrace("touch %s", fname)

		err := os.MkdirAll(filepath.Dir(fname), 0755)
		if CheckErr(err) {
			break
		}

		if ExistsFile(fname) {
			now := time.Now()
			if CheckErr(os.Chtimes(fname, now, now)) {
				break
			}
		} else {
			f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0666)
			if CheckErr(err) {
				break
			}
			f.Close()
		}
	}
}

func Cat(args ...string) Stream {
	return yieldStream(func(s Stream) {
		for fpath := range Glob(args...) {
			f, err := os.Open(fpath)
			if CheckErr(err) {
				return
			}

			xtrace("cat %s", fpath)
			for line := range FromReader(f) {
				s <- line
			}
			f.Close()
		}
	})
}

func Cd(dir string) {
	xtrace("cd %s", dir)
	absDir, err := filepath.Abs(dir)
	if CheckErr(err) {
		return
	}
	CheckErr(os.Chdir(absDir))
}

func Mkdir(args ...string) {
	argCnt := len(args)
	if argCnt < 1 {
		CheckErr(errors.New("mkdir: missing args"))
	}

	for _, arg := range args {
		xtrace("mkdir %s", arg)
		err := os.MkdirAll(arg, 0755)
		if CheckErr(err) {
			return
		}
	}
}

var (
	queueDir []string
)

func Pushd(dir string) {
	pwd := Pwd()

	xtrace("pushd %s", dir)
	absDir, err := filepath.Abs(dir)
	if CheckErr(err) {
		return
	}
	if CheckErr(os.Chdir(absDir)) {
		return
	}

	queueDir = append(queueDir, pwd)
}

func Popd() {
	if len(queueDir) == 0 {
		CheckErr(errors.New("popd: no dir in queue"))
		return
	}

	dir := queueDir[len(queueDir)-1]

	xtrace("popd %s", dir)
	if CheckErr(os.Chdir(dir)) {
		return
	}

	queueDir = queueDir[:len(queueDir)-1]
}

func Pwd() string {
	pwd, err := filepath.Abs(".")
	if CheckErr(err) {
		return "."
	}

	return pwd
}

func Chmod(perm int, fileList ...string) {
	for fname := range Glob(fileList...) {
		CheckErr(os.Chmod(fname, os.FileMode(perm)))
	}
}
