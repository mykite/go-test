package main

//测试变量的定义go会推断类型
var a = "test"
var b string = "test string"
var c bool

func main() {
	println(a, b, c)
}
