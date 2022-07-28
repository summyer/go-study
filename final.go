package main

import (
	"fmt"
	"time"
)

func App(s []int) {
	fmt.Println("方法中", s[0])
	fmt.Printf("方法中1 %p\n", &s)
	s = append(s, 3)
	fmt.Printf("方法中2 %p\n", &s)
}
func App2(s []int) []int {
	fmt.Println("方法中", s[0])
	fmt.Printf("方法中1 %p\n", &s)
	s = append(s, 3)
	fmt.Printf("方法中2 %p\n", &s)
	return s
}
func main() {
	// slice的坑
	s := make([]int, 1)
	s[0] = 1
	fmt.Println(s)
	fmt.Printf("%p\n", &s)
	App(s)
	fmt.Println(s)
	fmt.Printf("%p\n", &s)

	s2 := make([]int, 1)
	s2[0] = 1
	fmt.Println(s2)
	fmt.Printf("%p\n", &s2)
	s2 = App2(s2)
	fmt.Println(s2)
	fmt.Printf("%p\n", &s2)

	TimeTest()
	ForRangeTest()
}
func TimeTest() {
	//时间格式化 的坑
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC)) //layout用常量，不要自己改，时间会乱
}
func ForRangeTest() {
	//for range 闭包的坑
	s := []string{"a", "b", "c"}
	/*for _, v := range s {
		go func() {
			fmt.Println(v)  //有误
		}()
	}*/
	for _, v := range s {
		go func(v string) {
			fmt.Println(v) //对的 // ,因为用了select{}会提示，fatal error: all goroutines are asleep - deadlock!
			//下面这种就不会提示：fatal error: all goroutines are asleep - deadlock!
			//for {
			//	fmt.Println(v) //对的
			//}
		}(v)
	}
	select {}
}
