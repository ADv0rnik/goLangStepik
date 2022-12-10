package main

import "fmt"


func main() {
     var a, b, c int
     var workArray = [10]uint8{}
     for i := 0; i != 10; i++ {
	  fmt.Scan(&a)
	  workArray[i] = uint8(a)  
     }

     for i := 0; i < 3; i++ {
     	  fmt.Scan(&b, &c)
          workArray[b], workArray[c] = workArray[c], workArray[b]
     }
     fmt.Println(workArray)
}     
