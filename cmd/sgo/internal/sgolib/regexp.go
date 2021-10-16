// Code generated by 'yaegi extract regexp'. DO NOT EDIT.

//go:build go1.17
// +build go1.17

package sgolib

import (
	"reflect"
	"regexp"
)

func init() {
	Symbols["regexp/regexp"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Compile":          reflect.ValueOf(regexp.Compile),
		"CompilePOSIX":     reflect.ValueOf(regexp.CompilePOSIX),
		"Match":            reflect.ValueOf(regexp.Match),
		"MatchReader":      reflect.ValueOf(regexp.MatchReader),
		"MatchString":      reflect.ValueOf(regexp.MatchString),
		"MustCompile":      reflect.ValueOf(regexp.MustCompile),
		"MustCompilePOSIX": reflect.ValueOf(regexp.MustCompilePOSIX),
		"QuoteMeta":        reflect.ValueOf(regexp.QuoteMeta),

		// type definitions
		"Regexp": reflect.ValueOf((*regexp.Regexp)(nil)),
	}
}
