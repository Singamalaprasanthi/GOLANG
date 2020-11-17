// 1. go routine

package main 
  
import "fmt"
  
func display(str string) { 
    for w := 0; w < 6; w++ { 
        fmt.Println(str) 
    } 
} 
  
func main() { 
  
    // Calling Goroutine 
    go display("welcome") 
   display("prasanthi reddy") 
} 

$go run main.go
prasanthi reddy
prasanthi reddy
prasanthi reddy
prasanthi reddy
prasanthi reddy
prasanthi reddy

// 2.example

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
Welcome
prasanthi
prasanthi
Welcome
prasanthi
Welcome
Welcome
prasanthi

// 3. Anonymous Goroutine 

package main 
  
import ( 
    "fmt"
    "time"
) 
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

// 4.example

package main

import (  
    "fmt"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
    fmt.Println("main function")
}

$go run main.go
main function

// 5. example

package main

import (  
    "fmt"
    "time"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}

$go run main.go
Hello world goroutine
main function

// 6.Multiple go routines

package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}

$go run main.go
1 a 2 3 b 4 c 5 d e main terminated

// 7. example

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
	go say("prasanthi")
	say("hello")
}


$go run main.go
prasanthi
hello
prasanthi
hello
prasanthi
hello
prasanthi
hello

// 8. example

package main
import (
    "fmt"
    "time"
)
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {

    f("direct")


    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")


    time.Sleep(time.Second)
    fmt.Println("done")
}

$go run main.go
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done

// 9. example

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
    fmt.Println("Goroutine Tutorial")

    compute(10)
    compute(10)

    var input string
    fmt.Scanln(&input)

}

$go run main.go
Goroutine Tutorial
0
1
2
3
4
5
6
7
8

// 10.example

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
    fmt.Println("Goroutine ")

    go compute(10)
    go compute(10)
}

$go run main.go
Goroutine 




















