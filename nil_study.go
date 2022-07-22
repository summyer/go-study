package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
Go接口的nil判断
- nil在Go语言中只能被赋值给指针和接口
- 显示地将nil赋值给接口时，接口的type和data都将为nil。此时，接口与nil值判断是相等的。
- 将一个带有类型的nil赋值给接口时，只有data为nil，而type不为nil，此时，接口与nil判断将不相等。
*/
type Student interface {
	ShowInfo() string
}

type stud struct {
}

func (s *stud) ShowInfo() string {
	return "HI"
}
func getStudent() Student {
	var s *stud = nil //nil只能赋值给接口或者指针，不支持 var s stud = nil
	fmt.Println("getStudent:", s, reflect.TypeOf(s), nil == s)
	return s
}

/*
定义一个函数，返回该接口类型
*/
func getStudent2() Student {
	// 声明stu类型的指针变量--->带有类型的nil（类型为：*stu）
	var s *stud = nil
	fmt.Println("getStudent2:", s, reflect.TypeOf(s), nil == s)
	if nil == s {
		return nil
	}
	return s
}

func main() {
	var a interface{}
	fmt.Println(a == nil)
	a = nil
	fmt.Println(a == nil)
	//
	var ss Student
	ss = nil
	fmt.Println(ss, reflect.TypeOf(ss))
	if nil == ss {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	fmt.Println("==========带有类型的nil赋值给接口==========")
	ss = getStudent() //虽然执行：赋值nil操作，但是代码执行结果为："收到的信息是：not nil"
	fmt.Println(ss, reflect.TypeOf(ss))
	if nil == ss {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	//修改后
	fmt.Println("==========（修改后）带有类型的nil赋值给接口==========")
	ss = getStudent2()
	if nil == ss {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	testNil2()
}

/**
Go之容器之nil
nil的特点
- 是Go语言中一个预定义好的标识符
- 是"指针、切片、映射、通道、函数、接口的零值"
nil与null不同点
nil标识符是不能比较的
==对于 nil 来说是一种未定义的操作

*/
func testNil2() {
	//==对于 nil 来说是一种未定义的操作
	//fmt.Println(nil == nil) //invalid operation: nil == nil (operator == not defined on nil)

	//nil不是关键字或保留字,可以定义一个名为nil的变量
	//nil := "Hello"
	//fmt.Println(nil)
	//nil没有默认类型fmt.Printf("%T", nil)
	fmt.Println(nil)

	//不同类型nil的指针是一样的
	/*声明两个变量*/
	var arr []int
	var num *int
	/*分别利用格式化输出他们的nil值*/ //arr 和 num 的指针都是 0x0
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p", num)

	//指针指向的内存地址不一样
	///*分别利用格式化输出他们的nil值*/ //这两个的结果是不一样的
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p", &num)

	//不同类型的nil值不能比较
	/*声明两个变量*/
	var mapValue map[int]string
	var ptr *int
	fmt.Println(mapValue, ptr)
	//fmt.Printf(mapValue == ptr)//这会导致编译无法通过

	//两个相同类型的nil值也可能无法比较
	/*声明两个切片*/
	var num1 []int
	var num2 []int
	fmt.Println(num1, num2)
	//fmt.Printf(num1 == num2)
	//fmt.Printf(num1 == nil) //这也会导致编译无法通过--->不能将空值直接与nil标识符进行比较

	//nil 是 map、slice、func、pointer、channel、interface 的零值
	var m map[int]string
	var ptr2 *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr2)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)

	//不同类型的nil值占用的内存大小可能是不一样的
	/*
	* 一个类型的所有的值的内存布局都是一样的，nil 的大小与同类型中的非 nil 类型的大小是一样的。
	* 不同类型的 nil 值的大小可能不同。
	 */
	var p2 *struct{}
	fmt.Println(unsafe.Sizeof(p2)) // 8
	var s2 []int
	fmt.Println(unsafe.Sizeof(s2)) // 24
	var m2 map[int]bool
	fmt.Println(unsafe.Sizeof(m2)) // 8
	var c2 chan string
	fmt.Println(unsafe.Sizeof(c2)) // 8
	var f2 func()
	fmt.Println(unsafe.Sizeof(f2)) // 8
	var i2 interface{}
	fmt.Println(unsafe.Sizeof(i2)) // 16
}

/**
https://www.cnblogs.com/JunkingBoy/p/15196621.html
总结：
- nil不是Go语言当中预先定义好的标识符或保留字，可以作为变量名
- 不同类型的nil无法比较，相同类型的nil有可能也无法比较
- nil是引用类型的零值，map、slice、function的nil值无法比较--->其他的引用类型是可以的
*/
