package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic(1)
}

// What will be?
// - Compile failed
// - print 0
// - nothing print
