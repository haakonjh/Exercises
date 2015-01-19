package main
import (
. "fmt"
"runtime"
"time"
)

func thread_1(i* int,lock* Mutex) {
	for j := 0; j < 1000000; j++{
		lock.Lock()
		i++;
		lock.Unlock()
	}
}
func thread_2(i* int,lock* Mutex) {
	for j := 0; j < 1000000; j++{
		lock.Lock()
		i--;
		lock.Unlock()
	}
}
func main() {
var i int=0
var lock Mutex
runtime.GOMAXPROCS(runtime.NumCPU()) 

go thread_1(&i,&lock) 
go thread_2(&i,&Mutex)

Println(i)
}
