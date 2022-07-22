package main

/*
go get：获取远程包
go run:直接运行程序
go build:测试编译，检查是否有编译错误
go fmt:格式化源码
go install:编译包文件并编译整个程序
go test:运行测试文件  以_test.go结尾的就是测试文件
go doc:查看文档    开启本地文档服务：godoc -http=:8080    //godoc需要先安装

执行go install时有时需要将bin中的可执行程序拷贝到src中，因为程序中使用的相对路径
执行go build a.go时想产生可执行程序， a.go中的package名称需要是main
*/
