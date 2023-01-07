package sh

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func Exec(name string, args ...string) Stream {
	return Stream(nil).Exec(name, args...)
}

func (sin Stream) Exec(name string, args ...string) Stream {
	var fin io.Reader = os.Stdin
	if sin != nil {
		fin = sin.Reader()
	}

	xtrace("exec %s %s", name, strings.Join(args, " "))

	r, w := io.Pipe()
	go func() {
		defer w.Close()
		CheckErr(execV(fin, w, os.Stderr, name, args...))
	}()

	return FromReader(r)
}

func (sin Stream) Exec2(name string, args ...string) Stream {
	var fin io.Reader = os.Stdin
	if sin != nil {
		fin = sin.Reader()
	}

	xtrace("exec %s %s", name, strings.Join(args, " "))

	r, w := io.Pipe()
	go func() {
		defer w.Close()
		CheckErr(execV(fin, w, w, name, args...))
	}()

	return FromReader(r)
}

func Exec2(name string, args ...string) Stream {
	return Stream(nil).Exec2(name, args...)
}

func execV(fin io.Reader, fout, ferr io.WriteCloser, name string, args ...string) error {
	cmd := exec.Command(name, glob(args...).Slice()...)
	cmd.Env = os.Environ()
	cmd.Stdin = fin
	cmd.Stdout = fout
	cmd.Stderr = ferr

	return cmd.Run()
}
