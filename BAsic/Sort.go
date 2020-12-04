
	package main
 
import (
    "fmt"
    "math/rand"
   // "time"
)
 
func main() {
 
    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    quicksort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
func generateSlice(size int) []int {
 
    d := make([]int, size, size)
    //rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        d[i] = rand.Intn(100) - rand.Intn(100)
    }
   return d
}
   func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
	left, right := 0, len(a)-1
      
    pivot := rand.Int() % len(a)
      
    a[pivot], a[right] = a[right], a[pivot]
      
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
      
    a[left], a[right] = a[right], a[left]
      
    quicksort(a[:left])
    quicksort(a[left+1:])
      
    return a
}