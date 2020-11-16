
// EXAMPLE 1 :

package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 23, firstName: "SINGAMALA", lastName: "PRASANTHI", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
}  

$go run main.go

{SINGAMALA PRASANTHI 23}
SINGAMALA

// EXAMPLE 2 :

package main  
import (  
   "fmt"  
)  
type person struct {  
   fname string  
   lname string}  
type employee struct {  
   person  
   empId int  
}  
func (p person) details() {  
   fmt.Println(p, " "+" I am a person")  
}  
func (e employee) details() {  
   fmt.Println(e, " "+"I am a employee")  
}  
func main() {  
   p1 := person{"PRASANTHI", "REDDY"}  
   p1.details()  
   e1 := employee{person:person{"GEETHA", "TURAKA"}, empId: 11}  
   e1.details()  
}  

$go run main.go

{PRASANTHI REDDY}   I am a person
{{GEETHA TURAKA} 11}  I am a employee

 //EXAMPLE 3 :

// Accessing Structure Members


package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books    
 
  
   Book1.title = "Go Programming"
   Book1.author = "Prasanthi"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

  
   Book2.title = "java programming"
   Book2.author = " Geetha"
   Book2.subject = "java Programming Tutorial"
   Book2.book_id = 6495700
 
   
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

  
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

$go run main.go

Book 1 title : Go Programming
Book 1 author : Prasanthi
Book 1 subject : Go Programming Tutorial
Book 1 book_id : 6495407
Book 2 title : java programming
Book 2 author :  Geetha
Book 2 subject : java Programming Tutorial
Book 2 book_id : 6495700

// EXAMPLE 4 :

// Structures as Function Arguments

package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books    
 
   Book1.title = "Go Programming"
   Book1.author = "prasanthi"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407
  
   Book2.title = "java programming"
   Book2.author = "Zara Ali"
   Book2.subject = "java programming Tutorial"
   Book2.book_id = 6495700
 
   
   printBook(Book1)
   printBook(Book2)
}
func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}

$go run main.go

Book title : Go Programming
Book author : prasanthi
Book subject : Go Programming Tutorial
Book book_id : 6495407
Book title : java programming
Book author : Zara Ali
Book subject : java programming Tutorial
Book book_id : 6495700

// EXAMPLE 5 :

//  Pointers to Structures
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books   
   
   Book1.title = "Go Programming"
   Book1.author = "prasanthi"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   
   Book2.title = "java programming"
   Book2.author = "Zara Ali"
   Book2.subject = "java programming  Tutorial"
   Book2.book_id = 6495700
 
   
   printBook(&Book1)
   printBook(&Book2)
}
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}

$go run main.go

Book title : Go Programming
Book author : prasanthi
Book subject : Go Programming Tutorial
Book book_id : 6495407
Book title : java programming
Book author : Zara Ali
Book subject : java programming  Tutorial
Book book_id : 6495700

// EXAMPLE 6 :

package main 
  
import "fmt"
  
type Car struct { 
    Name, Model, Color string 
    WeightInKg         float64 
} 
  
func main() { 
    c := Car{Name: "Ferrari", Model: "GTC4", 
            Color: "Red", WeightInKg: 1920} 
    fmt.Println("Car Name: ", c.Name) 
    fmt.Println("Car Color: ", c.Color) 
    c.Color = "Black"
    fmt.Println("Car: ", c) 
} 

$go run main.go

Car Name:  Ferrari
Car Color:  Red
Car:  {Ferrari GTC4 Black 1920}

// EXAMPLE 7 :

package main 
  
import "fmt"
type Employee struct { 
    firstName, lastName string 
    age, salary int
} 
  
func main() { 
  emp8 := &Employee{"prasanthi", "Anderson", 23, 6000} 
    fmt.Println("First Name:", (*emp8).firstName) 
    fmt.Println("Age:", (*emp8).age) 
} 

$go run main.go

First Name: prasanthi
Age: 23

// EXAMPLE 8 :

package main 
  
import "fmt"
type Employee struct { 
    firstName, lastName string 
    age, salary         int
} 
  
func main() { 
   emp8 := &Employee{"Sam", "Anderson", 55, 6000} 
    fmt.Println("First Name: ", emp8.firstName) 
    fmt.Println("Age: ", emp8.age) 
} 

$go run main.go

First Name:  Sam
Age:  55

// EXAMPLE 9 :

package main

import "fmt"

type User struct {
    name       string
    occupation string
    age        int
}

func main() {

    u := User{"prasanthi", "gardener", 23}

    fmt.Printf("%s is %d years old and she is a %s\n", u.name, u.age, u.occupation)
}

$go run main.go
prasanthi is 23 years old and she is a gardener

// EXAMPLE 10 :

package main

import "fmt"

type Address struct {
    city    string
    country string
}

type User struct {
    name    string
    age     int
    address Address
}

func main() {

    p := User{
        name: "prasanthi",
        age:  23,
        address: Address{
            city:    "New York",
            country: "USA",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("Country:", p.address.country)
}

$go run main.go

Name: prasanthi
Age: 23
City: New York
Country: USA































