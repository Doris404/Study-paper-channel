### some interesting problem I met when study the language of go
####  exit status 2
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