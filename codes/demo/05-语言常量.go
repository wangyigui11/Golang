// 常量是一个简单值的标识符，在程序运行时，不会被修改的量
// 常量的数据类型只可以使布尔型、数字型(整数型、浮点型、复数型)和字符串型

// 常量还可以用作枚举
// const (
// 	UnKnown = 0
// 	Female = 1
// 	Male    = 2
// )

package main

import (
	"fmt"
	"unsafe"
)

const (
	abc  = "abc"
	bLen = len(abc)
	ccc  = unsafe.Sizeof(abc)
)

func main() {
	const LENGTH int = 10

	const WIDTH int = 5

	var area int

	const a, b, c = 1, false, "" // 多量赋值

	area = LENGTH * WIDTH

	fmt.Println("面积为 : ", area) //  50

	println()

	println(a, b, c) // 1 false ""

	println(abc, bLen, ccc) // abc 3 16
}
