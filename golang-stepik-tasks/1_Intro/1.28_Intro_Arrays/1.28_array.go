package main

import "fmt"

func main() {
     var n, v, count int
     
     fmt.Scan(&n)

     for i := 0; i <n; i++ {
         fmt.Scan(&v)
         if v > 0{
             count ++
         }
     }
     fmt.Println(count)
}
