package main

import "fmt"

func main() {
     var n, v int

     fmt.Scan(&n)

     for i := 0; i < n; i++ {
         fmt.Scan(&v)
         if i % 2 == 0{
             fmt.Printf("%d ",v)
 	 }
     }

}
