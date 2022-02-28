package n1115

import (
	"fmt"
)

type FooBar struct {
	N        int
	FooReady chan int
	BarReady chan int
	End      chan int
}

func (this *FooBar) foo() {
	for i := 0; i < this.N; i++ {
		<-this.BarReady
		fmt.Print("foo")
		this.FooReady <- 1
	}
}

func (this *FooBar) bar() {
	for i := 0; i < this.N; i++ {
		<-this.FooReady
		fmt.Print("bar")
		if i < this.N-1 {
			this.BarReady <- 1
		}
	}
	this.End <- 1
}

func PrintAlternately(n int) {
	fb := FooBar{N: n, 
		FooReady: make(chan int), 
		BarReady: make(chan int), 
		End: make(chan int)}
	go func() {
		fb.foo()

	}()
	go func() {
		fb.bar()
	}()
	fb.BarReady <- 1
	<-fb.End
}
