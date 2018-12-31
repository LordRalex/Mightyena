package logging

import "fmt"

type prefixLogger struct {
	Name string
}

func (pf *prefixLogger) Log(level Level, format string, args ...interface{}) {
	pf.log(level, format, args...)
}

func (pf *prefixLogger) log(level Level, format string, args ...interface{}) {
	fmt.Printf("[" + pf.Name + "] [" + level.Name() + "] " + format + "\n", args...)
}