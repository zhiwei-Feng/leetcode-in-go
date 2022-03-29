package main

import (
	"fmt"
)

func main() {
	mPtr := new([]int)
	*mPtr = append(*mPtr, 1)
	fmt.Println(*mPtr)
}
