package main

import (
	"fmt"
)

func main() {
	res := test()
  // What will print?
	fmt.Println(res)
}

func test() (closed bool) {
	defer func() {
		closed = true
	}()
	return false
}
