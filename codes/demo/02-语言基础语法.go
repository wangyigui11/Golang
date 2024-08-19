// Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。

// 一行代表一个语句的结束。
// 如果你打算将多个语句写在同一行，那么你需要使用; 来分隔它们。实际开发中一般不这样使用

// 注释不会被编译， 每一个包应该有相应的注释

//

package main

import "fmt"

func main() {

	if x > 0 {
		// do something
	}

	// result := add(2, 3)

	fmt.Println("hello world")
	fmt.Println("我是第二个语句")

	// hello world
	// 我是第二个语句
}

/**
	标识符 用来命名变量、类型的等程序实体。 一个标识符实际上就是一个或者多个字母 (A~Z和a~z)数字(0~9)、下划线_ 组成的序列，
	但是第一个字符必须是字母或斜划线而不能是数字

	以下是有效的标识符：
	mahesh   kumar   abc   move_name   a_123
	myname50   _temp   j   a23b9   retVal

	以下是无效的标识符：

	1ab（以数字开头）
	case（Go 语言的关键字）
	a+b（运算符是不允许的）
**/

// 字符串连接

// Go的字符串连接可以通过 + 实现  fmt.Println("YoStar" + "666")  输出  YoStar666

// 关键字
// 下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

// break	default	func	interface	select
// case	defer	go	map	struct
// chan	else	goto	package	switch
// const	fallthrough	if	range	type
// continue	for	import	return	var

// 除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

// append	bool	byte	cap	close	complex	complex64	complex128	uint16
// copy	false	float32	float64	imag	int	int8	int16	uint32
// int32	int64	iota	len	make	new	nil	panic	uint64
// print	println	real	recover	string	true	uint	uint8	uintptr

// 程序一般由关键字、常量、变量、运算符、类型和函数组成。

// 程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

// 程序中可能会使用到这些标点符号：.、,、;、: 和 …。
