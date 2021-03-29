package main

import "fmt"

func main() {
	// var int a = 1
	// var int b = 1

	// var a = 1
	// var b = 1
	// var (
	// 	a = 1
	// 	b = 1
	// )
	a := 1
	b := 1

	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换赋值
func Exchange() {
	a := 1
	b := 1

	for i := 0; i < 6; i++ {
		fmt.Print(" ", b)
		b = a + b
		a, b = b, a
	}
}
