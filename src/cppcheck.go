package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type cppcheck struct{}

func (m cppcheck) parse(filename string) []warningInfo {
	data, err := ioutil.ReadFile(filename)

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
			file:    m[1],
			lines:   m[2] + "-" + m[3],
			kind:    m[4],
			message: m[5],
		})
	}
	return warns
}
