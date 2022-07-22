package main_test

import (
	"fmt"
	"testing"
)

/**
接口interface
- 接口是一个或多个方法签名的集合
- 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显式声明实现了哪个接口，这被称为StructuralTyping
- 接口只有方法声明，没有实现，没有数据字段
- 接口可以匿名嵌入其他接口，或嵌入到结构中
- 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
- 只有当接口存储的类型和对象都为nil时，接口才等于nil
- 接口调用不会做receiver的自动转换
- 接口同样支持[匿名字段]方法
- 接口也可实现类似OOP中的多态
- 空接口可以作为任何类型数据的容器
*/

type USB interface {
	Name() string
	Connect()
}
type USB2 interface {
	Name() string
	Connecter //嵌入接口
}
type Connecter interface {
	Connect()
}

//不需要显式实现USB接口的方法，只要实现了对应接口的方法就行
type PhoneConnector struct {
	name string
}

func (p PhoneConnector) Name() string {
	return p.name
}

func (p PhoneConnector) Connect() {
	fmt.Println("connect:", p.name)
}

type PhoneConnector2 struct {
	name string
}

func (p *PhoneConnector2) Connect() {
	fmt.Println("connect:", p.name)
}
func Disconnect(usb USB) {
	fmt.Println("disconnect.", usb.Name())
}
func Disconnect2(usb USB2) {
	fmt.Println("disconnect.", usb.Name())
}
func Disconnect3(usb USB) {
	//ok pattern
	//判断usb是否是PhoneConnector
	//类似java里面的instance of
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("disconnect.", pc.name)
		return
	}
	fmt.Println("unknown device.")
}
func TestInterface(t *testing.T) {
	var usb USB
	usb = PhoneConnector{"phoneConnector"}
	usb.Connect()
	fmt.Println(usb.Name())
	u1 := PhoneConnector{"u1"}
	Disconnect(u1)
	Disconnect3(u1)
}

func TestInterface2(t *testing.T) {
	usb := PhoneConnector{"usb"}
	Disconnect2(usb)
}

//go中所有类型都实现了空接口，类似java中的object
type empty interface {
}

//接收参数：空接口
func Disconnect4(usb interface{}) {
	//ok pattern
	//判断usb是否是PhoneConnector
	//类似java里面的instance of
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("disconnect.", pc.name)
		return
	}
	fmt.Println("unknown device.")
}
func TestInterface4(t *testing.T) {
	u1 := PhoneConnector{"u1"}
	Disconnect4(u1)

	Disconnect5(u1)
}

func Disconnect5(usb interface{}) {
	//type switch
	//fmt.Println(usb.(type))  //编译错误，只能用于switch
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("disconnect.", v.name)
	default:
		fmt.Println("unknown device.")
	}
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("connect tv:", tv.name)
}

//类型间转换
func TestInterface5(t *testing.T) {
	pc := PhoneConnector{"phoneConnector"}
	var a Connecter
	a = Connecter(pc)
	a.Connect() //将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
	pc.name = "pc"
	a.Connect() //name并未改变成功

	pc2 := PhoneConnector2{"phoneConnector2"}
	pc2.Connect()
	var a2 Connecter
	a2 = Connecter(&pc2) //拷贝指针（上面一种拷贝的值） // 见errors.go中的 func New(text string) error { return &errorString{text}}
	a2.Connect()
	pc2.name = "pc2"
	a2.Connect()

	//tv := TVConnector{"tv"}
	//var b USB
	//b = USB(tv)   //编译问题，不能这样转换

}
func TestInterfaceNil(t *testing.T) {
	var a interface{}
	fmt.Println(a == nil) //true

	//只有当接口存储的类型和对象都为nil时，接口才等于nil
	var p *int = nil
	a = p
	fmt.Println(a == nil)
}

//接口调用不会做receiver的自动转换
//接口接收的是指针就必须传指针给它，不是指针就不能传指针给它，跟调用结构的方法不同，结构的方法会进行自动的转换
