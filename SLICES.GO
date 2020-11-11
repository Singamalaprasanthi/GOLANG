EXAMPLE 1 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating an array 
    arr := [7]string{"This", "is", "the", "tutorial", 
                         "of", "Go", "language"} 
  
    // Display array 
    fmt.Println("Array:", arr) 
  
    // Creating a slice 
    myslice := arr[1:6] 
  
    // Display slice 
    fmt.Println("Slice:", myslice) 
  
    // Display length of the slice 
    fmt.Printf("Length of the slice: %d", len(myslice)) 
  
    // Display the capacity of the slice 
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice)) 
} 

OUTPUT :

Array: [This is the tutorial of Go language]
Slice: [is the tutorial of Go]
Length of the slice: 5
Capacity of the slice: 6

EXAMPLE 2 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a slice 
    // using the var keyword 
    var my_slice_1 = []string{"PRASANTHI", "REDDY", "SINGAMALA"} 
  
    fmt.Println("My Slice 1:", my_slice_1) 
  
    // Creating a slice 
    //using shorthand declaration 
    my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45} 
    fmt.Println("My Slice 2:", my_slice_2) 
} 

OUTPUT :

My Slice 1: [Geeks for Geeks]
My Slice 2: [12 45 67 56 43 34 45]

EXAMPLE 3 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating an array 
    arr := [4]string{"PRASANTHI", "REDDY", "SINGAMALA", "SPR"} 
  
    // Creating slices from the given array 
    var my_slice_1 = arr[1:2] 
    my_slice_2 := arr[0:] 
    my_slice_3 := arr[:2] 
    my_slice_4 := arr[:] 
  
    // Display the result 
    fmt.Println("My Array: ", arr) 
    fmt.Println("My Slice 1: ", my_slice_1) 
    fmt.Println("My Slice 2: ", my_slice_2) 
    fmt.Println("My Slice 3: ", my_slice_3) 
    fmt.Println("My Slice 4: ", my_slice_4) 
} 

OUTPUT :

My Array:  [PRASANTHI REDDY SINGAMALA SPR]
My Slice 1:  [REDDY]
My Slice 2:  [PRASANTHI REDDY SINGAMALA SPR]
My Slice 3:  [PRASANTHI REDDY]
My Slice 4:  [PRASANTHI REDDY SINGAMALA SPR]

EXAMPLE 4 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating s slice 
    oRignAl_slice := []int{10, 20, 30, 40, 50, 60, 70} 
  
    // Creating slices from the given slice 
    var my_slice_1 = oRignAl_slice[1:5] 
    my_slice_2 := oRignAl_slice[0:] 
    my_slice_3 := oRignAl_slice[:6] 
    my_slice_4 := oRignAl_slice[:] 
    my_slice_5 := my_slice_3[2:4] 
  
    // Display the result 
    fmt.Println("Original Slice:", oRignAl_slice) 
    fmt.Println("New Slice 1:", my_slice_1) 
    fmt.Println("New Slice 2:", my_slice_2) 
    fmt.Println("New Slice 3:", my_slice_3) 
    fmt.Println("New Slice 4:", my_slice_4) 
    fmt.Println("New Slice 5:", my_slice_5) 
} 

OUTPUT :

Original Slice: [10 20 30 40 50 60 70]
New Slice 1: [20 30 40 50]
New Slice 2: [10 20 30 40 50 60 70]
New Slice 3: [10 20 30 40 50 60]
New Slice 4: [10 20 30 40 50 60 70]
New Slice 5: [30 40]

EXAMPLE 5 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating an array of size 7 
    // and slice this array  till 4 
    // and return the reference of the slice 
    // Using make function 
    var my_slice_1 = make([]int, 4, 7) 
    fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n", 
                   my_slice_1, len(my_slice_1), cap(my_slice_1)) 
  
    // Creating another array of size 7 
    // and return the reference of the slice 
    // Using make function 
    var my_slice_2 = make([]int, 7) 
    fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n", 
                   my_slice_2, len(my_slice_2), cap(my_slice_2)) 
      
} 

OUTPUT :

Slice 1 = [0 0 0 0], 
length = 4, 
capacity = 7
Slice 2 = [0 0 0 0 0 0 0], 
length = 7, 
capacity = 7

