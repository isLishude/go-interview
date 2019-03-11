// +build ignore
package main

var x interface{}

func f() (interface{}, interface{}) {
	return nil, nil
}

// Assuming x is declared and y is not declared, which clauses below are correct?
func main() {
	// x, _ := f() //incorrect
	x, _ = f()  //correct
	x, y := f() //correct
	x, y = f()  //incorrect
}
