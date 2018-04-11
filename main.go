package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// WarningInfo ...
type WarningInfo struct {
	wtype string
	file  string
	lines string
	logs  string
}

func main() {
	bytes, err := ioutil.ReadFile("cppcheck.log")

	if err != nil {
		fmt.Print(err)
	}

	fileText := string(bytes)

	lines := strings.Split(fileText, "\n")

	re := regexp.MustCompile(`\[([/\w-.]+):(\d+)\].*\[.*:(\d+)\]:\s\((\w+)\)\s(.*)`)

	warnings := make([]WarningInfo, len(lines))

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		var warning WarningInfo
		warning.file = match[1]
		warning.lines = match[2] + "-" + match[3]
		warning.wtype = match[4]
		warning.logs = match[5]
		warnings[i] = warning
	}

	for _, w := range warnings {
		fmt.Printf("(%s) %s [%s]: %s", w.wtype, w.file, w.lines, w.logs)
	}
}
