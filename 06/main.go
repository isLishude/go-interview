// +build ignore
package main

import (
	"fmt"
)

func test1() {
	s := []int{1, 2, 3}
	ss := s[1:]

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

func test2() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

func main() {
	test1()
	test2()
}

// Answer:
// [1 12 13]
// [1 2 3]
