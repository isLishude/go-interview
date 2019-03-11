// +build ignore
package main

import (
	"fmt"
)

func fn() (closed bool) {
	defer func() {
		closed = true
	}()
	return false
}

func main() {
	res := fn()
	fmt.Println(res)
}

// Answer:
// true
