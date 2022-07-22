package util2

import "fmt"

type AAA struct {
	name string
}

func (a *AAA) Print() {
	a.name = "123"
	fmt.Println(a.name)
}
