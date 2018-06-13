# golang 函数参数传递详解

参数传递是指在程序的传递过程中，实际参数就会将参数值传递给相应的形式参数，然后在函数中实现对数据处理和返回的过程。比较常见的参数传递有：值传递、指针传递、引用传递。

一直以为 Go 里面函数传参有值传递和引用传递两种方式，对指针传递和引用传递的区别也不是很清楚，查看了下官方文档才发现并不是这么回事。

> In a function call, the function value and arguments are evaluated in the usual order. After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution.

文档地址：https://golang.org/ref/spec#Calls
或者：http://docs.studygolang.com/ref/spec#Calls

官方文档已经明确说明：**Go 里边函数传参只有值传递一种方式**，为了加强自己的理解，再把每种传参方式进行一次梳理。

## 值传递

> 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

运行以下代码：

```go
func inc(x int) {
	fmt.Printf("x 内存地址:%p 值:%d \n", &x, x)
	x++
	fmt.Printf("x:%d \n", x)
}

func main() {
	a := 1

	fmt.Printf("a 内存地址:%p 值:%d \n", &a, a)
	inc(a)
	fmt.Printf("执行后 a:%d \n", a)
}
```

执行结果为：

```
a 内存地址:0xc42008c008 值:1
x 内存地址:0xc42008c018 值:1
x:2
执行后 a:1
```

可以看出程序中使用的是值传递，形参 x 是实参 a 在栈上的一份拷贝， 和实参拥有两个完全不同的地址，在函数内部改变了 x 的值，a 并没有改变

## 指针传递

> 形参为指向实参地址的指针，当对形参的指向操作时，就相当于对实参本身进行的操作。

修改上面代码

```go
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
```

执行结果为：

```
a 内存地址:0xc42008c008 值:1
x 内存地址:0xc4200a0020 值:0xc42008c008
x:2
执行后 a:2
```

可以看到指针&a 传给函数的形参 x 后，形参将会是它在栈上的一份拷贝，他们本身将各自拥有不同的地址，但是二者的值是一样的（都是变量 a 的地址），因此可以通过指针相关的操作来改变 a 的值。

## 引用传递

> 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数

由于 Go 中并不存在引用传递，我们看下一段 c++代码

```c++
void inc( int& x){
    printf("x 内存地址:%p 值:%d \n", &x, x);
    x++;
    printf("x:%d \n", x);
}

int main(int argc, const char * argv[]) {
    // insert code here...
    int a=1;
    printf("a 内存地址:%p 值:%d \n", &a, a);
    inc(a);
    printf("执行后 a:%d \n", a);
    return 0;
}
```

执行结果：

```
a 内存地址:0x7ffeefbff57c 值:1
x 内存地址:0x7ffeefbff57c 值:1
x:2
执行后 a:2
```

可以看到引用传递，操作地址就是实参地址 ，只是相当于实参的一个别名，对它的操作就是对实参的操作

那么传 slice 是不是传引用呢？

```go
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
```

执行结果为：

```
a: []int{1, 2, 3} a内存地址:0xc42000a0c0
a[0]内存地址：0xc4200182c0
a[1]内存地址：0xc4200182c8
a[2]内存地址：0xc4200182d0

x: []int{1, 2, 3} x内存地址:0xc42000a100
x[0]内存地址：0xc4200182c0
x[1]内存地址：0xc4200182c8
x[2]内存地址：0xc4200182d0

执行后 x:[]int{1, 8, 3}

执行后 a:[]int{1, 8, 3}
```

显而易见**Go 里边函数传参只有值传递一种方式**, 包括 slice/map/chan 在内所有类型, 没有传引用的说法.
