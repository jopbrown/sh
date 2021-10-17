package sh

import (
	"regexp"
)

func Grep(pattern string, fileList ...string) Stream {
	return Cat(Ls(fileList...).Slice()...).Grep(pattern)
}

func (sin Stream) Grep(pattern string) Stream {
	xtrace("grep '%s'", pattern)
	re, err := regexp.Compile(pattern)
	if CheckErr(err) {
		return nil
	}

	return yieldStream(func(sout Stream) {
		for line := range sin {
			if re.MatchString(line) {
				sout <- line
			}
		}
	})
}

func GrepV(pattern string, fileList ...string) Stream {
	return Cat(Ls(fileList...).Slice()...).GrepV(pattern)
}

func (sin Stream) GrepV(pattern string) Stream {
	xtrace("grep -v '%s'", pattern)
	re, err := regexp.Compile(pattern)
	if CheckErr(err) {
		return nil
	}

	return yieldStream(func(sout Stream) {
		for line := range sin {
			if !re.MatchString(line) {
				sout <- line
			}
		}
	})
}

func Sed(fn func(string) string, fileList ...string) Stream {
	return Cat(fileList...).Sed(fn)
}

func (sin Stream) Sed(fn func(string) string) Stream {
	return yieldStream(func(sout Stream) {
		for str := range sin {
			sout <- fn(str)
		}
	})
}
