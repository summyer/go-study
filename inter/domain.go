package main

import "fmt"

type people interface {
	GetName() string
	GetAge() int
}
type student struct {
	name string
	age  int
}

func (receiver *student) GetName() string {
	return ""
}
func (receiver *student) GetAge() int {
	return 0
}
func main() {
	stu := student{}
	ob, ok := interface{}(&stu).(people)
	fmt.Println(ob, ok)
	ob1, ok1 := interface{}(stu).(people)
	fmt.Println(ob1, ok1)
}
