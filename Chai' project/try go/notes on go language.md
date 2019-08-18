## This is a note I keep when I am learning the language of go

### ��go ����ָ�ϡ�
***
#### �ڶ����� �����±ʼ�
- һ��go������
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
ͬʱҲ��õ�·��

- ���ڱ�����ֵ���⣬�Ӻ����ж�һ���������и�ֵ�Ƿ��ı��ⲿ������ֵ
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
����ֵΪ: 
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
����ֵΪ��
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
����ֵΪ��
```go
GOG
```
��һ�ֵ�m���������¶�����һ����ͬ�ڵ�һ��a���µı���a����������������������m�����ڲ������������ΧʧЧ

�ڶ��ֵ�m�����еı���a���ⲿ�ı���a��һ�������������m�����ڲ��ı�a��ֵ�����ⲿҲ��Ч

������ͬ��һ�֣���m�����ڲ����¶���ı�������m����ʧЧ

- strings �� strconv�� 
>���ڴ����ַ����ĳ������������ the way to go 4.7 https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.7.md

