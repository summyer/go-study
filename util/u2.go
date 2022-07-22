package util2

import "fmt"

type AAA2 struct {
	name string
}

func (a *AAA2) Print() {
	a.name = "123"
	fmt.Println(a.name)
}
