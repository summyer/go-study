package main

/*
https://www.cnblogs.com/JunkingBoy/p/15944316.html
Go之error接口

Error接口作用
开发中常遇到的问题可分为:
- 异常
- 错误
Go语言的Error接口就是用来返回错误信息的
Go语言中引入error接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含error。error处理过程类似于C语言中的错误码，可逐层返回，直到被处理。

Go语言中Error的源码：
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}

分析：
- 定义了一个签名为Error() string的方法，实现了一个错误类型
- 一般情况下，如果函数需要返回错误，就将 error 作为多个返回值中的最后一个（并非是强制要求）
*/
import (
	"errors"
	"fmt"
	"math"
)

/* 创建error接口 */
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("方法报错!") //使用errors.New来返回一个错误信息
	}
	return math.Sqrt(f), nil
}

/* 方法调用 */
func main() {
	//创建一个简单error
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//自定义错误类型
	result2, err2 := Sqrt1(-13)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(result2)
	}

}

/* 定义错误类型 */
type definedError struct {
	Num     float64
	problem string
}

/* 定义类型实现的接口 */
func (e definedError) Error() string {
	return fmt.Sprintf("错误，原因是： \"%f\"是个自然数", e.Num)
}

/* 创建error接口 */
func Sqrt1(f float64) (float64, error) {
	if f < 0 {
		// 使用自定义错误类型进行返回
		return -1, definedError{Num: f}
	}
	return math.Sqrt(f), nil
}
