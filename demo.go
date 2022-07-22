package main

import "fmt"

var (
	a = 1
	b = 2
	c = 3
)

func t() {
	var (
		a = 1
		b = 2
	)
	fmt.Println(a, b)
}

func main() {
	num := 65
	str := string(num)
	fmt.Printf("%v, %T\n", str, str)
	a++

	type 中文 string
	var a 中文 = "ss"

	t()
	fmt.Println(a)
	//var a1 int = 65
	//b1 := string(a1)
	//fmt.Println(b1)
	switch {
	case true:
		fmt.Println("switch 0")

	}

LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
			//goto LABEL
		}
	}
}
