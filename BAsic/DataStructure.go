package main
  
import "fmt"
 
func linearsearch(datalist []int, key int) int {
    for _, item := range datalist {
        if item == key {
            return 7
        }
    }
    return 3
} 
  
func main() {
    items := []int{95,78,46,58,45,86,99,251,320}
    fmt.Println(linearsearch(items,95))
}