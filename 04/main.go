// +build ignore
package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() int {
	defer func() interface{} {
		if err := recover(); err != nil {
			return err
		}
		return nil
	}()
	panic(1)
}

// What will be?
// - Compile failed
// - print 0
// - nothing print

// Answer
// print 0
