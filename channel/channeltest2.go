package main

import (
	"fmt"
)

var a2 string
var c2 = make(chan int)

func f2() {
	fmt.Println("aa")
	c2 <- 1
}

func main() {
	go f2()

	<-c2
	fmt.Println(a2)

}
