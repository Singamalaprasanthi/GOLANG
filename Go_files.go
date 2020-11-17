// 1. create file

package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

$go run main.go
11 bytes written successfully

// 2. example

package main
 
import (
	"log"
	"os"
)
 
func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

$go run main.go
2020/11/17 06:54:29 &{0xc42009c060}

// 3. example

package main
import (
	"log"
	"os"
)
 
func main() {
	_, err := os.Stat("test")
 
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
 
	}
}

$go run main.go

// 4. example

package main
 
import (
	"log"
	"os"
)
 
func main() {
	oldName := "test.txt"
	newName := "testing.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

$go run main.go
2020/11/17 06:55:09 rename test.txt testing.txt: no such file or directory
exit status 1


// 5.example

package main
 
import (
	"io"
	"log"
	"os"
)
 
func main() {
 
	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()
 
	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
 
	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

 $go run main.go
2020/11/17 06:55:58 open /src/test.txt: no such file or directory
exit status 1

// 6.example

package main
 
import (
	"fmt"
	"log"
	"os"
)
 
func main() {
	fileStat, err := os.Stat("test.txt")
 
	if err != nil {
		log.Fatal(err)
	}
 
	fmt.Println("File Name:", fileStat.Name())        
	fmt.Println("Size:", fileStat.Size())             
	fmt.Println("Permissions:", fileStat.Mode())      
	fmt.Println("Last Modified:", fileStat.ModTime()) 
	fmt.Println("Is Directory: ", fileStat.IsDir())   
}

$go run main.go
2020/11/17 06:57:46 stat test.txt: no such file or directory
exit status 1

// 7.example 

package main
 
import (
	"log"
	"os"
)
 
func main() {
	err := os.Remove("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

$go run main.go
2020/11/17 06:58:39 remove /test.txt: no such file or directory
exit status 1

// 8.example

package main
 
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
 
func main() {
	filename := "test.txt"
 
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
 
	for data.Scan() {
		fmt.Print(data.Text())
	}
}
$go run main.go
open test.txt: no such file or directory
exit status 1

// 9.example

package main
 
import (
	"log"
	"os"
)
 
func main() {
	err := os.Truncate("test.txt", 100)
 
	if err != nil {
		log.Fatal(err)
	}
}

$go run main.go
2020/11/17 06:59:53 truncate test.txt: no such file or directory
exit status 1

// 10.example

package main
 
import (
	"log"
	"os"
	"time"
)
 
func main() {
	// Test File existence.
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")
 
	
	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}
 
	
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}
 
	
	addOneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

$go run main.go
2020/11/17 07:01:08 File does not exist.
exit status 1



