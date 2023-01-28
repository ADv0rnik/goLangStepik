package main

import (
	"fmt"
        "math"
       )


func main() {
    var n int
   
    fmt.Scan(&n)
    
    i := 0
    for {
	if n == 1{ 
	    fmt.Println(n)
	    break
	}    
	val := math.Pow(2, float64(i)) 
	if val <= float64(n) {	
            fmt.Print(val, " ")
	} else {
	  break
	}
	i++
    }
}

