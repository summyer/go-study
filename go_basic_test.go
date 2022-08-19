package main_test

import (
	"testing"
)

/*
内置关键字(25个均为小写)
break default func interface select
case defer go map struct
chan else goto package  switch
const fallthrough if range type
continue for import return var
*/
//go 注释方法
//    单行注释
/*多行注释*/
/*
Go程序的一般结构

* Go程序是通过package来组织的（与python类似）
* 一个目录中只能有一个package,  但是可以增加一个package [当前packageName]_test
* package必须放着非注释代码第一行
* 只有package名称为main的包可以包含main函数
* 一个可执行程序有且仅有一个main包，这个main包有且仅有一个main函数
* 通过import关键字来导入其他非main包
* 通过const关键字来进行常量的定义
* 通过在函数体外部使用var关键字来进行全局变量的声明和赋值
* 通过type关键字来进行结构(struct)或接口(interface)的声明
* 通过func关键字来进行函数的声明
*/

//行分隔符
//在 Go 程序中，一行代表一个语句结束--->如果打算将多个语句写在同一行，它们则必须使用 ; 人为区分

/*
import packageName
	//package别名与省略调用
	import(
		"fmt"
	"io"
	"os"
	"strings"
	"time"
	)
	//导入包之后，就可以使用<PackageName>.<FuncName>来对包中的函数进行调用
	//如果导入包之后未调用其中的函数或者类型将会报编译错误
	package 别名\省略
	import  a  fmt
	import .   fmt     //不建议使用，易混淆     ；  不可以和别名同时使用（使用的别名就不能通过省略调用）
    //使用【import _ 包路径】只是引用该包，仅仅是为了调用 init() 函数，所以无法通过包名来调用包中的其他函数。
*/

/*
Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或者函数是否可以被外部包所调用：
根据约定，函数名首字母小写即为private, 函数名首字母大写即为public
*/
func TestBasic1(t *testing.T) {

}

/*
各种组
*/
//下面称为组
//常量的定义
const (
	PI     = 3.14
	const1 = 1
)

//全局变量的声明和赋值， ~~在函数体内不能这样使用~~
var (
	name  = "1"
	name1 = "2"
)

//一般类型声明
type (
	newType int
	type1   float32
	type2   string
	//type3 [type]
)
