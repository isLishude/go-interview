// +build ignore
package main

import (
	"fmt"
	"strings"
)

func greet(names ...string) {
	names[0] = "hello"
	names[1] = "world"
}

func main() {
	params := []string{"init", "value"}
	greet(params...)
	res := strings.Join(params, ",")

	// What will print?
	fmt.Println(res)
}
