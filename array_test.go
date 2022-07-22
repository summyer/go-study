package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

/**
数组Array
1.定义数组的格式： var <varName> [n]<type> , n>=0
2.数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
3.注意区分指向数组的指针和指针数组
4.数组在go中为值类型
5.数组之间可以使用==或!=进行比较，但不可以使用<或>
6.可以使用new来创建数组，此方法放回一个指向数组的指针
*/

func TestName(t *testing.T) {
	var aa1 = [2]int{}
	fmt.Println(aa1)
	var a0 = [20]int{19: 1} //将索引19的元素赋值为1
	fmt.Println("设置指定索引的元素", a0)
	var a00 = [...]int{1, 2, 3} //省略数量
	fmt.Println(a00)
	var a000 = [...]int{19: 1}
	fmt.Println(a000)
	//查看变量的类型
	fmt.Println(reflect.TypeOf(a000))
	//数组的指针
	var p *[20]int = &a000
	fmt.Println("以数组的指针取索引元素的值", p[0])
	//指针数组
	x, y := 1, 2
	p1 := [...]*int{&x, &y}
	fmt.Println("数组指针", p1)
	fmt.Println("指针数组", p)
	var a1 [1]int
	a1[0] = 1
	var a2 [1]int
	a2[0] = 2
	var rs bool = a1 == a2
	fmt.Println(rs)

	var a3 = [3]int{1, 2, 3}
	fmt.Println(a3)
	var s1 = [10]string{"i", "love", "you"}
	size := len(s1)
	for i := 0; i < size; i++ {
		fmt.Println(s1[i])
	}
	for i, v := range s1 {
		fmt.Println(i, v)
	}
}

//make   vs   new
//make只用于slice、map以及channel的初始化；而new用于类型的内存分配
//make返回的是引用类型本身；new返回的是指向类型的指针
func TestNewArray(t *testing.T) {
	//返回指向数组的指针
	var a = new([10]int)
	fmt.Println(reflect.TypeOf(a))
	a[0] = 1
	a[1] = 2
	a[3] = 3
	fmt.Println(*a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func TestMakeArray(t *testing.T) {
	//这是一个slice
	var a = make([]int, 2)
	fmt.Println(a)
	var b [2]int
	fmt.Println(b)
	//var b =[2]int{}
	//fmt.Println(a == b)//   invalid operation: a == b (mismatched types []int and [2]int)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func TestMutilArray(t *testing.T) {
	a := [2][3]int{
		{1, 2, 2},
		{2, 3, 3},
	}
	fmt.Println(a)
	a1 := [2][3]int{
		{1: 1},
		{2: 1},
	}
	fmt.Println(a1)

	a2 := [...][3]int{ //非顶级的不能用三个点，如[3]不能换用[...]
		{1: 1},
		{2: 1},
	}
	fmt.Println(a2)
}
