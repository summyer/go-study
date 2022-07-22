package main_test

import (
	"fmt"
	"sort"
	"testing"
)

/**
map
- 类似其他语言中的哈希表或者字典，以key-value形式存储数据
- key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
- map使用make()创建，支持:=这种简写方式
- make(map[keyType]valueType,cap), cap表示容量，可省略
- 超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素个数
- 键值对不存在时自动添加，使用delete()删除某键值对
- 使用for range对map和slice进行迭代操作
*/

func TestBasic(t *testing.T) {
	//初始化
	var m1 = map[string]string{}
	fmt.Println(m1)
	m1 = make(map[string]string, 10)
	fmt.Println(m1)
	m2 := make(map[int]string, 10)
	fmt.Println(m2)
	m2[1] = "ok"
	fmt.Println(m2[1])
	fmt.Println(m2[2]) //空字符串
	delete(m2, 1)
	fmt.Println("m2[1]", m2[1])

	//key为int，value为map[int]string
	var m map[int]map[int]string = map[int]map[int]string{}
	m = make(map[int]map[int]string)
	m[1] = make(map[int]string) //需要初始化
	m[1][1] = "ok"
	fmt.Println(m[2][1])
	fmt.Println(m)
	//m[2][1] = "ok"
	a, ok := m[2][1] //返回的第二个参数用于判断value是否初始化,嵌套map需要提前初始化
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "ok"
	a, ok = m[2][1]
	fmt.Println(a, ok)
}

func TestFor(t *testing.T) {
	m := map[int]string{1: "2"}
	for k, v := range m {
		fmt.Println(k, v)
		v = "dd"
	}
	fmt.Println(m)
}
func TestFor1(t *testing.T) {
	m := map[int]string{1: "2"}
	for k, v := range m {
		fmt.Println(k, v)
		m[k] = "dd"
	}
	fmt.Println(m)
}
func TestFor2(t *testing.T) {
	m := map[int]string{1: "2"}
	for k := range m {
		fmt.Println(k)
		m[k] = "dd"
	}
	fmt.Println(m)
}
func TestComplex(t *testing.T) {
	//slice+map
	sm := make([]map[int]string, 10)
	for _, v := range sm {
		//v得到的其实是一个拷贝，下面初始化并不成功
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm) //还是都为空
}

func TestComplex2(t *testing.T) {
	//slice+map
	sm := make([]map[int]string, 10)
	for i := range sm {
		//这种方式可以
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm) //还是都为空
}

//对map排序
func TestSort(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for key := range m {
		fmt.Println(key)
		s[i] = key
		i++
	}
	sort.Ints(s) //引用
	for _, v := range s {
		fmt.Println(m[v])
	}
	fmt.Println(s, i)
}

func TestPractice(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	m2 := make(map[string]int, len(m))
	for k, v := range m {
		m2[v] = k
	}
	fmt.Println(m2)
}
