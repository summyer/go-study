package main_test

import (
	"fmt"
	"testing"
)

/**
指针  Go虽然保留了指针，但与其他编程语言不同的是，在GO当中不支持指针运算以及“->”运算符，而直接采用“.”选择符来操作指针目标对象的成员
- 操作符“&”取变量地址，使用“*”通过指针间接访问目标对象
- 默认值为nil而为NULL
*/

func TestControl1(t *testing.T) {
	a := 1
	a++ //不能是++a
	var p *int = &a
	fmt.Println(p)  //得到地址
	fmt.Println(*p) //得到p指针指向地址的值1
}

/**
递增递减语句
在Go中，++与--是作为语句而不是作为表达式
*/

/*
if判断语句
条件表达式没有括号
Go 没有三目运算符，所以不支持 ?: 形式的条件判断。
*/
func TestControl2(t *testing.T) {
	a1 := 2
	if a1 > 1 {
	}

	if a := 1; a > 1 { //此时a的作用域就是if这部分块
	}

	a := 10
	if a := 1; a > 0 {
		fmt.Println(a) //用得是内层的a，外部的被覆盖
	}
	fmt.Println(a) //用的就是外层的

	if a, b := 1, 2; a > 0 {
		fmt.Println(a, b)
	}
}

/*
 for循环语句
go只有for一个循环语句关键字，但支持3种形式
*/
func main1() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
	}
}
func main2() {
	a := 1
	for a <= 3 {
		a++
	}

}
func main3() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
	}
}

/*
选择语句
选择语句switch
- 可以使用任何类型或者表达式作为条件语句
- 不需要写break,一旦条件符合自动终止
- 如希望继续执行下一个case，需要使用fallthrough语句
- 支持一个初始表达式（可以是并行方式）,右侧需跟分号
- 左大括号必须和条件语句在同一行
*/

func main4() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}

func main5() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	}
}
func main6() {
	switch a := 1; { //作用域也仅限switch
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")
	}
}

/*
跳转语句
break、continue   配合标签可以跳多层循环  例如 break LABEL      continue LABEL
goto 配合标签调整执行位置
- 三个语法都可以配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- break与continue配合标签可用于多层循环的跳出
- goto是调整执行位置，与其他2个语句配合标签的结果并不相同
*/

func TestControl5(t *testing.T) {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")
}
func TestControl6(t *testing.T) {
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("OK")
}

/*
Go之条件语句
select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
*/
