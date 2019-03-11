// +build ignore
package main

type m struct{}

func (m *m) Lock()   {}
func (m *m) Unlock() {}

type ptr *m
type newM m

var a ptr
var b newM

// which code section can compile success?
func main() {
	// Why none of them can run ?
	// a.Lock()
	// b.Lock()
}

// None
// @see https://golang.org/ref/spec#Type_declarations
