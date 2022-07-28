package main

import "fmt"

func main() {
	a := 1
	for {
		switch a {
		case 1:
			fmt.Println("1")
			break
		case 2:
			fmt.Println("2")
			break
		}
	}
}
