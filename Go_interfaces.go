// 1.Interface

package main 
import "fmt"
type tank interface { 
    Tarea() float32 
    Volume() float32 
} 
  
type myvalue struct { 
    radius float32 
    height float32 
} 
  
func (m myvalue) Tarea() float32 { 
  
    return 2*m.radius*m.height + 
        2*3.14*m.radius*m.radius 
} 
  
func (m myvalue) Volume() float32 { 
  
    return 3.14 * m.radius * m.radius * m.height 
} 
  
// Main Method 
func main() { 
    var t tank 
    t = myvalue{10, 20} 
    fmt.Println("Area of tank :", t.Tarea()) 
    fmt.Println("Volume of tank:", t.Volume()) 
} 

$go run main.go
Area of tank : 1028
Volume of tank: 6280

// 2. example

package main
import (  
    "fmt"
)
type SalaryCalculator interface {  
    CalculateSalary() int
}
type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}
type Contract struct {  
    empId    int
    basicpay int
}
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}
func (c Contract) CalculateSalary() int {  
    return c.basicpay 
}

func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{
        empId:    1,
        basicpay: 1000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 2000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}

$go run main.go
Total Expense Per Month $6050

// 3. example

package main
import (  
    "fmt"
)
type Worker interface {  
    Work()
}
type Person struct {  
    name string
}
func (p Person) Work() {  
    fmt.Println(p.name, "is working")
}
func describe(w Worker) {  
    fmt.Printf("Interface type %T value %v\n", w, w)
}
func main() {  
    p := Person{
        name: "prasanthi",
    }
    var w Worker = p
    describe(w)
    w.Work()
}

$go run main.go
Interface type main.Person value {prasanthi}
prasanthi is working

/ / 4.Empty interface

package main
import (  
    "fmt"
)
func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}
func main() {  
    s := "Hello"
    describe(s)
    i := 67
    describe(i)
    strt := struct {
        name string
    }{
        name: "prasanthi S",
    }
    describe(strt)
}

$go run main.go
Type = string, value = Hello
Type = int, value = 67
Type = struct { name string }, value = {prasanthi S}

// 5. example

package main
import (  
    "fmt"
)
type VowelsFinder interface {  
    FindVowels() []rune
}
type MyString string
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}
func main() {  
    name := MyString("prasanthi reddy")
    var v VowelsFinder
    v = name 
    fmt.Printf("Vowels are %c", v.FindVowels())

}

$go run main.go
Vowels are [a  a i e]

// 6. example

package main 
import "fmt"
func myfun(a interface{}) { 
    val := a.(string) 
    fmt.Println("Value: ", val) 
} 
func main() { 
  
    var val interface { 
    } = "prasanthi"
      
    myfun(val) 
} 

$go run main.go
Value:  prasanthi

// 7. example

package main 
import "fmt"
func myfun(a interface{}) { 
    value, ok := a.(float64) 
    fmt.Println(value, ok) 
} 
func main() { 
  
    var a1 interface { 
    } = 90.01
  
    myfun(a1) 
  
    var a2 interface { 
    } = "prasanthi"
  
    myfun(a2) 
} 

$go run main.go
90.01 true
0 false

// 8.example

package main 
import "fmt"
type tank interface { 
    Tarea() float64 
    Volume() float64 
} 
func main() { 
    var t tank 
    fmt.Println("Value of the tank interface is: ", t) 
    fmt.Printf("Type of the tank interface is: %T ", t) 
} 

$go run main.go
Value of the tank interface is:  <nil>
Type of the tank interface is: <nil> 

// 9. example

package main
import "fmt"
type Guitarist interface {
    
    PlayGuitar()
}

type BaseGuitarist struct {
    Name string
}

type AcousticGuitarist struct {
    Name string
}

func (b BaseGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
    var player BaseGuitarist
    player.Name = "prasanthi"
    player.PlayGuitar()

    var player2 AcousticGuitarist
    player2.Name = "Ringo"
    player2.PlayGuitar()
}

$go run main.go
prasanthi plays the Bass Guitar
Ringo plays the Acoustic Guitar

// 10. example

package main
import (
    "fmt"
)
func myFunc(a interface{}) {
    fmt.Println(a)
}
func main() {
    var my_age int
    my_age = 23

    myFunc(my_age)
}

$go run main.go
23


