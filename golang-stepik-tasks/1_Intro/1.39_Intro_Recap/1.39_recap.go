package main

import "fmt"


func main() {
    var n uint

    fmt.Scan(&n)

    last := n % 10
    
    switch {
    case last == 0 || 11 <= n && n <= 14  || 5 <= last && last <= 9: fmt.Printf("%d korov", n)
    case last == 1: fmt.Printf("%d korova", n)
    case 2 <= last && last <= 4: fmt.Printf("%d korovy", n)
    }
}

