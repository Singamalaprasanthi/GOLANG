//1. Channels

package main 
  
import "fmt"
  
func main() { 
  
    
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel) 
    fmt.Printf("Type of the channel: %T ", mychannel) 
  
    
    mychannel1 := make(chan int) 
    fmt.Println("\nValue of the channel1: ", mychannel1) 
    fmt.Printf("Type of the channel1: %T ", mychannel1) 
} 

$go run main.go
Value of the channel:  <nil>
Type of the channel: chan int 
Value of the channel1:  0xc000094060
Type of the channel1: chan int 

// 2. Creating a channel 

package main 
  
import "fmt"
  
func myfunc(ch chan int) { 
  
    fmt.Println(450 + <-ch) 
} 
func main() { 
    fmt.Println("start Main method") 
    // Creating a channel 
    ch := make(chan int) 
  
    go myfunc(ch) 
    ch <- 50 
    fmt.Println("End Main method") 
} 

$go run main.go
start Main method
510
End Main method

// 3.example
package main 
  
import "fmt"
func myfun(mychnl chan string) { 
  
    for v := 0; v < 5; v++ { 
        mychnl <- "prasanthi"
    } 
    close(mychnl) 
} 
  func main() { 
  
    // Creating a channel 
    c := make(chan string) 
  
    // calling Goroutine 
    go myfun(c) 
  
   
    for { 
        res, ok := <-c 
        if ok == false { 
            fmt.Println("Channel Close ", ok) 
            break
        } 
        fmt.Println("Channel Open ", res, ok) 
    } 
} 

$go run main.go
Channel Open  prasanthi true
Channel Open  prasanthi true
Channel Open  prasanthi true
Channel Open  prasanthi true
Channel Open  prasanthi true
Channel Close  false

// 4.example

package main  
import "fmt"

func main() { 
  mychnl := make(chan string) 
  go func() { 
        mychnl <- "hik"
        mychnl <- "gfg"
        mychnl <- "hari"
        mychnl <- "prasanthi"
        close(mychnl) 
    }() 
  for res := range mychnl { 
        fmt.Println(res) 
    } 
} 

$go run main.go
hik
gfg
hari
prasanthi

//5.example
package main  
import "fmt"

func main() { 
    
    mychnl := make(chan string, 4) 
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Hari"
    mychnl <- "prasanthi"
    fmt.Println("Length of the channel is: ", len(mychnl)) 
}

$go run main.go
Length of the channel is:  4

//6.example
package main 
  
import "fmt"
func main() { 
  
    
    mychnl := make(chan string, 5) 
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "hari"
    mychnl <- "prasanthi"
  
    fmt.Println("Capacity of the channel is: ", cap(mychnl)) 
}
 
$go run main.go
Capacity of the channel is:  5

// 7.example

package main
import "fmt"
func main() {
    messages := make(chan string)

    go func() { messages <- "hello" }()


    msg := <-messages
    fmt.Println(msg)
}

$go run main.go
hello

// 8.example

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c 

	fmt.Println(x, y, x+y)
}

$go run main.go
-5 17 12

// 9.example
package main 
    
import "fmt"
    
func main() { 
  
    // Creating a channel using make() function 
    ch:= make(chan int, 5) 
    fmt.Printf("\nChannel: %d", len(ch)) 
    ch <- 0 
    ch <- 1 
    ch <- 2 
    fmt.Printf("\nChannel: %d", len(ch))
}
 
$go run main.go

Channel: 0
Channel: 3

// 10. example

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 
}

func main() {
	s := []int{5, 2, 4, -9, -2, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c 

	fmt.Println(x, y, x+y)
}

$go run main.go
-11 11 0









