package main

import "fmt"

/*
类型断言
https://www.cnblogs.com/JunkingBoy/p/15944274.html

类型断言语法格式
value, ok := x.(T)
x表示一个接口的类型，T表示一个具体的类型（也可为接口类型）
语法说明：

断言表达式会返回x的值和一个布尔值，可根据该布尔值判断x是否为T类型：
如果T是具体某个类型，类型断言会检查x的动态类型是否等于具体类型T。如果检查成功，类型断言返回的结果是x的动态值，其类型是T。
如果T是接口类型，类型断言会检查x的动态类型是否满足T。如果检查成功，x的动态值不会被提取，返回值是一个类型为T的接口值。
无论T是什么类型，如果x是nil接口值，类型断言都会失败。
*/

func main() {
	// 声明接口变量
	var x interface{}
	// 赋值
	x = 10
	// 使用接口断言判断(判断接口当中的值和类型满不满足条件)
	value, ok := x.(int) //可以看到x变量是接口变量。x.(int)判断的是接口当中的值的类型是否是int类型
	fmt.Println(value, ",", ok)

	//类型断言配合Switch使用的示例
	// 调用类型判断方法进行判断
	getType(x)
}

/*
定义一个接口类型判断的函数。根据传参判断接口的参数类型
*/
func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("入参类型为int类型!")
	case string:
		fmt.Println("入参的类型为string类型!")
	case float64:
		fmt.Println("入参的类型为float类型!")
	default:
		fmt.Println("未知类型!")
	}
}
