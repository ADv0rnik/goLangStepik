package main


import "fmt"


func main() {
    var a, b, max int

    fmt.Scan(&a, &b)

    for ; b >= a; b-- {
        if b % 7 == 0 {
           max = b
           break
        }     
    }
    if max >= a && max <= b {
        fmt.Println(max)
    } else {
        fmt.Println("No")
    }    
}
