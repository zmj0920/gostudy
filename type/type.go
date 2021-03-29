package main

import "fmt"

type MyInt int64

func main() {
	// Go 语⾔言不不允许隐式类型转换
	// str := "hello world"
	// fmt.Println("str的长度" + len(str))

	// var a = 1
	// var b int32
	// b = a

	// 别名也是不支持隐式类型转换

	// var str = 1
	// var c MyInt
	// c = str
	// Point()
	TypeString()
}

func Implicit() {
	//go 不支持隐士类型转换
	//var a int = 1
	var a int32 = 1
	var b int64
	// 普通类型通过显示类型转换，int/float 可以使用 type(var) 的形式来进行强制类型转换
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	fmt.Print(a, b, c)
}

func Point() {
	a := 1
	aPtr := &a
	//go 不支持指针运算
	// aPtr = aPtr + 1
	fmt.Print(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
}

func TypeString() {
	var s string
	// var a *int
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
	if s == "" {
		fmt.Print("s is \"\"")
	}

	// 不使用显式类型，无法使用“nil”来初始化变量
	// if a == nil {
	// 	fmt.Print("s is nil")
	// }
}
