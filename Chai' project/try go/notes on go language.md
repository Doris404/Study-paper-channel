### This is a note I keep when I am learning the language of go

- go 可以获取系统类型例如
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

- 子函数是否会更改变量值，请看以下三个例子
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