// +build ignore
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
