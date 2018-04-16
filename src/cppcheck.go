package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type cppcheck struct{}

func (m cppcheck) warnings() []warningInfo {
	data, err := ioutil.ReadFile(cfg.CppCheck.InputFile)

	if err != nil {
		log.Fatal(err)
	}

	text := string(data)

	lines := strings.Split(text, "\n")

	//TODO: fix regexp (now it doesn't handle strings with one line)
	// like [/home/fexolm/git/SLama/src/block.c:13]: (portability) 'mem' is of type 'void *'. When using void pointers in calculations, the behaviour is undefined.

	r := regexp.MustCompile(`\[([/\w-.]+):(\d+)\].*\[.*:(\d+)\]:\s\((\w+)\)\s(.*)`)

	warns := make([]warningInfo, 0, len(lines))

	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		warns = append(warns, warningInfo{
			File:    m[1],
			Lines:   m[2] + "-" + m[3],
			Kind:    m[4],
			Message: m[5],
			Module:  "cppcheck",
		})
	}
	return warns
}
