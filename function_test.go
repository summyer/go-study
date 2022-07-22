package main_test

import (
	"fmt"
	"testing"
)

/**
函数function
- go函数不支持嵌套、重载和默认参数
- 但支持以下特性：
	无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
- 定义函数使用关键字func,且左大括号不能另起一行
- 函数也可以作为一种类型使用
*/
func A(a int, b int) {
}
func A1(a int, b int) int {
	return 1
}
func A2(a int, b int) (int, int) {
	return 1, 1
}

//a,b,c都是int时简写
func A3(a, b, c int) (int, int) {
	return 1, 1
}
func A4(a, b int) (a1 int, b2 int, c2 int) {
	return 1, 1, 1
}

//命名返回值不要return后面带参数,不过为了可读性还是像上面那要在return中返回较好
func A41(a, b int) (a1 int, b2 int, c2 int) {
	a1, b2, c2 = 1, 2, 3
	return
}
func A42(a, b int) (a1 int, b2 int, c2 int) {
	a1, b2, c2 = 1, 2, 3
	return a1, b2, c2
}
func A43(a, b int) (a1 int, b2 int, c2 int) {
	a1, b2, c2 = 1, 2, 3
	return a1, b2, b
}

//返回值简写需要命名返回值
func A5(a, b int) (a1, b2, c2 int) {
	return 1, 1, 1
}

//不定长变参
func A6(a ...int) {
	fmt.Println(a) //接收的是slice
}

//不定长变成后面不能有其他变量，例如(a ...int,b string)这种就不行
func A7(b string, a ...int) {
	fmt.Println(a) //接收的是slice
}

func B(s ...int) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}
func B1(s []int) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}
func B2(a *int) {
	*a = 2
	fmt.Println(*a)
}

func TestA(t *testing.T) {
	A6(1, 2, 3)
}
func TestB(t *testing.T) {
	a, b := 1, 2
	B(a, b) // 传值
	fmt.Println(a, b)
	s1 := []int{3, 4}
	B1(s1) // 传引用
	fmt.Println(s1)
	b1 := 3
	B2(&b1) //传地址
	fmt.Println(b1)

	var f = B
	f(1, 2)
	//匿名函数
	var f1 = func() {
		fmt.Println("匿名函数")
	}
	f1()
	f2 := closure(1)
	fmt.Println(f2(2))
	fmt.Println(f2(3))
}

//闭包  返回值为函数类型
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
