// Code generated by 'yaegi extract github.com/jopbrown/sh'. DO NOT EDIT.

package sgolib

import (
	"github.com/jopbrown/sh"
	"reflect"
)

func init() {
	Symbols["github.com/jopbrown/sh/sh"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Arg":            reflect.ValueOf(sh.Arg),
		"ArgN":           reflect.ValueOf(sh.ArgN),
		"Args":           reflect.ValueOf(sh.Args),
		"Cat":            reflect.ValueOf(sh.Cat),
		"Cd":             reflect.ValueOf(sh.Cd),
		"CheckErr":       reflect.ValueOf(sh.CheckErr),
		"Chmod":          reflect.ValueOf(sh.Chmod),
		"Cp":             reflect.ValueOf(sh.Cp),
		"Echo":           reflect.ValueOf(sh.Echo),
		"Echof":          reflect.ValueOf(sh.Echof),
		"Err":            reflect.ValueOf(sh.Err),
		"Exec":           reflect.ValueOf(sh.Exec),
		"Exec2":          reflect.ValueOf(sh.Exec2),
		"Exists":         reflect.ValueOf(sh.Exists),
		"ExistsDir":      reflect.ValueOf(sh.ExistsDir),
		"ExistsFile":     reflect.ValueOf(sh.ExistsFile),
		"Exit":           reflect.ValueOf(sh.Exit),
		"ExitIfErr":      reflect.ValueOf(sh.ExitIfErr),
		"From":           reflect.ValueOf(sh.From),
		"FromFields":     reflect.ValueOf(sh.FromFields),
		"FromReader":     reflect.ValueOf(sh.FromReader),
		"FromSlice":      reflect.ValueOf(sh.FromSlice),
		"Grep":           reflect.ValueOf(sh.Grep),
		"GrepV":          reflect.ValueOf(sh.GrepV),
		"Ls":             reflect.ValueOf(sh.Ls),
		"Mkdir":          reflect.ValueOf(sh.Mkdir),
		"Mv":             reflect.ValueOf(sh.Mv),
		"NewMakeFile":    reflect.ValueOf(sh.NewMakeFile),
		"Popd":           reflect.ValueOf(sh.Popd),
		"Pushd":          reflect.ValueOf(sh.Pushd),
		"Pwd":            reflect.ValueOf(sh.Pwd),
		"Rm":             reflect.ValueOf(sh.Rm),
		"Sed":            reflect.ValueOf(sh.Sed),
		"SetExitOnError": reflect.ValueOf(sh.SetExitOnError),
		"SetXtrace":      reflect.ValueOf(sh.SetXtrace),
		"Touch":          reflect.ValueOf(sh.Touch),

		// type definitions
		"MakeAuto":  reflect.ValueOf((*sh.MakeAuto)(nil)),
		"MakeFile":  reflect.ValueOf((*sh.MakeFile)(nil)),
		"MakeFunc":  reflect.ValueOf((*sh.MakeFunc)(nil)),
		"MakeTask":  reflect.ValueOf((*sh.MakeTask)(nil)),
		"MakeTasks": reflect.ValueOf((*sh.MakeTasks)(nil)),
		"Stream":    reflect.ValueOf((*sh.Stream)(nil)),
	}
}
