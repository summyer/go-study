package main_test

import (
	"fmt"
	"testing"
)

/*
常量的定义
- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右边必须是常量或者常量表达式
- 常量表达式中的函数必须是内置函数
初始化规则
- 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式，如果前面一行是并行定义，下一行也需要相同定义个数
- 使用相同的表达式不代表具有相同的值
- iota是常量的计数器，从0开始，组中每定义1个常量自动增1
- 通过初始化规则与iota可以达到枚举的效果
- 每遇到一个const关键字，iota就会重置为0
*/
/*
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。--->基本类型
:=符号会默认左侧声明的是一个变量
*/
func TestConst1(t *testing.T) {
	const (
		MAX_COUNT = 'A' //也可以用MaxCount,但是还是建议大写通过底线连接
	)
	//上面可见性为公开，如果想变成私有，可以换成_MAX_COUNT或者cMAX_COUNT  ,其中c表示const
}

const (
	BYTE = 1 << (10 * (iota + 1))
	KB
	MB
	GB
)

func TestConst2(t *testing.T) {
	fmt.Println(BYTE)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}

/*
运算符
优先级从高到低
^ !   //一元运算符
* / % <<  >> & &^   //二元运算符
+ - | ^
== !=  <  <=  >  >=
<-  //专门用于channel
&&
||
*/
