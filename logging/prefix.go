package logging

import "fmt"

type prefixLogger struct {
	Name string
}

func (pf *prefixLogger) Log(level Level, format string, args ...interface{}) {
	pf.log(level, format, args...)
}

func (pf *prefixLogger) Info(format string, args ...interface{}) {
	pf.Log(Info, format, args...)
}

func (pf *prefixLogger) Warning(format string, args ...interface{}) {
	pf.Log(Warning, format, args...)
}

func (pf *prefixLogger) Error(format string, args ...interface{}) {
	pf.Log(Error, format, args...)
}

func (pf *prefixLogger) Critical(format string, args ...interface{}) {
	pf.Log(Critical, format, args...)
}

func (pf *prefixLogger) Debug(format string, args ...interface{}) {
	pf.Log(Debug, format, args...)
}

func (pf *prefixLogger) log(level Level, format string, args ...interface{}) {
	fmt.Printf("["+pf.Name+"] ["+level.Name()+"] "+format+"\n", args...)
}
