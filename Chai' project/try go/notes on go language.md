### This is a note I keep when I am learning the language of go

- go ���Ի�ȡϵͳ��������
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
������δ����õ�
```go
The operating system is: windows
```