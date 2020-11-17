
// 1.Go Select

package main
import "fmt"
func main() {
var c1, c2, c3 chan int
var i1, i2 int
select {
case i1 = <-c1:
fmt.Printf("received ", i1, " from c1\n")
case c2 <- i2:
fmt.Printf("sent ", i2, " to c2\n")
case i3, ok := (<-c3): // same as: i3, ok := <-c3
if ok {
fmt.Printf("received ", i3, " from c3\n")
} else {
fmt.Printf("c3 is closed\n")
}
default:
fmt.Printf("communication is not there\n")
}
}

$go run main.go
communication is not there

// 2.Example

package main
import (
    "fmt"
    "time"
)
func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "three"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "five"
    }()


    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

$go run main.go
received three
received five

// 3.example

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

$go run main.go
0
1
1
2
3
5
8
13
21
34
quit

// 4.example

package main 
   
import("fmt"
 "time") 
       
    
    func portal1(channel1 chan string) { 
  
        time.Sleep(3*time.Second) 
        channel1 <- "Welcome to channel 1"
    } 
       
    
     func portal2(channel2 chan string) { 
  
        time.Sleep(9*time.Second) 
        channel2 <- "Welcome to channel 2"
    } 

func main(){ 
       
    
   R1:= make(chan string) 
   R2:= make(chan string) 
      
   
   go portal1(R1) 
   go portal2(R2) 
  
   select{ 
  
        
       case op1:= <- R1: 
       fmt.Println(op1) 
   
       
       case op2:= <- R2: 
       fmt.Println(op2) 
   } 
      
} 

$go run main.go
Welcome to channel 1

// 5.example

package main 
import "fmt"
func main() { 
    mychannel:= make(chan int) 
   select{ 
     case <- mychannel:  
       
     default:fmt.Println("Not found") 
}    
} 

$go run main.go
Not found

// 6.example

package main 
import "fmt"
      
    func portal1(channel1 chan string){ 
        for i := 0; i <= 3; i++{ 
            channel1 <- "Welcome to channel 1"
        } 
          
    } 
      
   
     func portal2(channel2 chan string){ 
        channel2 <- "Welcome to channel 2"
    } 
  
func main() { 
      
    
   R1:= make(chan string) 
   R2:= make(chan string) 

   go portal1(R1) 
   go portal2(R2) 
     
   
   select{ 
       case op1:= <- R1: 
       fmt.Println(op1) 
       case op2:= <- R2: 
       fmt.Println(op2) 
   } 
} 

$go run main.go
Welcome to channel 2

// 7.example

package main
import (  
    "fmt"
    "time"
)

func server1(ch chan string) {  
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {  
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}
func main() {  
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}

$go run main.go
from server2

// 8.example

package main

import (  
    "fmt"
    "time"
)

func process(ch chan string) {  
    time.Sleep(10500 * time.Millisecond)
    ch <- "process successful"
}

func main() {  
    ch := make(chan string)
    go process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default:
            fmt.Println("no value received")
        }
    }

}

$go run main.go
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received

// 9.example

package main

import "fmt"

func main() {  
    var ch chan string
    select {
    case v := <-ch:
        fmt.Println("received value", v)
    default:
        fmt.Println("default case executed")

    }
}

$go run main.go
default case executed

// 10.example

package main

import "fmt"

func main() {  
    ch := make(chan string)
    select {
    case <-ch:
    default:
        fmt.Println("case executed")
    }
}

$go run main.go
case executed


