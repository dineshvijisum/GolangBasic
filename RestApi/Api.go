package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	
)
type Art struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Art1 []Art

func allArt(w http.ResponseWriter, r *http.Request){
	art1 :=Art1{
		Art{Title:"Test Title",Desc:"Test Description",Content:"Hello World"},
	}
	fmt.Println("hhhh:hhhh")
	json.NewEncoder(w).Encode(art1)
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi dinesh")
}
func handleRequest(){
	myRouter :=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/art1",allArt)
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}
func main() {
	handleRequest()
}