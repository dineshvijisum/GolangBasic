package main
import (
	"fmt"
	"net/http"
)
//func dinesh(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w,"<h1>Hi dinesh</h1>")

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"<h1>Hi dinesh</h1>")})
	
	http.ListenAndServe(":2000",nil)
	
}