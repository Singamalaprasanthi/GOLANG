//1.Defer and Panic

package main
import "fmt"
func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}
func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

$go run main.go
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.

//2.example

package main
import (  
    "fmt"
)

func recoverFullName() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverFullName()
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}

$go run main.go
recovered from  runtime error: last name cannot be nil
returned normally from main
deferred call in main

//3.Example

package main 
import ( 
    "fmt"
) 
  
func entry(lang *string, aname *string) { 
  
    
    defer fmt.Println("Defer statement in the entry fun ") 
  
    
    if lang == nil { 
        panic("Error: Language cannot be nil") 
    } 
      
    
    if aname == nil { 
        panic("Error: Author name cannot be nil") 
    } 
  
    
    fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname) 
} 
  
func main() { 
  
    A_lang := "GO Language"
  
    
    defer fmt.Println("Defer statement in the Main fun") 
  
    entry(&A_lang, nil) 
} 

$go run main.go
Defer statement in the entry fun
Defer statement in the Main fun
panic: Error: Author name cannot be nil

goroutine 1 [running]:
main.entry(0xc42005df48, 0x0)
	/home/cg/root/9163813/main.go:22 +0x20d
main.main()
	/home/cg/root/9163813/main.go:42 +0xda
exit status 2

//4.example

package main
import "os"
func main() {
    panic("a problem")
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}

$go run main.go
panic: a problem
goroutine 1 [running]:
main.main()
	/home/cg/root/9163813/main.go:4 +0x64
exit status 2

//5.example

package main
import "fmt"
func main() {
	defer fmt.Println("prasanthi")

	fmt.Println("hello")
}

$go run main.go
hello
prasanthi

//6.example

package main
import (
    "fmt"
    "os"
)
func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

$go run main.go
creating
writing
closing

//7.example

package main 
import "fmt"

func mul(a1, a2 int) int { 
    res := a1 * a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
  
func show() { 
    fmt.Println("Hello!, prasanthi") 
} 

func main() { 
    mul(23, 45) 
  
    defer mul(23, 56) 
    show() 
} 

$go run main.go
Result:  1035
Hello!, prasanthi
Result:  1288

//8.example

package main 
import "fmt"
func add(a1, a2 int) int { 
    res := a1 + a2 
    fmt.Println("Result: ", res) 
    return 0 
} 

func main() { 
  
    fmt.Println("Start") 
  
    defer fmt.Println("End") 
    defer add(34, 56) 
    defer add(10, 10) 
} 

$go run main.go
Start
Result:  20
Result:  90
End

//9.example

package main

import "fmt"

func foo() {
  defer fmt.Println("Done!")
  fmt.Println(" who knows what?")
}

func main() {
  foo()
}

$go run main.go
who knows what?
Done!

//10.example

package main

import "fmt"

func foo() {
  defer fmt.Println("Hi!")
  defer fmt.Println("How are you?")
  fmt.Println("what about office work?")
}

func main() {
  foo()
}

$go run main.go
what about office work?
How are you?
Hi!

