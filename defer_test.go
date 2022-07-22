package main_test

import (
	"fmt"
	"testing"
)

/*
defer panic recover
- 执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
- 即使函数发生严重错误也会执行
- 支持匿名函数的调用
- 常用于资源清理、文件关闭、解锁以及记录时间等操作
- 通过与匿名函数配合可在return之后修改函数计算结果
- 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址
- go没有异常机制，但有panic/recover模式来处理错误
- panic可以在任何地方引发，但recover只有defer调用的函数中有效
*/
func Test1(t *testing.T) {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
func Test2(t *testing.T) {
	for i := 0; i < 3; i++ {
		//值
		defer fmt.Println(i)
	}
}
func Test3(t *testing.T) {
	for i := 0; i < 3; i++ {
		//闭包
		defer func() {
			//引用
			fmt.Println(i) //都是3
		}()
	}
}

func TestPanic(t *testing.T) {
	PA()
	PB()
	defer func() {
		fmt.Println("中间defer")
	}()
	PC()
}
func PA() {
	fmt.Println("A")
}
func PB() {
	//defer需要放在panic之前
	defer func() {
		fmt.Println("第一个defer")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B:", err)
		}
	}()
	panic("panic in b")
}
func PC() {
	fmt.Println("C")
}

func TestPractice2(t *testing.T) {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() {
			fmt.Println("defer_closure =", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i=", i)
		}
	}
	for _, f := range fs {
		f()
	}
}
