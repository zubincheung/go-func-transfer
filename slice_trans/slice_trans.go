package main

import "fmt"

func change(x []int) {
	fmt.Printf("\nx: %#v x内存地址:%p \n", x, &x)
	for i, _ := range x {
		fmt.Printf("x[%d]内存地址：%p\n", i, &x[i])
	}
	x[1] = 8
	fmt.Printf("\n执行后 x:%#v \n", x)
}

func main() {
	a := []int{1, 2, 3}

	fmt.Printf("a: %#v a内存地址:%p \n", a, &a)
	for i, _ := range a {
		fmt.Printf("a[%d]内存地址：%p\n", i, &a[i])
	}

	change(a)

	fmt.Printf("\n执行后 a:%#v \n", a)
}
