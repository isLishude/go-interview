package main

// this section is from `https://github.com/goquiz/goquiz.github.io`
import "fmt"

type st struct{}

func (s st) F() {}

type intF interface {
	F()
}

func initType() st {
	var s st
	return s
}

func initPointer() *st {
	var s *st
	return s
}
func initEfaceType() interface{} {
	var s st
	return s
}

func initEfacePointer() interface{} {
	var s *st
	return s
}

func initIntFType() intF {
	var s st
	return s
}

func initIntFPointer() intF {
	var s *st
	return s
}

func main() {
	// What will be printed when the code below is executed?
	// And explain why fmt.Println(InitType() == nil) cannot be compiled?
	// println(InitType() == nil)
	fmt.Println(initPointer() == nil)
	fmt.Println(initEfaceType() == nil)
	fmt.Println(initEfacePointer() == nil)
	fmt.Println(initIntFType() == nil)
	fmt.Println(initIntFPointer() == nil)
}
