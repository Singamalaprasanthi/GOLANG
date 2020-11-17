// 1. go mutex

package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    balance int
)

func init() {
    balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
    balance += value
    mutex.Unlock()
    wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
    balance -= value
    mutex.Unlock()
    wg.Done()
}

func main() {
    fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
    go withdraw(700, &wg)
    go deposit(500, &wg)
    wg.Wait()

    fmt.Printf("New Balance %d\n", balance)
}

$go run main.go
Go Mutex Example
Depositing 500 to account with balance: 1000
Withdrawing 700 from account with balance: 1500
New Balance 800

// 2. example

package main  
import (  
   "sync"  
   "time"  
   "math/rand"  
   "fmt"  
)  
var wait sync.WaitGroup  
var count int  
var mutex sync.Mutex  
func  increment(s string)  {  
   for i :=0;i<10;i++ {  
      mutex.Lock()  
      x := count  
      x++;  
      time.Sleep(time.Duration(rand.Intn(10))*time.Millisecond)  
      count = x;  
      fmt.Println(s, i,"Count: ",count)  
      mutex.Unlock()  
        
   }  
   wait.Done()  
     
}  
func main(){  
   wait.Add(2)  
   go increment("foo: ")  
   go increment("bar: ")  
   wait.Wait()  
   fmt.Println("last count value " ,count)  
}  

$go run main.go
bar:  0 Count:  1
bar:  1 Count:  2
bar:  2 Count:  3
bar:  3 Count:  4
bar:  4 Count:  5
bar:  5 Count:  6
bar:  6 Count:  7
bar:  7 Count:  8
bar:  8 Count:  9
bar:  9 Count:  10
foo:  0 Count:  11
foo:  1 Count:  12
foo:  2 Count:  13
foo:  3 Count:  14
foo:  4 Count:  15
foo:  5 Count:  16
foo:  6 Count:  17
foo:  7 Count:  18
foo:  8 Count:  19
foo:  9 Count:  20
last count value  20

// 3. example

package main

import (
	"fmt"
	"sync"
	"time"
)


type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}


func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}


func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

$go run main.go
1000

// 4.example

package main
import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)
func main() {

    var state = make(map[int]int)


    var mutex = &sync.Mutex{}


    var readOps uint64
    var writeOps uint64


    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
  

                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)


                time.Sleep(time.Millisecond)
            }
        }()
    }


    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }


    time.Sleep(time.Second)


    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}

$go run main.go
readOps: 85535
writeOps: 8560
state: map[0:50 2:70 3:87 1:2 4:86]

// 5.example

package main

import (
	"fmt"
	"sync"
)


func main() {
	var b sync.Mutex
	
	b.Lock()
	b.Lock()
	fmt.Println("This never executes as we are in deadlock") 
}

$go run main.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0xc42000e1fc)
	/usr/lib/golang/src/runtime/sema.go:62 +0x34
sync.(*Mutex).Lock(0xc42000e1f8)
	/usr/lib/golang/src/sync/mutex.go:87 +0x9d
main.main()
	/home/cg/root/4724938/main.go:13 +0x59
exit status 2



















