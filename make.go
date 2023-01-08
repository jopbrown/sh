package sh

import (
	"fmt"
	"os"
)

type MakeFile struct {
	firstTarget string
	targetDict  map[string]*MakeTask
}

type MakeFunc func(auto *MakeAuto) bool

type MakeTask struct {
	target        string
	prerequisites []string
	isPhony       bool
	fn            MakeFunc
}

type MakeTasks []*MakeTask

type MakeAuto struct {
	Target            string
	FirstPrerequisite string
	Prerequisites     []string
}

func NewMakeFile() *MakeFile {
	mk := &MakeFile{}
	mk.targetDict = make(map[string]*MakeTask)
	return mk
}

func (mk *MakeFile) Tasks(targets ...string) MakeTasks {
	if len(targets) == 0 {
		panic("at least one target for task")
	}

	tasks := make(MakeTasks, 0, len(targets))

	for _, target := range targets {
		tasks = append(tasks, mk.Task(target))
	}

	return tasks
}

func (mk *MakeFile) Task(target string) *MakeTask {
	task := &MakeTask{}
	task.target = target
	task.prerequisites = make([]string, 0)
	mk.targetDict[target] = task

	if mk.firstTarget == "" {
		mk.firstTarget = target
	}

	return task
}

func (tasks MakeTasks) ForEach(fn func(task *MakeTask)) {
	for _, task := range tasks {
		fn(task)
	}
}

func (task *MakeTask) Target() string {
	return task.target
}

func (task *MakeTask) Depend(prerequisites ...string) *MakeTask {
	task.prerequisites = append(task.prerequisites, prerequisites...)
	return task
}

func (task *MakeTask) Phony() *MakeTask {
	task.isPhony = true
	return task
}

func (task *MakeTask) Command(fn MakeFunc) *MakeTask {
	task.fn = fn
	return task
}

func (mk *MakeFile) Make(targets ...string) {
	if len(targets) == 0 {
		CheckErr(mk.make(mk.firstTarget))
	}

	for _, target := range targets {
		if CheckErr(mk.make(target)) {
			return
		}
	}
}

func (mk *MakeFile) make(target string) error {
	task, found := mk.targetDict[target]
	if !found {
		return fmt.Errorf("not found task of target: %s", target)
	}

	needRunCommand := false
	if task.isPhony || !ExistsFile(target) {
		needRunCommand = true
	}

	for _, pre := range task.prerequisites {
		preTask, prefound := mk.targetDict[pre]
		if !prefound && !ExistsFile(pre) {
			return fmt.Errorf("not task to make prerequisite of %s: %s", target, pre)
		}

		if prefound {
			err := mk.make(pre)
			if err != nil {
				return fmt.Errorf("fail to make prerequisite of %s: %w", target, err)
			}
		}

		if needRunCommand {
			continue
		}

		if prefound && preTask.isPhony {
			continue
		}

		tinfo, _ := os.Stat(target)
		pinfo, err := os.Stat(pre)
		if err != nil {
			return fmt.Errorf("make: prerequisite not exist: %s", pre)
		}

		if tinfo.ModTime().Before(pinfo.ModTime()) {
			needRunCommand = true
		}
	}

	if needRunCommand {
		if task.fn != nil {
			auto := &MakeAuto{}
			auto.Target = target
			auto.Prerequisites = task.prerequisites
			if len(auto.Prerequisites) > 0 {
				auto.FirstPrerequisite = auto.Prerequisites[0]
			}
			if !task.fn(auto) {
				return fmt.Errorf("task command execute fail: %s", target)
			}
		}

		if !task.isPhony && !ExistsFile(target) {
			return fmt.Errorf("task command not generate target: %s", target)
		}
	}

	return nil
}

func (mk *MakeFile) QueryTask(target string) *MakeTask {
	return mk.targetDict[target]
}
