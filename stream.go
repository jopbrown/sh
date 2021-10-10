package sh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Stream chan string

func yieldStream(fn func(Stream)) Stream {
	s := make(Stream)
	go func() {
		defer close(s)
		fn(s)
	}()

	return s
}

func From(r io.Reader) Stream {
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

func FromSlice(ss []string) Stream {
	return yieldStream(func(s Stream) {
		for _, line := range ss {
			s <- line
		}
	})
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
