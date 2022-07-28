package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 10)

func f() {
	time.Sleep(2 * time.Second)
	fmt.Println("go")
	c <- 0
}

func main() {
	go f()
	<-c
	print("end")

}
