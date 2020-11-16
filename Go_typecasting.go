// 1.  Explicit typecasting

package main 
  
import "fmt"
  
func main() { 
  
    var totalsum int = 200
    var number int = 10 
    var avg float32 
  
    // explicit type conversion 
    avg = float32(totalsum) / float32(number) 
   
    fmt.Printf("Average = %f\n", avg) 
} 

$go run main.go
Average = 20.000000

// 2. example

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 4, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

$go run main.go
4 5 6

// 3. example

package main

import "fmt"

func main() {
	v := 42.0 
	fmt.Printf("v is of type %T\n", v)
}

$go run main.go
v is of type float64

// 4. example

package main

import "fmt"

func main() {
   var sum int = 15
   var count int = 4
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("Value of mean : %f\n",mean)
}

$go run main.go
Value of mean : 3.750000

// 5 . example

package main
 
import (
    "fmt"
)
 
func main() {
    var a int = 40
    f := float64(a)
    fmt.Println(f)     
}

$go run main.go
40

// 6 . example

package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    var s string = "42"
    v, _ := strconv.Atoi(s)       
    fmt.Println(v)    // 42
     
    var i int = 42
    str := strconv.Itoa(i)       
    fmt.Println(str) 
}

$go run main.go
42
42

// 7 . example

package main
 
import (
    "fmt"
)
 
func main() {
    f := 12.34567
    i := int(f)  // loses precision
    fmt.Println(i)      
     
    ii := 34
    ff := float64(ii)
     
    fmt.Println(ff)    
}

$go run main.go
12
34

//  8. example

package main
 
import (
    "fmt"
)
 
func main() {
    var s string = "Hello"
    var b []byte = []byte(s)     
    fmt.Println(b)  
     
    ss := string(b)             
     
    fmt.Println(ss)     
}

$go run main.go
[72 101 108 108 111]
Hello

//  9. example

package main
 
import (
    "fmt"
)
 
func main() {
    a := 6/3        
    f := 6.3/3  
     
    fmt.Println(a, f)    
}

$go run main.go
2 2.1

//  10. example

package main
 
import (
    "fmt"
)
 
func main() {
    f := 12.34567
    i := int(f)  // loses precision
    fmt.Println(i)      
     
    ii := 34.23
    ff := float64(ii)
     
    fmt.Println(ff)    
}

$go run main.go
12
34.23



























