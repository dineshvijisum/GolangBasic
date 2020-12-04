package main
import( 
	"fmt"
	"encoding/json"
)
type Person struct{
	Name string
	Age int
	Car []string
}
func main(){
	
	p1 :=&Person{Name:"Dinesh",Age:25,Car:[]string{"Hundai","maruthi"}}
	data,_ :=json.Marshal(p1)
	fmt.Println(string(data))
}