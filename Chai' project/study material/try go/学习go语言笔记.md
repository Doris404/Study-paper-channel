## 学习go语言笔记

> https://tour.go-zh.org/welcome/1

#### 切片

```[]Slice```为一个元素类型为Slice的切片。切片```Slice[1:4]```代表一个包含Slice中下标从1到3的切片

> https://tour.go-zh.org/moretypes/7


切片可以用内建函数```make```来创建，这也是创建动态数组的方式。

```a := make([]int,5)```即创建一个长度为5，元素种类为整形的切片

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

> https://tour.go-zh.org/moretypes/13

为切片追加新的元素是常见操作，使用```append```函数

```s = append(s, 2, 3, 4)```即在s中加入2,3,4这三个元素

> https://tour.go-zh.org/moretypes/15

#### 位移操作

```
package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a)
	a = a << 1
	fmt.Println(a)
}
```
最终输出结果为：
```
2
4
```
```<<```为左移符号，代表乘2，```>>```为右移符号，代表除以2

#### 映射

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

```
> https://tour.go-zh.org/moretypes/19

#### 函数传参

函数也可以作为一个参数传入另外一个函数作为参数

```
package main

import "fmt"

func calculate(fn func(float64,float64) float64, a float64, b float64) float64{
	return fn(a, b)
}

func add(a float64, b float64) float64{
	return a + b 
}

func main(){
	fmt.Println(calculate(add,3,4))
}
```

#### 函数的闭包

闭包：能够读取其他函数内部变量的函数。例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。

> https://tour.go-zh.org/moretypes/26

```
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	var result [20]int
	result[0] = 0
	result[1] = 1	
	x := -1	
	fmt.Println("outside x = ",x)
	return func() int{
		x = x + 1
		fmt.Println("inside x = ",x)
		if (x == 0 || x == 1){
			return result[x]
		}
		result[x] = result[x-1]+result[x-2]
		return result[x]
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

```
输出：

```
outside x = -1
inside x =  0
0
inside x =  1
1
inside x =  2
1
inside x =  3
2
inside x =  4
3
inside x =  5
5
inside x =  6
8
inside x =  7
13
inside x =  8
21
inside x =  9
34

```

#### interface

```
package main

import "fmt"

type I interface {
	M()
	K()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func (t T) K() {
	fmt.Println("string is ",t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
	i.K()
}
```
最终输出：
```
hello
string is  hello
```

一个接口使得不同的数据类型可以用一个函数名进行计算：
```
package main

import "fmt"

type calcu interface{
	print_out ()
	show_type ()
}

type vertex struct{
	X float64
	Y float64
}

type F float64

func (v vertex) print_out(){
	fmt.Println("print_out output : ", "vertex.X: ", v.X, "vertex.Y: ", v.Y)
}

func (v vertex) show_type(){
	fmt.Println("show_type output : ", "vertex")
}

func (f F) print_out(){
	fmt.Println("print_out output : ", "float64 : ", f)
}

func (f F) show_type(){
	fmt.Println("show_type output : ", "float64")
}

func main(){
	var i calcu
	i = vertex{3,4}
	i.print_out()
	i.show_type()
	i = F(2.3)
	i.print_out()
	i.show_type()
}
```
输出：
```
print_out output :  vertex.X:  3 vertex.Y:  4
show_type output :  vertex
print_out output :  float64 :  2.3
show_type output :  float64
```

**类型判断**

```
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
```

#### 信道

```
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```
输出
```
1
2
```
信道(channel)就像一个队列(queue)一样，先进先出.

创建一个信道的语句是：
```
new_channel := make(chan channel_type, column)

其中channel_type即信道里元素的种类，column是信道的大小
```

#### select 语句

```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```
输出：

```
0
1
1
2
3
5
8
13
21
34
quit
```
大多数情况下，select语句与for语句一起出现，以达到多次选择的作用,请看下面这个例子

```
package main

import (
	"fmt"
	"time"
)
func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(1*time.Second)
		c1 <- "one"
	}()
	go func(){
		time.Sleep(1*time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++{
		select{
		case msg1 := <- c1:
			fmt.Println("received",msg1)
		
		case msg2 := <- c2:
			fmt.Println("received",msg2)
		}
	}

}
```
在这个例子中，两个go func()仅仅是将两个channel装满，在后面的for循环中每一次循环中要打印一个channel的消息。打印的顺序是随机的。

>更多学习： https://tour.go-zh.org/concurrency/11

#### 并发

#### 疑惑

- select 







