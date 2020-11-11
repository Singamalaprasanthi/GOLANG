EXAMPLE 1 :

package main

import "fmt"

func main() {
   var n [10]int /* n is an array of 10 integers */
   var i,j int

   /* initialize elements of array n to 0 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* set element at location i to i + 100 */
   }
   
   /* output each array element's value */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}

OUTPUT :

Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109

EXAMPLE 2 :

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

OUTPUT :

Hello World
[Hello World]
[2 3 5 7 11 13]

EXAMPLE 3 :

package main 
  
import "fmt"
  
func main() { 
  
// Creating an array of string type  
// Using var keyword 
var myarr[3]string
  
// Elements are assigned using index 
myarr[0] = "SPR"
myarr[1] = "SINGAMALAPRASANTHIREDDY"
myarr[2] = "SINGAMALA"
  
// Accessing the elements of the array  
// Using index value 
fmt.Println("Elements of Array:") 
fmt.Println("Element 1: ", myarr[0]) 
fmt.Println("Element 2: ", myarr[1]) 
fmt.Println("Element 3: ", myarr[2]) 
} 

OUTPUT :

Elements of Array:
Element 1:  SPR
Element 2:  SINGAMALAPRASANTHIREDDY
Element 3:  SINGAMALA

EXAMPLE 4 :

package main 
  
import "fmt"
  
func main() { 
  
// Shorthand declaration of array 
arr:= [4]string{"SINGAMLA", "PRASANTHI", "REDDY" } 
  
// Accessing the elements of  
// the array Using for loop 
fmt.Println("Elements of the array:") 
  
for i:= 0; i < 3; i++{ 
fmt.Println(arr[i]) 
} 
  
} 

OUTPUT :

Elements of the array:
SINGAMLA
PRASANTHI
REDDY

EXAMPLE 5 :

package main 
  
import "fmt"
  
func main() { 
  
// Creating and initializing  
// 2-dimensional array  
// Using shorthand declaration 
// Here the (,) Comma is necessary 
arr:= [3][3]string{{"C#", "C", "Python"},  
                   {"Java", "Scala", "Perl"}, 
                    {"C++", "Go", "HTML"},} 
  
// Accessing the values of the  
// array Using for loop 
fmt.Println("Elements of Array 1") 
for x:= 0; x < 3; x++{ 
for y:= 0; y < 3; y++{ 
fmt.Println(arr[x][y]) 
} 
} 
  
// Creating a 2-dimensional 
// array using var keyword 
// and initializing a multi 
// -dimensional array using index 
var arr1 [2][2] int
arr1[0][0] = 100 
arr1[0][1] = 200 
arr1[1][0] = 300 
arr1[1][1] = 400 
  
// Accessing the values of the array 
fmt.Println("Elements of array 2") 
for p:= 0; p<2; p++{ 
for q:= 0; q<2; q++{ 
fmt.Println(arr1[p][q]) 
  
} 
} 
} 

OUTPUT :

Elements of Array 1
C#
C
Python
Java
Scala
Perl
C++
Go
HTML
Elements of array 2
100
200
300
400

EXAMPLE 6 :

package main 
  
import "fmt"
  
func main() { 
  
// Creating an array of int type 
// which stores, two elements 
// Here, we do not initialize the  
// array so the value of the array 
// is zero 
var myarr[2]int 
fmt.Println("Elements of the Array :", myarr) 
  
} 

OUTPUT :

Elements of the Array : [0 0]

EXAMPLE 7 :

package main 
  
import "fmt"
  
func main() { 
  
// Creating array 
// Using shorthand declaration     
arr1:= [3]int{9,7,6} 
arr2:= [...]int{9,7,6,4,5,3,2,4} 
arr3:= [3]int{9,3,5} 
  
// Finding the length of the  
// array using len method 
fmt.Println("Length of the array 1 is:", len(arr1)) 
fmt.Println("Length of the array 2 is:", len(arr2)) 
fmt.Println("Length of the array 3 is:", len(arr3)) 
} 

OUTPUT :

Length of the array 1 is: 3
Length of the array 2 is: 8
Length of the array 3 is: 3

EXAMPLE 8 :

package main 
  
import "fmt"
  
func main() { 
      
// Creating an array whose size is determined  
// by the number of elements present in it 
// Using ellipsis 
myarray:= [...]string{"PRA", "HRA", "MAHI", 
                    "GGK", "GOV"} 
  
fmt.Println("Elements of the array: ", myarray) 
  
// Length of the array 
// is determine by  
// Using len() method 
fmt.Println("Length of the array is:", len(myarray)) 
} 

OUTPUT :

Elements of the array:  [PRA HRA MAHI GGK GOV]
Length of the array is: 5

EXAMPLE 9 :

package main 
  
import "fmt"
  
func main() { 
      
// Creating an array whose size 
// is represented by the ellipsis 
myarray:= [...]int{29, 79, 49, 39, 
                   20, 49, 48, 49} 
  
// Iterate array using for loop 
for x:=0; x < len(myarray); x++{ 
fmt.Printf("%d\n", myarray[x]) 
} 
} 

OUTPUT :

29
79
49
39
20
49
48
49

EXAMPLE 10 :

package main 
  
import "fmt"
  
func main() { 
      
// Creating an array whose size  
// is represented by the ellipsis 
my_array:= [...]int{100, 200, 300, 400, 500} 
fmt.Println("Original array(Before):", my_array) 
  
// Creating a new variable  
// and initialize with my_array 
new_array := my_array 
  
fmt.Println("New array(before):", new_array) 
  
// Change the value at index 0 to 500 
new_array[0] = 500 
  
fmt.Println("New array(After):", new_array) 
  
fmt.Println("Original array(After):", my_array) 
} 

OUTPUT :

Original array(Before): [100 200 300 400 500]
New array(before): [100 200 300 400 500]
New array(After): [500 200 300 400 500]
Original array(After): [100 200 300 400 500]

EXAMPLE 11 :

package main 
  
import "fmt"
  
func main() { 
  
// Arrays     
arr1:= [3]int{9,7,6} 
arr2:= [...]int{9,7,6} 
arr3:= [3]int{9,5,3} 
  
// Comparing arrays using == operator 
fmt.Println(arr1==arr2) 
fmt.Println(arr2==arr3) 
fmt.Println(arr1==arr3) 
  
// This will give and error because the  
// type of arr1 and arr4 is a mismatch  
/* 
arr4:= [4]int{9,7,6} 
fmt.Println(arr1==arr4) 
*/
} 

OUTPUT :

true
false
false
























