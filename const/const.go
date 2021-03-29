package main

import "fmt"

//连续常量赋值
// iota 表示取值为0(const关键字）所在行，后续常量依次加1
func main() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	test2()
}

func test2() {
	const (
		Readable = 1 << iota
		Writable
		Executable
	)
	// 0001 0010 0100 即 1 2 4
	fmt.Println(Readable, Writable, Executable)
	// 0111 即 7,表示可读,可写,可执行
	accessCode := 7
	fmt.Println(accessCode&Readable == Readable, accessCode&Writable == Writable, accessCode&Executable == Executable)
}
