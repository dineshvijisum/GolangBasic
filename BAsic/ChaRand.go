package main 
import ( 
    "fmt"
   "math/rand"
  // "text/tabwriter"
  // "os"
   // "time"
) 
func main() { 
	NewSource :=make(chan int)
	go func(){NewSource <- 100}()
    d :=<-NewSource
	fmt.Println(d)
	r := rand.New(rand.NewSource(99))
/*	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t\n", name, v1)
	}*/
	fmt.Println("int", r.Int())
}