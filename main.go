package main

import (
	"fmt"
	"time"
)

func printA(c chan int, a chan int) {
	for range c {
		fmt.Print("A")
		a <- 1
	}
}

func printB(a chan int, b chan int) {
	for range a {
		fmt.Print("B")
		b <- 1
	}
}

func printC(b chan int, c chan int) {
	for range b {
		fmt.Print("C")
		c <- 1
	}
}

func main() {
	cha := make(chan int)
	chb := make(chan int)
	chc := make(chan int)
	go printA(chc, cha)
	go printB(cha, chb)
	go printC(chb, chc)
	chc <- 1
	time.Sleep(3 * time.Second)
}
