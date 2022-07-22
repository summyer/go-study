package main

//https://www.cnblogs.com/JunkingBoy/p/15239775.html

import (
	"errors"
	"fmt"
)

//自定义一个错误
var errorZero = errors.New("division by zero")

/*
定义一个除法函数，被除数除以除数。如果商为0返回错误
*/
func div(dividend, divisor int) (int, error) {
	//判断除数为0抛出异常
	if divisor == 0 {
		return 0, errorZero
	}
	//正常计算，返回空值
	return dividend / divisor, nil //正常计算返回空的error
	/*
	   进行正常的除法计算，没有发生错误时，错误对象返回 nil
	*/
}

/*
声明一个结构体，包含错误的解析
*/
type ParseError struct {
	Filename string //文件名
	Line     int    //行号
	/*
	   声明了一个解析错误的结构体，解析错误包含有 2 个成员，分别是文件名和行号
	*/
}

/*
声明一个error接口，指向结构体，返回错误的描述
*/
func (e *ParseError) Error() string { //实现了错误接口，将成员的文件名和行号格式化为字符串返回
	return fmt.Sprintf("%s,%d", e.Filename, e.Line)
}

//创建一些解析错误??
/*
根据给定的文件名和行号创建一个错误实例--->这是一个错误实例
*/
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
func main() {
	var errorZero = errors.New("division by zero")
	fmt.Println(errorZero)
	fmt.Println(div(1, 0))

	//自定义error
	var e error //声明一个错误接口类型
	//创建一个错误实例
	e = newParseError("main.go", 1)
	/*
	   创建一个实例，这个错误接口内部是 *ParserError 类型，携带有文件名 main.go 和行号 1
	*/

	//通过error接口查看错误描述
	fmt.Println(e.Error())

	//通过断言获取详细的错误信息
	switch detail := e.(type) {
	case *ParseError: //解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default:
		//其他类型的错误
		fmt.Println("other error")
	}
}
