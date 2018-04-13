package main

import (
	"fmt"
)

func main() {
	var m module

	var cpp cppcheck

	m = cpp

	warns := m.parse("../cppcheck.log")

	for _, w := range warns {
		fmt.Println(w)
	}
}
