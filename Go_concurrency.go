
// 1.Concurrency

package main 
  
import ( 
    "fmt"
    "time"
) 
  
func display(str string) { 
    for w := 0; w < 6; w++ { 
        time.Sleep(1 * time.Second) 
        fmt.Println(str) 
    } 
} 
  
func main() { 
  
    // Calling Goroutine 
    go display("Welcome") 
  
    display("prasanthi") 
} 

$go run main.go
Welcome
prasanthi
prasanthi
Welcome
prasanthi
Welcome
prasanthi
Welcome
Welcome
prasanthi
Welcome
prasanthi

// 2.Example

package main 
  
import ( 
    "fmt"
    "time"
) 
  
// Main function 
func main() { 
  
    fmt.Println("Welcome!! to Main function") 
  
    go func() { 
  
        fmt.Println("Welcome!! to prasanthi") 
    }() 
  
    time.Sleep(1 * time.Second) 
    fmt.Println("GoodBye!! to Main function") 
} 

$go run main.go
Welcome!! to Main function
Welcome!! to prasanthi
GoodBye!! to Main function

// 3.example

package main
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

$go run main.go
world
hello
hello
world
world
hello
hello
world
world
hello

// 4.example

package main  
import (  
   "fmt"  
   "time"  
   "sync"  
)  
var wg = sync.WaitGroup{}  
  
func main() {  
   wg.Add(2)  
   go fun1()  
   go fun2()  
   wg.Wait()  
}  
func fun1(){  
   for  i:=0;i<10;i++{  
      fmt.Println("fun1,  ->",i)  
      time.Sleep(time.Duration(5*time.Millisecond))  
   }  
   wg.Done()  
}  
func fun2(){  
   for i:=0;i<10;i++{  
      fmt.Println("fun2,  ->",i)  
      time.Sleep(time.Duration(10*time.Millisecond))  
   }  
   wg.Done()  
}
  
$go run main.go
fun2,  -> 0
fun1,  -> 0
fun1,  -> 1
fun2,  -> 1
fun1,  -> 2
fun1,  -> 3
fun2,  -> 2
fun1,  -> 4
fun1,  -> 5
fun2,  -> 3
fun1,  -> 6
fun1,  -> 7
fun2,  -> 4
fun1,  -> 8
fun1,  -> 9
fun2,  -> 5
fun2,  -> 6
fun2,  -> 7
fun2,  -> 8
fun2,  -> 9

// 5.example

package main
import (
    "fmt"
    "time"
)

func compute(value int) {
    for i := 0; i < value; i++ {
        time.Sleep(time.Second)
        fmt.Println(i)
    }
}

func main() {
    fmt.Println("prasanthi")
    compute(10)
    compute(10)
    var input string
    fmt.Scanln(&input)

}

$go run main.go
prasanthi
0
1
2
3
4
5
6
7
8