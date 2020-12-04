package main
import "fmt"
type Vertex struct {
Lat, Long float64
bb string
}
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967,"dij"},
// same as "Bell Labs": Vertex{40.68433, -74.39967}
"Google": {37.42202, -122.08408,"kk"},
//"aaa":{"dinesh"},
}
func main() {
fmt.Println(m)
}