EXAMPLE 6 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a slice 
    myslice := []string{"This", "is", "the", "tutorial", 
        "of", "Go", "language"} 
  
    // Iterate using for loop 
    for e := 0; e < len(myslice); e++ { 
        fmt.Println(myslice[e]) 
    } 
}

OUTPUT :

This
is
the
tutorial
of
Go
language

EXAMPLE 7 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a slice 
    myslice := []string{"This", "is", "the", "tutorial", 
                                 "of", "Go", "language"} 
  
    // Iterate slice 
    // using range in for loop 
    for index, ele := range myslice { 
        fmt.Printf("Index = %d and element = %s\n", index+3, ele) 
    } 
} 

OUTPUT :

Index = 3 and element = This
Index = 4 and element = is
Index = 5 and element = the
Index = 6 and element = tutorial
Index = 7 and element = of
Index = 8 and element = Go
Index = 9 and element = language

EXAMPLE 8 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a slice 
    myslice := []string{"This", "is", "the", 
        "tutorial", "of", "Go", "language"} 
  
    // Iterate slice 
    // using range in for loop 
    // without index 
    for _, ele := range myslice { 
        fmt.Printf("Element = %s\n", ele) 
    } 
} 

OUTPUT :

Element = This
Element = is
Element = the
Element = tutorial
Element = of
Element = Go
Element = language

EXAMPLE 9 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a zero value slice 
    var myslice []string 
    fmt.Printf("Length = %d\n", len(myslice)) 
    fmt.Printf("Capacity = %d ", cap(myslice)) 
  
} 

OUTPUT :

Length = 0
Capacity = 0 

EXAMPLE 10 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating a zero value slice 
    arr := [6]int{55, 66, 77, 88, 99, 22} 
    slc := arr[0:4] 
  
    // Before modifying 
  
    fmt.Println("Original_Array: ", arr) 
    fmt.Println("Original_Slice: ", slc) 
  
    // After modification 
    slc[0] = 100 
    slc[1] = 1000 
    slc[2] = 1000 
  
    fmt.Println("\nNew_Array: ", arr) 
    fmt.Println("New_Slice: ", slc) 
} 

OUTPUT :

Original_Array:  [55 66 77 88 99 22]
Original_Slice:  [55 66 77 88]

New_Array:  [100 1000 1000 88 99 22]
New_Slice:  [100 1000 1000 88]

EXAMPLE 11 :

package main 
  
import "fmt"
  
func main() { 
  
    // creating slices 
    s1 := []int{12, 34, 56} 
    var s2 []int
  
    // If you try to run this commented 
    // code compiler will give an error 
    /*s3:= []int{23, 45, 66} 
      fmt.Println(s1==s3) 
    */
  
    // Checking if the given slice is nil or not 
    fmt.Println(s1 == nil) 
    fmt.Println(s2 == nil) 
} 

OUTPUT :

false
true

EXAMPLE 12 :

package main 
  
import "fmt"
  
func main() { 
  
    // Creating multi-dimensional slice 
    s1 := [][]int{{12, 34}, 
        {56, 47}, 
        {29, 40}, 
        {46, 78}, 
    } 
  
    // Accessing multi-dimensional slice 
    fmt.Println("Slice 1 : ", s1) 
  
    // Creating multi-dimensional slice 
    s2 := [][]string{ 
        []string{"PRASANTHI", "PRA"}, 
        []string{"REDDY", "RD"}, 
        []string{"MAGI", "MAG"}, 
    } 
  
    // Accessing multi-dimensional slice 
    fmt.Println("Slice 2 : ", s2) 
  
} 

OUTPUT :

Slice 1 :  [[12 34] [56 47] [29 40] [46 78]]
Slice 2 :  [[PRASANTHI PRA] [REDDY RD] [MAGI MAG]]

EXAMPLE 13 :

package main 
  
import ( 
    "fmt"
    "sort"
) 
  
func main() { 
  
    // Creating Slice 
    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"} 
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89} 
  
    fmt.Println("Before sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
  
    // Performing sort operation on the 
    // slice using sort function 
    sort.Strings(slc1) 
    sort.Ints(slc2) 
  
    fmt.Println("\nAfter sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
  
} 

OUTPUT :

Before sorting:
Slice 1:  [Python Java C# Go Ruby]
Slice 2:  [45 67 23 90 33 21 56 78 89]

After sorting:
Slice 1:  [C# Go Java Python Ruby]
Slice 2:  [21 23 33 45 56 67 78 89 90]
















