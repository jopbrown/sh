package sh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const _MAX_BUFFER_LINES = 1000

type Stream chan string

func yieldStream(fn func(Stream)) Stream {
	s := make(Stream, _MAX_BUFFER_LINES)
	go func() {
		defer close(s)
		fn(s)
	}()

	return s
}

func FromReader(r io.Reader) Stream {
	sc := bufio.NewScanner(r)

	return yieldStream(func(s Stream) {
		for sc.Scan() {
			s <- sc.Text()
		}

		if err := sc.Err(); err != nil {
			CheckErr(err)
		}
	})
}

func From(ss ...string) Stream {
	return yieldStream(func(s Stream) {
		for _, line := range ss {
			s <- line
		}
	})
}

func FromFields(fields string) Stream {
	return yieldStream(func(s Stream) {
		for _, line := range strings.Fields(fields) {
			s <- line
		}
	})
}

func FromSlice(ss []string) Stream {
	return yieldStream(func(s Stream) {
		for _, line := range ss {
			s <- line
		}
	})
}

func Echo(msg string) Stream {
	r := strings.NewReader(msg)
	xtrace("echo %s", msg)
	return FromReader(r)
}

func Echof(format string, v ...interface{}) Stream {
	return Echo(fmt.Sprintf(format, v...))
}

func (s Stream) Print() {
	s.PrintTo(os.Stdout)
}

func (s Stream) PrintToErr() {
	s.PrintTo(os.Stderr)
}

func (s Stream) PrintToDevNull() {
	s.PrintTo(io.Discard)
}

func (s Stream) PrintTo(w io.Writer) {
	for line := range s {
		fmt.Fprintln(w, line)
	}
}

func (s Stream) PrintToFile(fname string, append bool) {
	flag := os.O_CREATE | os.O_WRONLY
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}

	err := os.MkdirAll(filepath.Dir(fname), 0755)
	if CheckErr(err) {
		return
	}

	f, err := os.OpenFile(fname, flag, 0644)
	if CheckErr(err) {
		return
	}
	defer f.Close()

	s.PrintTo(f)
}

func (s Stream) Slice() []string {
	ss := make([]string, 0, 10)
	for line := range s {
		ss = append(ss, line)
	}

	return ss
}

func (s Stream) String() string {
	sb := &strings.Builder{}
	s.PrintTo(sb)
	return sb.String()
}

func (s Stream) Reader() io.Reader {
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		s.PrintTo(w)
	}()

	return r
}
