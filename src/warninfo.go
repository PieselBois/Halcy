package main

import (
	"fmt"
)

type warningInfo struct {
	Kind    string `json:"kind"`
	File    string `json:"file"`
	Lines   string `json:"lines"`
	Message string `json:"message"`
	Module  string `json:"module"`
}

func (w warningInfo) String() string {
	return fmt.Sprintf("(%s) %s [%s]: %s", w.Kind, w.File, w.Lines, w.Message)
}
