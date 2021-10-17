package sh

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetStopOnError(true)
	SetXtrace(true)
	os.Exit(m.Run())
}
