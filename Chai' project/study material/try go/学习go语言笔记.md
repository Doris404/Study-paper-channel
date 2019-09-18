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


> end point:https://tour.go-zh.org/methods/4








