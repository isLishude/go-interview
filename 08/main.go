package main

var x interface{}

func f() (interface{}, interface{}) {
	return nil, nil
}

// Assuming x is declared and y is not declared, which clauses below are correct?
func main() {
	// x, _ := f()
	// x, _ = f()
	// x, y := f()
	// x, y = f()
}
