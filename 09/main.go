package main

// see spec: https://golang.org/ref/spec#Assignability
// x's type V and T have identical underlying types and at least one of V or T is not a defined type.

func main() {
	type hash []byte

	var a hash
	var b []byte

	// right here?
	a = b
	b = a

	// right here?
	a = make([]byte, 0, 0)

	type number float64
	var i number
	var j float64

	// right here?
	// i = j
	// j = i

	// righer here?
	i = number(j)
	j = float64(i)
}
