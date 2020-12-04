package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*myChannel :=make(chan string)
	go func (){
		myChannel <-"Hello world"
		//myChannel <-"World"
		}()
	text :=<-myChannel
	//fmt.Println(text)*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "text")
	})
	//http.HandleFunc("/",myHandlerfunc)
	http.ListenAndServe(":9000",nil)

}