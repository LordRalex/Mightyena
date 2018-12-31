package logging

import "fmt"

type prefixLogger struct {
	Name string
}

func (pf *prefixLogger) Print(msg string) {
	pf.print(msg)
}

func (pf *prefixLogger) Printf(format string, args ...interface{}) {
	pf.print(format, args...)
}

func (pf *prefixLogger) print(format string, args ...interface{}) {
	fmt.Printf("[" + pf.Name + "]" + format + "\n", args...)
}