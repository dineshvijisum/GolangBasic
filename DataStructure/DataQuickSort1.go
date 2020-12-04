package main
import (
	"fmt"
)
func partition(a []int, b, c int) int {
	p := a[c]
	for j := b; j < c; j++ {
		if a[j] < p {
			a[j], a[b] = a[b], a[j]
			b++
		}
	}
	a[b], a[c] = a[c], a[b]
	return b
}
func quickSort(a []int, b, c int) {
	if b > c {
		return
	}
	p := partition(a, b, c)
	quickSort(a, b, p-1)
	quickSort(a, p+1, c)
}
func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14}	
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}