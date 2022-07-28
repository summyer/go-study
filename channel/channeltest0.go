package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("element 1 received from channel ch1: %v\n", elem1)
	elem2 := <-ch1
	fmt.Printf("element 2 received from channel ch1: %v\n", elem2)
	elem3, ret := <-ch1
	if ret == true {
		fmt.Printf("element 3 received from channel ch1: %v\n", elem3)
	}
}
