package main


import "fmt"


func main() {
    var fib1, fib2 int = 1, 1
    var val int

    fmt.Scan(&val)
    
    for i:=1; i <= val; i++ {
        fib1, fib2 = fib2, fib1 + fib2
	
	if fib2 == val {
	    fmt.Println(i+2)
	    return
	}
    }
    fmt.Println(-1)
}
	    
