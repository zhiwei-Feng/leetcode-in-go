package main

import (
	"C"
)

func main() {
	var ch chan struct{}
	<-ch
}
