
// EXAMPLE 1 :

package main  
import "fmt"  
func main ()  {  
   x := map[string]int{"PRASANTHI":23,"John":37, "Raj":20}  
   fmt.Print(x)  
   fmt.Println("\n",x["Raj"])  
}  

$go run main.go
map[John:37 PRASANTHI:23 Raj:20]
 20


// Go Map Insert and Update operation

// EXAMPLE 2 :

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   fmt.Println(m)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   m["Key2"] = 50 
   fmt.Println(m)  
}  

$go run main.go

map[]
map[Key1:10 Key2:20 Key3:30]
map[Key1:10 Key2:50 Key3:30]

//Go Map Delete operation

// EXAMPLE 3 :

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   delete(m, "Key3")  
   fmt.Println(m)  
}    

$go run main.go

map[Key1:10 Key2:20 Key3:30]
map[Key1:10 Key2:20]

// Go Map Retrieve Element

// EXAMPLE 4 :

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   fmt.Println("The value:", m["Key2"])  
}  

$go run main.go

map[Key1:10 Key2:20 Key3:30]
The value: 20

// EXAMPLE 5 :

package main  
import "fmt"  
func main() {  
   m := make(map[string]int)  
   m["Key1"] = 10  
   m["Key2"] = 20  
   m["Key3"] = 30  
   fmt.Println(m)  
   v, ok := m["Key2"]  
   fmt.Println("The value:", v, "Present?", ok)  
   i, j := m["Key4"]  
   fmt.Println("The value:", i, "Present?", j)  
}  

$go run main.go

map[Key1:10 Key2:20 Key3:30]
The value: 20 Present? true
The value: 0 Present? false

// Go Map of Struct

// EXAMPLE 6 :

package main  
import "fmt"  
type Vertex struct {  
   Latitude, Longitude float64  
}  
var m = map[string]Vertex{  
   "Jpoint": Vertex{     40.68433, -74.39967,   },  
   "SSS-IT": Vertex{     37.42202, -122.08408,  },  
}  
func main() {  
   fmt.Println(m)  
}  

$go run main.go

map[Jpoint:{40.68433 -74.39967} SSS-IT:{37.42202 -122.08408}]

// EXAMPLE 7 :

package main 
  
import "fmt"
  
func main() { 
  var map_1 map[int]int
   if map_1 == nil { 
      
        fmt.Println("True") 
    } else { 
      
        fmt.Println("False") 
    } 
   map_2 := map[int]string{ 
      
            90: "Dog", 
            91: "Cat", 
            92: "Cow", 
            93: "Bird", 
            94: "Rabbit", 
    } 
      
    fmt.Println("Map-2: ", map_2) 
} 

$go run main.go

True
Map-2:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]


// EXAMPLE 8 :

package main 
  
import "fmt"
  
func main() { 
 
    var My_map = make(map[float64]string) 
    fmt.Println(My_map) 
    My_map[1.3] = "PRASANTHI"
    My_map[1.5] = "HARITHA"
    fmt.Println(My_map) 
} 

$go run main.go
map[]
map[1.3:PRASANTHI 1.5:HARITHA]

//EXAMPLE 9 :

package main 
  
import "fmt"
  

func main() { 
  m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    fmt.Println("Original map: ", m_a_p) 
     m_a_p[95] = "Parrot"
    m_a_p[96] = "Crow"
    fmt.Println("Map after adding new key-value pair:\n", m_a_p) 
    m_a_p[91] = "PIG"
    m_a_p[93] = "DONKEY"
    fmt.Println("\nMap after updating values of the map:\n", m_a_p) 
} 

$go run main.go

Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
Map after adding new key-value pair:
 map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit 95:Parrot 96:Crow]

Map after updating values of the map:
 map[90:Dog 91:PIG 92:Cow 93:DONKEY 94:Rabbit 95:Parrot 96:Crow]

// EXAMPLE 10 :

package main 
  
import "fmt"
  
func main() { 
   m_a_p := map[int]string{ 
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
    fmt.Println("Original map: ", m_a_p) 
  new_map := m_a_p 
    new_map[96] = "Parrot"
    new_map[98] = "Pig"
    fmt.Println("New map: ", new_map) 
    fmt.Println("\nModification done in old map:\n", m_a_p) 
} 

$go run main.go

Original map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
New map:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit 96:Parrot 98:Pig]

Modification done in old map:
 map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit 96:Parrot 98:Pig]
























