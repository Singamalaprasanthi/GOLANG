// 1. example

package main
import (
    "fmt"
    "os"
)
func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
You can get individual args with normal indexing.

    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

// 2. example

package main  
  
import (  
    "fmt"  
    "os"  
)  
func main() {  
    var s, arg string  
    for i := 1; i < len(os.Args); i++ {  
        s += arg + os.Args[i]+" "  
    }  
    fmt.Println(s)  
}  

go build ProgramName.go  
./ProgramName Tom Dick Harry 

Tom Dick Harry   

// 3. example

package main  
import "os"  
import "fmt"  
func main() {  
    arumentWithPath := os.Args 
    arumentSlice:= os.Args[1:] 
    arumentAt2 := os.Args[2] 
    fmt.Println(arumentWithPath)  
    fmt.Println(arumentSlice)  
    fmt.Println(arumentAt2)  
}  

[Tom Dick Harry]
Dick


// 4. example

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}

$ go run hello.go
Hello World!

// 5. example

// flag.go
package main

import (
    "flag"
    "fmt"
)

func main() {
    var name string
    flag.StringVar(&name, "opt", "", "Usage")

    flag.Parse()

    fmt.Println(name)
}

$ go run hello.go -opt option
option

$ go run hello.go
Hello World!

$ go run hello.go -name Gopher
Hello Gopher

// 6. example

// args.go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args)
    fmt.Println(os.Args[1])
}

$ go build -o args args.go 
$ ./args prasanthi
[./args prasanthi]
prasanthi

// 7. example

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(len(os.Args), os.Args)
}

$go run main.go
1 [/tmp/go-build486280812/command-line-arguments/_obj/exe/main]

// 8.example

package main

import ("fmt"
        "os"
)

func main() {
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

$ go run test.go 1 2 3 4 5
[/tmp/go-build162373819/command-line-arguments/_obj/exe/modbus 1 2 3 4 5]
[1 2 3 4 5]
3


// 9.example

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "default value", "a string for description")
    flag.Parse()
    fmt.Println("word:", *wordPtr)

}

go run main.go -word=hello
 word: hello

// 10.example

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "default value", "a string for description")
    flag.Parse()
    fmt.Println("word:", *wordPtr)

}

go run main.go -word=welcome
 word : welcome












































