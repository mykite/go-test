package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

//值传递引用传递，应用传递使用 &xxx
