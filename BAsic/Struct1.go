/*package main
import "fmt"
type Point struct {
X, Y int
}
var (
p = Point{1, 2} // has type Point
q = &Point{1, 2} // has type *Point
r = Point{X: 1} // Y:0 is implicit
s = Point{} // X:0 and Y:0
)
func main() {
fmt.Println(p, q, r, s)
}*/
package main
import (
"fmt"
"time"
)
type Bootcamp struct {
Lat, Lon float64
Date time.Time
}
func main() {
	event := Bootcamp{
	Lat: 34.012836,
	Lon: -118.495338,
	//event.Date = time.Now(),
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)",
	event.Date, event.Lat, event.Lon)
	}