package main

type module interface {
	parse(filename string) []warningInfo
}
