## This is a note I keep when I am learning the language of go

### 《go 入门指南》
***
#### 第二部分 第四章笔记
- 一个go的例子
```go
package main

import (
	"fmt"
   "runtime"
	"os"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
```
运行这段代码会得到
```go
The operating system is: windows
```
同时也会得到路径

- 关于变量赋值问题，子函数中对一个变量进行赋值是否会改变外部变量的值
```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   a := "O"
   print(a)
}
```
返回值为: 
```go
GOG
```
```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() {
   print(a)
}

func m() {
   a = "O"
   print(a)
}
```
返回值为：
```go
GOO
```
```go
package main

var a string

func main() {
   a = "G"
   print(a)
   f1()
}

func f1() {
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}
```
返回值为：
```go
GOG
```
第一种的m函数是重新定义了一个不同于第一个a的新的变量a，这个变量的作用域仅仅在m函数内部，出了这个范围失效

第二种的m函数中的变量a和外部的变量a是一个变量，因此在m函数内部改变a的值，到外部也有效

第三种同第一种，在m函数内部重新定义的变量出了m函数失效

- strings 和 strconv包 
>用于处理字符串的常规操作，来自 the way to go 4.7 https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.7.md

