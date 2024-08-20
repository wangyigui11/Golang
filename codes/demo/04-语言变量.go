// Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字
// 声明变量的一般形式是使用var关键字

// var identifier type

// 可以一次声明多个变量
// var identifier1, identifier2 type

// 变量声明
// 第一种，指定变量类型，如果没有初始化，则变量默认为零值
// var v_name v_type
// v_name = value

// 数值类型(包括complex64/128) 为0

// 布尔类型为false
// 字符串为“” (空字符串)
// 以下几种类型为 nil
// var a *int
// var a []int
// var a map[string] int
// var a chan int
// var a func(string) int
// var a error

// 第二种，根据值自行判定变量类型
// var d = true

// 第三种， 如果变量已经使用var声明过了， 再使用 := 声明变量， 就产生错误编译， 格式
// v_name := value

// 多变量声明
/****/
// var vname1, vname2, vname3 type

// vname1, vname2, vname3 = v1, v2, v3

// var vname1, vname2, vname3 = v1, v2, v3 // 和python比较相似， 不需要显示声明类型， 自动推断

// vname1, vname2, vname3 := v1, v2, v3 // 出现在:=左侧的变量不应该是已经被声明过的。否则会导致编译错误

// // 这种因式分解关键字的写法一般用于声明全局变量

// var (
// 	vname1 v_type1
// 	vname2 v_type2
// )

// 值类型和引用类型

// 所有像 int float bool string这些基本类型都属于值类型， 使用这些类型的变量直接指向存在内存中的值
// 当使用 = 将一个变量赋值给另一个变量时， 如： j = i, 实际上是在内存中将i的值进行了拷贝

// 更复杂的数据通畅会需要使用多个字， 这些数据一般使用引用类型保存。

package main

import "fmt"

func main() {
	// 声明一个变量并初始化
	var a string = "YoStar"

	fmt.Println(a) // YoStar

	var b, c int = 1, 2

	fmt.Println(b, c) // 1 2

	// 默认为零值
	var d int
	fmt.Println(d) // 0

	var e bool // false
	fmt.Println(e)

	var f *int
	fmt.Println(f) // <nil>

	// var intVal int
	// intVal := 1 // 这个时候会产生编译错误， 因为intVal 已经被声明过了， 不需要重新声明， 直接使用 intVal := 1即可

	intVal := 1
	fmt.Println(intVal) // 1
	// 相当于 var intVal
	// int intVal = 1

	g := "YoStar"
	fmt.Println(g) // YoStar

	var x, y int // 0 0
	var (
		aa int  // 0
		bb bool // false
	)

	var cc, dd int = 1, 2 // 1, 2

	var ee, ff = 123, "Hello" // 123 Hello

	// 这种不带声明格式的只能在函数体中出现
	gg, hh := 345, "Hello :=" // 345 Hello :=

	fmt.Println(x, y, aa, bb, cc, dd, ee, ff, gg, hh)
}
