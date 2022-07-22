package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
Go基本类型
- 布尔型  bool
    - 长度1字节
    - 取值范围 true,false
    - 注意事项 不可以用数字代表true或false
- 整型  int/uint
    - 根据运行平台可能为32或64
- 8位整型 int8/uint8
    - 长度为1字节
    - 取值范围 -128~127/0~255
- 字节 byte    //uint8的别名
- 16位整型  int16/uint16
    - 长度2字节
- 32位整型 int32( rune )/uint32
    - 长度4字节
- 64位整型 int64/uint64
    - 长度8字节
- 浮点型 float32/float64
    - 长度4/8字节
    - 小数位 精确到7/15小数位
- 复数 complex64/complex128
    - 长度 8/16字节
- 足够保持指针的32位或64位整数型 uintptr
- 其他值  array、struct、string
- 引用类型 slice 、 map、 chan
- 接口类型 interface
- 函数类型  func
*/

/*
Go之字符类型(byte和rune)
- uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符(这是一种无符号类型)。
- 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型(这是一种有符号类型)。
*/

/*
类型零值
值类型都是0，bool为false， string为空字符串，引用类型为nil
*/

/*
类型别名
*/
func TestType1(t *testing.T) {
	//类型别名
	type (
		byte int8
		文本   string
	)
}

/*
变量的声明与赋值
单个变量的声明于赋值
- 变量的声明格式：var <变量名称>  <变量类型>
- 变量的赋值格式：<变量名称>=<表达式>
- 声明的同时赋值：var <变量名称> [变量类型] =<表达式>
- 声明的同时赋值最简写法 d:=456   //不能用于全局变量
多个变量的声明与赋值
- 全局变量的声明可以使用var()的方式进行简写
- 全局变量的声明不可以省略var,但可使用并行方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用var()方式简写，只能使用并行方式
*/
func TestVar2(t *testing.T) {
	//并行、类型推导
	var (
		aaa      = "hello"
		sss, bbb = 1, 2
	)
	a, b, c := 1, 2, 3
	fmt.Println(aaa, sss, bbb, a, b, c)
}

/*
类型转换
- Go中不存在隐私转换，所有类型转换必须显示声明
- 转换只能发生在两种相互兼容的类型之间
- 类型转换的格式：
  <ValueA>[:]=<TypeOfValueA>(<ValueB>)
*/

func TestConvert(t *testing.T) {
	//int转字符串
	//var a = string(65) //->实际是A
	//fmt.Println(reflect.TypeOf(a))

	fmt.Println(strconv.Itoa(65)) // ->这个是字符串65
	fmt.Println(strconv.Atoi("65"))
}

func TestConvert2(t *testing.T) {
	num := 65
	str := string(num) //单元测试会提示conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)。
	//但是在main方法是可以执行成功的
	fmt.Printf("%v, %T\n", str, str)
}

//A：65, string；B：A, string；C：65, int；D：报错
//正确答案：B。
//不过，如果你使用 go vet 检查，会提示：conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)。
//也就是说，如果要将数字 65 转为字符串，不能使用 string(num)，如果使用这种方式转，得到的是一个 rune 的字符串表示，因为字面 A 的 ASCII 码是 65，因此这里输出结果是 A,string。
//如果要禁用go vet的运行，请使用-vet=off标志。
