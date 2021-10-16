// Code generated by 'yaegi extract path'. DO NOT EDIT.

//go:build go1.17
// +build go1.17

package sgolib

import (
	"path"
	"reflect"
)

func init() {
	Symbols["path/path"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Base":          reflect.ValueOf(path.Base),
		"Clean":         reflect.ValueOf(path.Clean),
		"Dir":           reflect.ValueOf(path.Dir),
		"ErrBadPattern": reflect.ValueOf(&path.ErrBadPattern).Elem(),
		"Ext":           reflect.ValueOf(path.Ext),
		"IsAbs":         reflect.ValueOf(path.IsAbs),
		"Join":          reflect.ValueOf(path.Join),
		"Match":         reflect.ValueOf(path.Match),
		"Split":         reflect.ValueOf(path.Split),
	}
}
