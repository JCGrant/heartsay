package renderer

import (
	"strings"
)

type Frame struct {
	rows []string
}

func NewFrame(frameStr string) *Frame {
	return &Frame{
		strings.Split(frameStr, "\n"),
	}
}

func (f *Frame) String() string {
	return strings.Join([]string(f.rows), "\n")
}
