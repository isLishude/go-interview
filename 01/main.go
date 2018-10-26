package main

import (
	"fmt"
	"log"
)

var (
	todo bool
)

func init() {
	todo, err := fn()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("init:", todo)
}

func fn() (bool, error) {
	return true, nil
}

func main() {
	fmt.Println("main:", todo)
}
