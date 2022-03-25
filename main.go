package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)


func main() {
	uint
	file, err := os.Open("a.txt")
	if err!=nil {
		panic("")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}
