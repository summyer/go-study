package main

import (
	"fmt"
	"reflect"
)

/*
反射 reflection
- 反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
- 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
- 反射会将匿名字段作为独立字段（匿名字段本质）
- 想要利用反射修改对象状态，前提是 interface.data是settable,即pointer-interface
- 通过反射可以"动态"调用方法
*/
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}
func (u User) Hello2(name string) {
	fmt.Println("hello ", name, ",my name is", u.Name)
}

//获取对象字段信息，类型信息,方法信息
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type:", t.Name())
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("非struct结构不能反射取对象信息")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v=%v\n", field.Name, field.Type, val)
	}
	fmt.Println("Methods:")
	for i := 0; i < t.NumMethod(); i++ { //只能列出public方法
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type) //Hello:func(main.User)
	}
}

func main() {
	u := User{1, "nik", 20}
	Info(u)
	Info(&u) //这种指针，会报错

	/*反射匿名或嵌入字段*/
	m := Manager{User: User{1, "ok", 20}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0)) //reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10d0440), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	fmt.Printf("%#v\n", t.Field(1)) //reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x10c3e20), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}
	//取User中的id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) //reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x10c3700), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	//取User中的name
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1})) //reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x10c3ee0), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}

	/*反射修改内容*/
	//基本类型
	x := 123
	v := reflect.ValueOf(&x) //修改时需要传指针  pointer-interface
	v.Elem().SetInt(999)
	fmt.Println(x)
	//结构
	u1 := User{1, "ok", 12}
	Set(&u1)
	fmt.Println(u1)

	/*反射调用方法*/
	u2 := User{1, "ok", 12}
	u2.Hello2("ni")
	v2 := reflect.ValueOf(u2) //传&u2也行
	mv := v2.MethodByName("Hello2")
	if !mv.IsValid() {
		fmt.Println("BAD methodName")
		return
	}
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args) //如果有返回值，可以直接接收call返回值 rs:=mv.Call(args)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	//if f := v.FieldByName("Name"); f.Kind() == reflect.String {
	//	f.SetString("BYEBYE")
	//}
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}

type Manager struct {
	User
	title string
}
