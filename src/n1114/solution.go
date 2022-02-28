package n1114

import (
	"fmt"
	"sync"
)

type Foo struct {
	firstReady  chan int
	secondReady chan int
}

func (f Foo) First() {
	fmt.Print("first")
	f.firstReady <- 1
}
func (f Foo) Second() {
	<-f.firstReady
	fmt.Print("second")
	f.secondReady <- 1
}
func (f Foo) Third() {
	<-f.secondReady
	fmt.Print("third")
}

func print(nums []int) {
	foo := Foo{
		firstReady: make(chan int),
		secondReady: make(chan int),
	}
	var wg sync.WaitGroup
	wg.Add(3)
	for _, v := range nums {
		switch v {
		case 1:
			go func() {
				foo.First()
				wg.Done()
			}()
		case 2:
			go func() {
				foo.Second()
				wg.Done()
			}()
		case 3:
			go func() {
				foo.Third()
				wg.Done()
			}()
		}
	}

	wg.Wait()
}
