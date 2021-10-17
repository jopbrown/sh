package sh

import (
	"strings"
	"testing"
)

func TestMake(t *testing.T) {
	Rm("tmp")

	targets := FromFields("1.txt 2.txt 3.txt").Sed(func(s string) string {
		return "tmp/target/" + s
	}).Slice()

	mappingFn := func(s string) string {
		s = strings.Replace(s, "tmp/target", "tmp/source", 1)
		s = strings.Replace(s, ".txt", ".src", 1)
		return s
	}

	prerequisites := FromSlice(targets).Sed(mappingFn).Slice()

	Touch(prerequisites...)

	Make(targets, prerequisites, func(target string, prerequisites []string) bool {
		Cp(mappingFn(target), target)
		return true
	})
}
