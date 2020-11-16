// 1. Simple for loop

package main 
  
import "fmt"
  

func main() { 
      
      for i := 0; i < 4; i++{ 
      fmt.Printf("prasanthi\n")   
    } 
    
} 

$go run main.go
prasanthi
prasanthi
prasanthi
prasanthi

// 2. For loop as Infinite Loop

package main 
  
import "fmt"
  
 func main() { 
      
    // Infinite loop 
    for { 
      fmt.Printf("prasanthi\n")   
    } 
    
}

$go run main.go
prasanthi
prasanthi
prasanthi
prasanthi
prasanthi
prasanthi

// 3. for loop as while Loop

package main 
  
import "fmt"
  
 func main() { 
      
    i:= 0 
    for i < 3 { 
       i += 2 
    } 
  fmt.Println(i)  
} 

$go run main.go
4

// 4. Simple range in for loop

package main 
  
import "fmt"
  
 func main() { 
      
    rvariable:= []string{"spr", "prasanthi", "reddy"}  
      
    for i, j:= range rvariable { 
       fmt.Println(i, j)  
    } 
    
} 

$go run main.go
0 spr
1 prasanthi
2 reddy

// 5. Using for loop for strings

package main 
  
import "fmt"
  
 func main() { 
      
    // String as a range in the for loop 
    for i, j:= range "XabCd" { 
       fmt.Printf("The index number of %U is %d\n", j, i)  
    } 
    
} 

$go run main.go
The index number of U+0058 is 0
The index number of U+0061 is 1
The index number of U+0062 is 2
The index number of U+0043 is 3
The index number of U+0064 is 4

// 6. For Maps

package main 
  
import "fmt"
  
func main() { 
      
    mmap := map[int]string{ 
        22:"spr", 
        33:"prasanthi", 
        44:"reddy", 
    } 
    for key, value:= range mmap { 
       fmt.Println(key, value)  
    } 
    
} 

$go run main.go
22 spr
33 prasanthi
44 reddy

// 7. For Channel

package main 
  
import "fmt"
  
func main() { 
chnl := make(chan int) 
    go func(){ 
        chnl <- 100 
        chnl <- 1000 
       chnl <- 10000 
       chnl <- 100000 
       close(chnl) 
    }() 
    for i:= range chnl { 
       fmt.Println(i)  
    } 
    
} 

$go run main.go
100
1000
10000
100000

// 8. Count with a basic while loop

package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 5 {
		fmt.Println("i =", i)
		i++
	}
}

$go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4

// 9. while loop in go

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


$go run main.go
1024

// 10 .do while in go

package main

import "fmt"

func main() {

	fmt.Println("Demo 1")
	Demo1()

	fmt.Println("\nDemo 2")
	Demo2()

	fmt.Println("\nDemo 3")
	Demo3()

	fmt.Println("\nDemo 4")
	Demo4()

}

func Demo1() {
	i := 1
	for {
		fmt.Printf("%3d", i)
		i++
		if i == 10 {
			break
		}
	}
}

func Demo2() {
	ch := 'a'
	for {
		fmt.Printf("%3c", ch)
		ch++
        if ch == 'z' {
			break
		}
	}
}

func Demo3() {
	var total int = 0
	i := 1
	for {
		total += i
		i++
        if i == 10 {
			break
		}
	}
	fmt.Println("Total: ", total)
}

func Demo4() {
	var counter int = 0
	i := 1
	for {
		if i%2 == 0 {
			counter++
		}
		i++
		if i == 10 {
			break
		}
	}
	fmt.Println("Counter: ", counter)
}

$go run main.go
Demo 1
  1  2  3  4  5  6  7  8  9
Demo 2
  a  b  c  d  e  f  g  h  i  j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y
Demo 3
Total:  45

Demo 4
Counter:  4





































