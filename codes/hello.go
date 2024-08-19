package main

import "fmt"

func main() {

	// 这是单行注释
	/**
	 这是多行注释
	 第一行
	 第二行
	**/
	// fmt.Println("Hello World!")
	// fmt.Println("菜鸟教程：runoob.com")

	// fmt.Println("Google" + "Runoob")

	// // %d 表示整型数字，%s 表示字符串
	// var stockcode = 123
	// var enddate = "2020-12-31"
	// var url = "Code=%d&endDate=%s"
	// fmt.Printf(url, stockcode, enddate)

	// 语言变量
	// var a string = "Runoob"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// 输出结果为
	// Runoob
	// 1 2

	/**

		// 语言变量
		// var identifier type

		// var identifier1, identifier2 type

		var a string = "Runoob"

		fmt.Println(a)

		var b, c int = 1, 2
		fmt.Println(b, c)

		// Runoob
		// 1 2
	**/

	// 变量声明

	// 第一种， 指定变量类型， 如果没有初始化， 则变量默认为零值

	// 数值类型为0

	// 布尔类型

	// var v_name v_type

	// v_name = value

	var a = "RUNOOB"

	fmt.Println(a) // RUNOOB

	// 没有初始化就为零值
	var b int

	fmt.Println(b) // 0

	var c bool

	fmt.Println(c) // false
}
