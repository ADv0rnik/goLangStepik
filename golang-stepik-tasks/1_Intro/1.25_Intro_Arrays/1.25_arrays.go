/* Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/

package main


import "fmt"


func main() {
     var n int
     
     fmt.Scan(&n)     
     a := make([]int, n) 
    
     for i := 0; i < n; i++ {
	 fmt.Scan(&a[i])	     
     }

     fmt.Println(a[3])
}



