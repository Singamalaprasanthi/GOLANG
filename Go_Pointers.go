// 1. Pointers

package main
import "fmt"
func main() {
   var a int = 40  
   fmt.Printf("Address of a variable: %x\n", &a  )
}

$go run main.go
Address of a variable: c00002c008

// 2.example

package main
import "fmt"
func main() {
   var a int = 10   
   var ip *int      
   ip = &a  
   fmt.Printf("Address of a variable: %x\n", &a  )
   fmt.Printf("Address stored in ip variable: %x\n", ip )
   fmt.Printf("Value of *ip variable: %d\n", *ip )
}
$go run main.go
Address of a variable: c00002c008
Address stored in ip variable: c00002c008
Value of *ip variable: 10


//3.example

package main
import "fmt"
func main() {
   var  ptr *int
   fmt.Printf("The value of ptr is : %x\n", ptr  )
}
$go run main.go
The value of ptr is : 0

// 4. example

package main 
import "fmt"
func main() { 
    x := 0xFF 
    y := 0x9C 
    fmt.Printf("Type of variable x is %T\n", x) 
    fmt.Printf("Value of x in hexdecimal is %X\n", x) 
    fmt.Printf("Value of x in decimal is %v\n", x) 
    fmt.Printf("Type of variable y is %T\n", y) 
    fmt.Printf("Value of y in hexdecimal is %X\n", y) 
    fmt.Printf("Value of y in decimal is %v\n", y)     
} 
$go run main.go
Type of variable x is int
Value of x in hexdecimal is FF
Value of x in decimal is 255
Type of variable y is int
Value of y in hexdecimal is 9C
Value of y in decimal is 156

// 5. example

package main 
import "fmt"
func main() { 
    var x int = 2356 
      
    var p *int
      
    p = &x 
  
    fmt.Println("Value stored in x = ", x) 
    fmt.Println("Address of x = ", &x) 
    fmt.Println("Value stored in variable p = ", p) 
} 
$go run main.go
Value stored in x =  2356
Address of x =  0xc00002c008
Value stored in variable p =  0xc00002c008

// 6. examples

package main 
import "fmt"
func main(){
    var y = 324
    var p = &y 
    fmt.Println("Value stored in y before changing = ", y)
    fmt.Println("Address of y = ", &y) 
    fmt.Println("Value stored in pointer variable p = ", p) 
    fmt.Println("Value stored in y(*p) Before Changing = ", *p)
    *p = 500 
     fmt.Println("Value stored in y(*p) after Changing = ",y) 
  
} 
$go run main.go
Value stored in y before changing =  324
Address of y =  0xc0000b6020
Value stored in pointer variable p =  0xc0000b6020
Value stored in y(*p) Before Changing =  324
Value stored in y(*p) after Changing =  500


// 7. example

package main
import (  
    "fmt"
)
func change(val *int) {  
    *val = 30
}
func main() {  
    a := 40
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
$go run main.go
value of a before function call is 40
value of a after function call is 30

// 8. example

package main
import (  
    "fmt"
)
func helloworld() *int {  
    i := 20
    return &i
}
func main() {  
    d := helloworld()
    fmt.Println("Value of d", *d)
}
$go run main.go
Value of d 20

// 9. example

package main
import (  
    "fmt"
)
func main() {  
    b := 250
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)
}
$go run main.go
address of b is 0xc000080010
value of b is 250
new value of b is 251

// 10. example 

package main
import (  
    "fmt"
)
func modify(arr *[3]int) {  
    arr[1] = 89
}
func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
$go run main.go
[89 89  91]