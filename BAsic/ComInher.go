/*package main
import "fmt"
type User struct {
Id int
Name string
Location string
}
type Player struct {
	Id int
	Name string
	Location string
	GameId int
	}
	func main() {
		s := User{}
		s.Id=44
		s.Name="fff"
		s.Location="ghdhd"
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
	fmt.Printf("%+v", s)
	}*/
	package main
import "fmt"
type User struct {
Id int
Name, Location string
}
type Player struct {
User
GameId int
}
func main() {
p := Player{
User{Id: 42, Name: "Matt", Location: "LA"},
90404,
}
fmt.Printf(
"Id: %d, Name: %s, Location: %s, Game id: %d\n",
p.Id, p.Name, p.Location, p.GameId)
// Directly set a field defined on the Player struct
//p.Id = 11
fmt.Printf("%+v", p)
}