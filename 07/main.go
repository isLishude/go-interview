package main

type m struct{}

func (m *m) Lock()   {}
func (m *m) Unlock() {}

type ptr *m
type newM m

// which code section can compile success?
func main() {
	/*
		var a ptr
		a.Lock()

		var b newM
		b.Lock()
	*/
}
