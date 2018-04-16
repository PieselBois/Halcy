package main

type module interface {
	warnings() []warningInfo
}
