package main_test

import (
	"fmt"
	"sort"
	"testing"
)

/**
切片Slice
- 其本身并不是数组，它指向底层的数组
-作为变成数组的替代方案，可以关联底层数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组获取生成
- 使用len()获取元素个数，cap()获取容量
- 一般使用make()创建
- 如果多个slice指向相同底层数组，其中一个的值改变会影响全部
- make([]T,len,cap)
- 其中cap可以省略，则和len的值相同
- len表示存储的元素个数，cap表示容量
*/
func Test(t *testing.T) {
	var aa []int
	fmt.Println(aa)
	var array = [10]int{1, 3, 2}
	var s0 = array[1:10] //基于数组创建切片，包含起始索引，不包含终止索引
	s0 = array[1:]       //从索引1开始，到最后
	s0 = array[:5]       //从0到索引5
	fmt.Println(s0)
	var s1 = []int{1, 2}
	fmt.Println(s1)
	var a = make([]int, 3, 10)
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println(a)

	fmt.Println(len(a), cap(a))
	//只能给slice排序，不能是array
	sort.Ints(s0)
	//遍历切片
	for i := 0; i < len(s0); i++ {
		fmt.Println(s0[i])
	}
	for i, v := range array {
		fmt.Println(i, v)
	}
}

//Reslice:基于一个slice通过[x:y]这种方式产生新的slice
/**
  - Reslice时索引以被slice的切片为准
  - 索引不可以超越被slice的切片的容量的cap()值
  - 索引越界不会导致底层数组的重新分配而是引发错误
*/
func TestReslice(t *testing.T) {
	var s4 = []int{1, 2, 3, 4, 5, 6, 7}
	var s5 = s4[2:]
	fmt.Println(s5)
}

/**
Append
- 可以在slice尾部追加元素
- 可以将一个slice追加在另一个slice尾部
- 如果最终长度未超过追加到slice的容量则返回原始slice
- 如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
*/
func TestAppend(t *testing.T) {
	s1 := make([]int, 3, 6)
	fmt.Printf("原地址%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p \n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p \n", s1, s1)
}

func TestChange(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	//append个数超过slice容量的话会创建一个新的数组
	s2 = append(s2, 1, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	s1[0] = 9
	fmt.Println(s1, s2)
}

/**
Copy
*/
func TestCopy(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	//将s2拷贝到s1
	//copy(s1, s2)
	//fmt.Println(s1)

	//copy(s2, s1)
	//fmt.Println(s2)

	//copy(s2, s1[1:3])
	//fmt.Println(s2)
	copy(s2[1:3], s1[1:3])
	fmt.Println(s2)

	s3 := s1[:] //拷贝s1全部
	fmt.Println(s3)
}

func TestForeach(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println(i, v)
		slice[i] = 0
	}
	fmt.Println(slice)
}
