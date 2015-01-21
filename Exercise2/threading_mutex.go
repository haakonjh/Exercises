//go run threading_mutex.go
package main
import (
	"fmt"
    "runtime"
    "sync"
)

var i int=0
var mutex = &sync.Mutex{}
func thread_1() {
	for j := 0; j < 1000000; j++{
		mutex.Lock()
		i++;
		mutex.Unlock()
	}
}
func thread_2() {
	for j := 0; j < 1000000; j++{
		mutex.Lock()
		i--;
		mutex.Unlock()
	}
}
func main() {


runtime.GOMAXPROCS(runtime.NumCPU()) 

go thread_1() 
go thread_2()

fmt.Println(i)

}