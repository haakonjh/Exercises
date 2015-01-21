//go run threading_chan.go
package main
import (
. "fmt"
"runtime"
)


var i int = 0
var buffer = make(chan int, 1)
var done_add = make(chan int, 1)
var done_sub = make(chan int, 1)

func thread_1() {
	for j := 0; j < 1000000; j++{
		<-buffer
		i++	
		buffer<-1	
	}
	done_add<-1
}
func thread_2() {
	for j := 0; j < 1000000; j++{
		<-buffer
		i--
		buffer<-1
	}
	done_sub<-1
}
func main() {

runtime.GOMAXPROCS(runtime.NumCPU()) 
buffer<-1
go thread_1()
go thread_2()
<-done_add
<-done_sub
Println(i)
}
