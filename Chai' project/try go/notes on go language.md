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