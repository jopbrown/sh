package sh

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	exitOnErrorMode bool
	xtraceMode      bool
	lastErr         error
)

func Err() error {
	return lastErr
}

func Exit(code int) {
	os.Exit(code)
}

func SetExitOnError(enable bool) bool {
	old := exitOnErrorMode
	exitOnErrorMode = enable
	return old
}

func SetXtrace(enable bool) bool {
	old := xtraceMode
	xtraceMode = enable
	return old
}

func ExitIfErr() {
	if lastErr != nil {
		log.Fatalf("%+v", lastErr)
	}
}

func CheckErr(err error) bool {
	if err == nil {
		return false
	}

	lastErr = err

	if exitOnErrorMode {
		ExitIfErr()
	}

	fmt.Fprintf(os.Stderr, "error: %+v", lastErr)

	return true
}

func xtrace(format string, v ...interface{}) {
	if xtraceMode {
		fmt.Fprintf(os.Stderr, "+ "+format, v...)
		io.WriteString(os.Stderr, "\n")
	}
}

func Args() []string {
	return os.Args[1:]
}

func ArgN() int {
	return len(os.Args) - 1
}

func Arg(n int) string {
	if n < 0 || n > ArgN() {
		// CheckErr(fmt.Errorf("arg[%d] is not available", n))
		return ""
	}
	return os.Args[n]
}
