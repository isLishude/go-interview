# go-interview

## 1

```go
package main

import (
	"fmt"
)

func fn() (closed bool) {
	defer func() {
		closed = true
	}()
	return false
}

func main() {
	res := fn()
	fmt.Println(res)
}

// Answer:
// true
```

## 2

```go
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

// Answer:
// init: true
// main: false
```

## 3

```go
package main

import (
	"fmt"
)

func main() {
	var a []int
	b := a[:]
	fmt.Println(b == nil)
}

// Answer:
// true
```

## 4

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() int {
	defer func() interface{} {
		if err := recover(); err != nil {
			return err
		}
		return nil
	}()
	panic(1)
}

// What will be?
// - Compile failed
// - print 0
// - nothing print

// Answer
// print 0
```

## 5

```go
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
	a.Lock()
	b.Lock()
}

// None
// @see https://golang.org/ref/spec#Type_declarations
```

## 6

```go
package main

// see spec: https://golang.org/ref/spec#Assignability
// x's type V and T have identical underlying types and at least one of V or T is not a defined type.

/*
可将类型分为命名和未命名两⼤֒类。

命名类型包括 bool、int、string 等，⽽ array、slice、map 等和具体元素类型、⻓度等有关，属于未命名类型。
具有相同声明的未命名类型被视为同⼀类型。
- 具有相同基类型的指针。
- 具有相同元素类型和⻓度的 array
- 具有相同元素类型的 slice
- 具有相同键值类型的 map
- 具有相同元素类型和传送⽅ٛ向的 channel
- 具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct
- 签名相同 (参数和返回值，不包括参数名称) 的 function
- ⽅ٛ法集相同 (⽅ٛ法名、⽅ٛ法签名相同，和次序⽆ٞ关) 的 interface

                                 -- Go语言学习笔记
*/

type hash []byte
type number float64

var a hash
var b []byte
var i number
var j float64

func main() {

	// right here?
	a = b
	b = a

	// right here?
	a = make([]byte, 0)

	// right here?
	a = make(hash, 0)

	// why not correct here?
	// i = j
	// j = i

	// righer here?
	i = number(j)
	j = float64(i)
}
```

## 7

```go
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
```

## 8

```go
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

// hello,world
```

## 9

```go
package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = 1
	e = iota
	f
	g = 1
	h
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h)
}
// 0 1 2 1 1 2 1 1
```
