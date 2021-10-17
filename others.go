package sh

import (
	"fmt"
	"os"
)

type MakeFunc func(target string, prerequisites []string) bool

func Make(targets []string, prerequisites []string, fn MakeFunc) {

targetLoop:
	for _, target := range targets {
		tinfo, err := os.Stat(target)
		needMake := false
		if err != nil {
			needMake = true
		} else {
			for _, pre := range prerequisites {
				pinfo, err := os.Stat(pre)
				if err != nil {
					CheckErr(fmt.Errorf("make: prerequisite not exist: %s", pre))
					break targetLoop
				}

				if tinfo.ModTime().Before(pinfo.ModTime()) {
					needMake = true
					break
				}
			}
		}

		if needMake {
			if !fn(target, prerequisites) {
				break targetLoop
			}
		}
	}
}
