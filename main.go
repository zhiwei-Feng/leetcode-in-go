package main

import (
	"sync"
)

var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func main() {
	go doprint()
	go doprint()
}
