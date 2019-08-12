### some interesting problem  and fact I met when study the language of go
####  exit status 2 problem
This kind of problem has happened many times. The reason is not clearly found yet. Here are some situations that the problem can happen.
- When you try the example on the website of go example https://gobyexample.com/timers:
``` go
package main
import "fmt"

func main() {
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	// close(queue)
	for elem := range queue{
		fmt.Println(elem)
	}
}
```
It will return the follows:
```go
one
two
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        C:/Users/value.go:11 +0x113
```
- sometimes the problem of exit status 2 happens because of overflow
``` go
package main

import (
	"fmt"
	"time"
)

func main() {
	burstyLimiter := make (chan time.Time,3)

	for i:=0; i<3; i++{
		burstyLimiter <- time.Now()
		fmt.Println(<-burstyLimiter)
	}

	fmt.Println(<-burstyLimiter)
}
```
The result is as follows:
``` go
2019-08-13 07:53:45.9807829 +0800 CST m=+0.020969301
2019-08-13 07:53:46.0276332 +0800 CST m=+0.067819601
2019-08-13 07:53:46.0286316 +0800 CST m=+0.068818001
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        C:/Users/ÀîÏþÍ©/Desktop/value.go:27 +0x17c
exit status 2
```