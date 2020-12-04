package main
 
import (
    "fmt"
    "math/rand"
  //  "time"
)
 
func main() {
 
    d := generateSlice(10)
    fmt.Println("\n--- Unsorted --- \n\n", d)
   // quicksort(slice)
    //fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
 
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    d := make([]int, size, size)
    //rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        d[i] = rand.Intn(100) - rand.Intn(100)
    }
   return d
}