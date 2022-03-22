package main

import (
	"fmt"
	"runtime"
	"strings"
)

func New(format string, args ...interface{}) error {
	pcList := pcs()
	callersFrames := runtime.CallersFrames(pcList)

	var frames []*frameInfo

	for {
		f, ok := callersFrames.Next()
		if !ok {
			break
		}
		frames = append(frames, &frameInfo{
			file: f.File,
			name: f.Function,
			line: f.Line,
		})
	}

	return &BoringError{
		msg:    fmt.Sprintf(format, args...),
		frames: frames,
	}
}

func pcs() []uintptr {
	buf := make([]uintptr, 127)
	for {
		n := runtime.Callers(3, buf)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]uintptr, 2*len(buf))
	}
}

type BoringError struct {
	msg    string
	frames []*frameInfo
}

func (be *BoringError) Error() string {
	buf := strings.Builder{}
	buf.WriteString(be.msg)
	buf.WriteString("\n")
	for _, f := range be.frames {
		buf.WriteString(f.String())
		buf.WriteString("\n")
	}
	return buf.String()
}

type frameInfo struct {
	name string
	file string
	line int
}

func (fi *frameInfo) String() string {
	return fmt.Sprintf("%s %s:%d", fi.name, fi.file, fi.line)
}

func foo() error {
	return bar()
}

func bar() error {
	return New("foo")
}

func main() {
	err := foo()
	berr := err.(*BoringError)
	fmt.Println(berr)
}
