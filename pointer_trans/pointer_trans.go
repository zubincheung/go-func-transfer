package main

import (
	"fmt"
)

func inc(x *int) {
	fmt.Printf("x 内存地址:%p 值:%p \n", &x, x)
	*x++
	fmt.Printf("x:%d \n", *x)
}

func main() {
	a := 1

	fmt.Printf("a 内存地址:%p 值:%d \n", &a, a)
	inc(&a)
	fmt.Printf("执行后 a:%d \n", a)
}
