package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myChannel :=make(chan string)
	go func (){
		myChannel <-"Hello w"
		//myChannel <-"World"
		}()
	text :=<-myChannel
	fmt.Println(text)
		fmt.Fprintf(w, text)
	})
	//http.HandleFunc("/",myHandlerfunc)
	http.ListenAndServe(":9000",nil)

}