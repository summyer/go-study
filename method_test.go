package main_test

import (
	"fmt"
	"testing"
)

/**
方法method   类似class中的方法
- Go中虽没有class,但依旧有method
- 通过显示说明receiver来实现与某个类型的组合
- 只能为同一个包中的类型定义方法
- Receiver可以是类型的值或者指针
- 不存在方法重载
- 可以使用值或指针来调用方法，编译器会自动完成转换
- 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第1个参数(Method Value vs. Method Expression)
- 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
- 类型别名不会拥有底层类型所附带的方法
- 方法可以调用结构中的非公开字段
*/
type AA struct {
	Name string
}
type BB struct {
	Name string
}

//receiver   (a AA) 值传递
func (a AA) Print() {
	a.Name = "AA-1"
	fmt.Println("A-print")
}

//receiver (a *AA)指针传递
func (a *AA) Print2() {
	a.Name = "AA-2"
	fmt.Println("A-print2")
}

//编译错误，不支持重载
//func (a AA) Print(i int) {
//	fmt.Println("A")
//}

//这种可以，跟不同的结构做的绑定
func (b BB) Print() {
	fmt.Println("A")
}

func TestMethod1(t *testing.T) {
	a := AA{}
	a.Print()
	fmt.Println(a.Name) //修改不成功
	a.Print2()
	fmt.Println(a.Name) //修改成功
	b := BB{}
	b.Print()
}

//类型别名
type TZ int

func (a *TZ) Print() {
	fmt.Println("TZ")
}
func TestMethod2(t *testing.T) {
	var a TZ
	a.Print() //支持，很灵活  TZ和int类型之前也需要强制类型转换，同时不会将int的方法带到TZ中

	//同时不能为int类型绑定方法，因为方法的绑定 只能为同一个包中的类型定义方法

	//(Method Value vs. Method Expression)
	a.Print()       //Method Value
	(*TZ).Print(&a) //Method Expression

	//方法名称冲突解决规则和字段名称冲突一样，根据优先级从最高级找到最低级(注意同名问题)
}

//方法访问权限，方法可以调用结构中的非公开字段
type AAA struct {
	name string
}

func (a *AAA) Print() {
	//方法中都是可以访问私有字段的，因为方法都是要和结构同一包下面才有用。只要这个方法被暴露到其他包中就行
	a.name = "123"
	fmt.Println(a.name)
}

func TestMethod3(t *testing.T) {
	a := AAA{}
	a.Print()
	fmt.Print(a.name) //name小写私有，也是可以访问的，权限是以包(package)为级别的，包内公有，包外私有
}

type Int int

func (a *Int) Increase(num Int) {
	*a += num
}
func (a *Int) Increase1(num int) {
	*a += Int(num) //需要强制转换
}
func TestIncrease(t *testing.T) {
	var a Int
	a = 0
	a.Increase1(100)
	fmt.Println(a)
	a.Increase1(100)
	fmt.Println(a)

}
