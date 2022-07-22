package main

import (
	util2 "demo/util" //导入时是相对go文件所在目录的路径，这也是为什么一个目录下只能有一个package
)

func main() {
	a := util2.AAA{} //调用时是package的名称，（package名不一定和go文件名一样）
	a.Print()
}
