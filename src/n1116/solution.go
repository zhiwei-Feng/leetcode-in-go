package n1116

import (
	"fmt"
	"sync"
)

type ZeroEvenOdd struct {
	N                              int
	ZeroReady, OddReady, EvenReady chan int
}

func (this ZeroEvenOdd) zero() {
	for i := 0; i < this.N; i++ {
		<-this.ZeroReady
		fmt.Print("0")
		if i%2 == 0 {
			this.OddReady <- 1
		} else {
			this.EvenReady <- 1
		}
	}
}

func (this ZeroEvenOdd) even() {
	for i := 0; i < this.N/2; i++ {
		<-this.EvenReady
		fmt.Print(2*i + 2)
		if this.N%2 == 1 || i < this.N/2-1 {
			this.ZeroReady <- 1
		}
	}
}

func (this ZeroEvenOdd) odd() {
	for i := 0; i < (this.N+1)/2; i++ {
		<-this.OddReady
		fmt.Print(2*i + 1)
		if this.N%2 == 0 || i < (this.N+1)/2-1 {
			this.ZeroReady <- 1
		}
	}
}

func Print(n int) {
	printer := ZeroEvenOdd{
		N:         n,
		ZeroReady: make(chan int),
		OddReady:  make(chan int),
		EvenReady: make(chan int),
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		printer.zero()
		wg.Done()
	}()
	go func() {
		printer.odd()
		wg.Done()
	}()
	go func() {
		printer.even()
		wg.Done()
	}()
	printer.ZeroReady <- 1
	wg.Wait()
}
