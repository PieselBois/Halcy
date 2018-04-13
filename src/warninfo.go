package main

import "fmt"

type warningInfo struct {
	kind    string
	file    string
	lines   string
	message string
}

func (w warningInfo) String() string {
	return fmt.Sprintf("(%s) %s [%s]: %s", w.kind, w.file, w.lines, w.message)
}
