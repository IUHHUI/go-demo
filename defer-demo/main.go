package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	for i := 0; i < 3; i++ {
		//defer 函数的参数 在 defer 时就已经计算完成并确定
		defer fmt.Printf("%d ", i)
	}

	defer fmt.Println()

	defer func() {
		//使用参数在匿名函数内部,打印时,还是3 4 5
		for i := 3; i < 6; i++ {
			fmt.Printf("%d ", i)
		}
	}()

	b()

	/*
		entering: b
		in b
		entering: a
		in a
		leaving: a
		leaving: b
		3 4 5
		2 1 0
	*/
}
