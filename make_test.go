package sh

import (
	"strings"
	"testing"
)

func TestMake(t *testing.T) {
	SetXtrace(false)

	Rm("tmp")

	targets := FromFields("aaa.o bbb.o ccc.o ddd.o").Sed(func(s string) string {
		return "tmp/target/" + s
	}).Slice()

	sharePrerequisites := FromFields("eee.h fff.h").Sed(func(s string) string {
		return "tmp/source/" + s
	}).Slice()

	substFn := func(s string) string {
		s = strings.Replace(s, "tmp/target", "tmp/source", 1)
		s = strings.Replace(s, ".o", ".c", 1)
		return s
	}

	prerequisites := FromSlice(targets).Sed(substFn).Slice()
	prerequisites = append(prerequisites, sharePrerequisites...)

	mk := NewMakeFile()
	mk.Task("all").Phony().Depend("prepare").Depend(targets...).Command(func(auto *MakeAuto) bool {
		Echo("done").Print()
		return true
	})
	mk.Task("prepare").Phony().Command(func(auto *MakeAuto) bool {
		Echo("prepare").Print()
		return true
	})

	mk.Tasks(targets...).ForEach(func(task *MakeTask) {
		task.Depend(substFn(task.Target())).Depend(sharePrerequisites...).Command(func(auto *MakeAuto) bool {
			Echof("generate %s from %s", auto.Target, auto.FirstPrerequisite).Print()
			Cp(auto.FirstPrerequisite, auto.Target)
			return true
		})
	})

	mk.Tasks(prerequisites...).ForEach(func(task *MakeTask) {
		task.Command(func(auto *MakeAuto) bool {
			Echof("generate %s", auto.Target).Print()
			Touch(auto.Target)
			return true
		})
	})

	// mk.Task(targets...).Depend(prerequisites...).Command(func(auto *MakeAuto) bool {
	// 	Cp(mappingFn(auto.Target), auto.Target)
	// 	return true
	// })

	// mk.Task(prerequisites...).Command(func(auto *MakeAuto) bool {
	// 	Touch(auto.Target)
	// 	return true
	// })

	Echo("First make").Print()
	mk.Make()

	Echo("Second make").Print()
	Rm("tmp/source/aaa.c")
	Rm("tmp/target/bbb.o")
	mk.Make()

	Echo("Third make").Print()
	Touch("tmp/source/ccc.c")
	Touch("tmp/target/ddd.o")
	mk.Make()

	Echo("Fouth make").Print()
	Rm("tmp/source/eee.h")
	mk.Make()

	Echo("Fiveth make").Print()
	Rm("tmp/source/aaa.c")
	Touch("tmp/source/fff.h")
	mk.Make()
}
