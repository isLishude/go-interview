// +build ignore
package main

import (
	"fmt"
)

func main() {
	var a []int
	b := a[:]
	fmt.Println(b == nil)
}

// Answer:
// true
