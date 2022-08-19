# 静态资源  https://jishuin.proginn.com/p/763bfbd61c89
静态文件，也有人叫资产或资源，是一些被程序使用、没有代码的文件。在 Go 中，这类文件就是非 .go 的文件。它们大部分被用在 Web 内容，像 HTML 、javascript 还有网络服务器处理的图片，然而它们也可以以模板、配置文件、图片等等形式被用在任何程序中。主要问题是这些文件不会随代码一起被编译。

# go语言规范文档
https://golang.google.cn/ref/spec

# 
int(a)
usb.()
reflect.Typeof()
v := reflect.ValueOf(o)
if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {

}

str := `hello
world
v2.0`
fmt.Println(str)


string到int
int,err:=strconv.Atoi(string)

string到int64
int64, err := strconv.ParseInt(string, 10, 64)

int到string
string:=strconv.Itoa(int)

int64到string
string:=strconv.FormatInt(int64,10)


# 交叉编译
## 在Mac上编译可运行在Linux, Windows上的GO程序
编译运行在 amd64位 linux系统
>CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main_run.go

linux上执行：
> chmod 777 main_run
> ./main_run
## 编译运行在 amd64位 windows系统
>CGO_ENABLED=0 GOOS=windows  go build main_run.go

window上执行：
>main_run.exe

## 此程序在linux上执行报错:找不到config/global_config.yml
解决:在可执行程序目录下新建config目录，放入global_config.yml