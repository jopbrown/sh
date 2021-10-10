package sh

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	stopOnErrorMode  bool
	panicOnErrorMode bool
	xtraceMode       bool
	lastErr          error
)

func Err() error {
	return lastErr
}

func Exit(code int) {
	os.Exit(code)
}

func SetStopOnError(enable bool) bool {
	old := stopOnErrorMode
	stopOnErrorMode = enable
	return old
}

func SetPanicOnError(enable bool) bool {
	old := panicOnErrorMode
	panicOnErrorMode = enable
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

func PanicIfErr() {
	if lastErr != nil {
		panic(fmt.Errorf("%+v", lastErr))
	}
}

func CheckErr(err error) bool {
	if err == nil {
		return false
	}

	lastErr = err

	if stopOnErrorMode {
		ExitIfErr()
	}

	if panicOnErrorMode {
		PanicIfErr()
	}

	fmt.Fprintf(os.Stderr, "error: %+v", lastErr)

	return true
}

func xtrace(format string, v ...interface{}) {
	if xtraceMode {
		fmt.Fprintf(os.Stderr, "+x "+format, v...)
		io.WriteString(os.Stderr, "\n")
	}
}
