// +build ignore
package main

import (
	"fmt"
)

// panic or print nil or print 10?
func main() {
	c := make(chan int, 1)
	c <- 10
	close(c)

	fmt.Println(<-c)
}

// print 10
