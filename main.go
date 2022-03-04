package main

import (
	"fmt"
	"time"
)

var a int

func main() {
	go func() {
		a = 1
	}()
	go func() {
		a = 2
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}
