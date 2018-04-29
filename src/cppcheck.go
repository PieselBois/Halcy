package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type cppcheck struct{}

func (m cppcheck) warnings() []warningInfo {
	// data, err := ioutil.ReadFile(cfg.CppCheck.InputFile)

	c := exec.Command("cppcheck",
		fmt.Sprintf("--project=%s", cfg.CppCheck.CompileCommands),
		"--enable=all", "-q")
	out, err := c.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	text := string(out)

	lines := strings.Split(text, "\n")

	r1 := regexp.MustCompile(
		`\[([/\w-.]+):(\d+)\].*\[.*:(\d+)\]:\s\((\w+)\)\s(.*)`)
	r2 := regexp.MustCompile(
		`\[([/\w-.]+):(\d+)\]:\s\((\w+)\)\s(.*)`)

	warns := make([]warningInfo, 0, len(lines))

	for _, line := range lines {

		if m := r1.FindStringSubmatch(line); m != nil {
			warns = append(warns, warningInfo{
				File:    m[1],
				Lines:   m[2] + "-" + m[3],
				Kind:    m[4],
				Message: m[5],
				Module:  "cppcheck",
			})
		} else if m = r2.FindStringSubmatch(line); m != nil {
			warns = append(warns, warningInfo{
				File:    m[1],
				Lines:   m[2],
				Kind:    m[3],
				Message: m[4],
				Module:  "cppcheck",
			})
		}
	}
	return warns
}
