package main

import "fmt"

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Printf("uint %d\n", a-b)

	var c uint8 = 1
	var d uint8 = 2
	fmt.Printf("uint %d\n", c-d)

	var e uint16 = 1
	var f uint16 = 2
	fmt.Printf("uint %d\n", e-f)

	var g uint32 = 1
	var h uint32 = 2
	fmt.Printf("uint %d\n", g-h)

	var i uint64 = 1
	var j uint64 = 2
	fmt.Printf("uint %d\n", i-j)

	/**
	uint 18446744073709551615
	uint 255
	uint 65535
	uint 4294967295
	uint 18446744073709551615
	**/

}
