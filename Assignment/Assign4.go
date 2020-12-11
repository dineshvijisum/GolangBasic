package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Wep App", News: "Click the link"}
	t, _ := template.ParseFiles("Html.html")
	t.Execute(w, p)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>New Web</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":10000", nil)
}
