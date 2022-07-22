package main_test

import (
	"fmt"
	"testing"
)

/*
结构struct   类似class
- go中的struct与C中的struct非常相似，并且GO没有class
- 使用type <Name> struct{}定义结构，名称遵循可见性规则
- 支持指向自身的指针类型成员
- 支持匿名结构，可用作成员或定义成员变量
- 匿名结构也可以用于map的值
- 可以使用字面值对结构进行初始化
- 允许直接通过指针来读写结构成员
- 相同类型的成员可进行直接拷贝赋值
- 支持==与!=比较运算符，但不支持>或<
- 支持匿名字段，本质上是定义了以某个类型为名称的字段
- 嵌入结构作为匿名字段看起来像继承，但不是继承
- 可以使用匿名字段指针
*/
type test struct {
}
type person0 struct {
	Name string
	Age  int
}
type person struct {
	Name string
	Age  int
}

func ChangePerson(per person) {
	per.Age = 20
	per.Name = "lisa"
	fmt.Println(per)
}
func ChangePerson2(per *person) {
	per.Age = 20
	per.Name = "lisa"
	fmt.Println(per)
}
func ChangePerson3(per *person) {
	per.Age = 20
	per.Name = "lisa"
	fmt.Println(per)
}

func TestStructTest(t *testing.T) {
	//值
	a := test{}
	fmt.Println(a)
	b := person{}
	fmt.Println(b)
	b.Name = "joe"
	b.Age = 19
	fmt.Println(b)

	//字面值初始化
	c := person{
		Name: "joe",
		Age:  19, //注意此处要加,
	}
	fmt.Println(c)
	ChangePerson(c) //传值，这里并不能改变c中的属性
	fmt.Println(c)
	ChangePerson2(&c) //传引用，会改变里面的属性
	fmt.Println(c)

	//开发中也推荐这种形式
	//初始化时就取指针
	d := &person{
		Name: "joe",
		Age:  19, //注意此处要加,
	}
	//不用这样操作指针属性 *d.Name=""，直接d.Name="ok"就行
	ChangePerson2(d)
	fmt.Println(d)
	ChangePerson3(d)
	fmt.Println(d)
}

func TestStruct2(t *testing.T) {
	a := struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
	b := &struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(b)
}

type person2 struct {
	Name    string
	Age     int
	Contact struct {
		Phone string
		City  string
	}
}

type person3 struct {
	Name    string
	Age     int
	Contact struct { //匿名结构
		Phone, City string //这样很方便，也是为什么go将类型放着后面
	}
}

func TestStruct3(t *testing.T) {
	a := person3{}
	fmt.Println(a)
	b := person3{
		Name: "",
		Age:  0,
	}
	//匿名结构初始化
	b.Contact.City = "hz"
	b.Contact.Phone = "182"
	fmt.Println(b)
	c := person3{
		Name: "",
		Age:  0,
		Contact: struct { //匿名结构初始化
			Phone, City string
		}{},
	}
	fmt.Println(c)
}

type person4 struct {
	string //匿名字段
	int    //匿名字段
}

func TestStruct4(t *testing.T) {
	a := person4{"joe", 19} //保持字段顺序
	//b := person4{19,"joe"}  //这种不行
	fmt.Println(a)
}
func TestStruct5(t *testing.T) {
	a := person{"joe", 19}
	//相同结构直接赋值
	b := a
	fmt.Println(b)
}
func TestStruct6(t *testing.T) {
	//a := person0{"joe", 19}
	//b := person{"joe", 19}
	// fmt.Println(a == b)   这种情况a和b不同类型，不能比较
	a := person{"joe", 19}
	b := person{"joe", 19}
	fmt.Println(a == b) //结果为true

	c := struct {
		Name string
		Age  int
	}{"joe", 19}
	d := struct {
		Name string
		Age  int
	}{"joe", 19}
	fmt.Println(c == d) //这两个匿名结构一样的可以比较
}

type person5 struct {
	string //匿名字段
	int    //匿名字段
	Name   string
}

func TestStructP5(t *testing.T) {
	a := person5{"1", 19, "1"}
	//a := person5{"1", 19, Name:"1"} //不支持这种，编译错误
	b := person5{string: "1", int: 19, Name: "1"}
	fmt.Println(a, b)
}

//go中的"继承"
//go中不存在继承，以组合(嵌入结构)实现"继承"
type human struct {
	Sex int
}
type teacher struct {
	human //匿名字段
	Name  string
	Age   int
}
type student struct {
	human //嵌入字段作为匿名字段，本质上将结构的名称作为字段的名称，即此处名称为human
	Name  string
	Age   int
}

type student2 struct {
	p    human
	Name string
	Age  int
}

func TestStruct7(t *testing.T) {
	a := teacher{Name: "joe", Age: 19}
	b := student{Name: "joe", Age: 19}
	fmt.Println(a, b)

	//a := teacher{Name: "joe", Age: 19, Sex: 0} //不支持这样，编译就出错
	//b := student{Name: "joe", Age: 19, Sex: 1}

	//a := teacher{Name: "joe", Age: 19,human{Sex:0}} //也不支持这样，编译就出错
	//b := student{Name: "joe", Age: 19,human{Sex:1}}

	//a := teacher{human{Sex: 0}, Name: "joe", Age: 19} //也不支持这样，编译就出错
	//b := student{human{Sex: 1}, Name: "joe", Age: 19}
	//初始化human方法一
	a1 := teacher{human: human{Sex: 0}, Name: "joe", Age: 19}
	b1 := student{human: human{Sex: 1}, Name: "joe", Age: 19}
	fmt.Println(a1, b1)
	a1.human.Sex = 100
	fmt.Println(a1) //支持
	a1.Sex = 1000   //嵌入结构的字段默认给了子类
	fmt.Println(a1) //支持

	b2 := student2{p: human{Sex: 1}, Name: "joe", Age: 20}
	fmt.Println(b2)
	b2.p.Sex = 100
	fmt.Println(b2)
	//b2.Sex=1000 //不支持，需要匿名字段才行
}

//嵌入结构和被嵌入结构存在同名字段
type Astruct struct {
	Bstruct
	Name string
}

type A2struct struct {
	Bstruct
	Cstruct
}
type Bstruct struct {
	Name string
}
type Cstruct struct {
	Name string
}

func TestStruct11(t *testing.T) {
	a := Astruct{Bstruct: Bstruct{"in"}, Name: "out"}
	fmt.Println(a) //{{in} out}
	//a.Name优先取当前结构中的字段 ，没有就找嵌入结构中的字段
	fmt.Println(a.Name, a.Bstruct.Name) //out in

	a2 := A2struct{Bstruct: Bstruct{"B"}, Cstruct: Cstruct{"C"}}
	//fmt.Println(a2.Name) //编译失败，存在冲突重名字段
	fmt.Println(a2.Bstruct.Name)
}
