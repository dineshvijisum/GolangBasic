package main

import (
    "io"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)
type CalcsHandler struct {}
func (handler *CalcsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "I will perform calculations for you...")
}
func main() {
    port := 8080 
    router := mux.NewRouter() 
    router.Handle("/calcs", &CalcsHandler{})  
    log.Printf("Starting server. Listening on port %d", port)
    err := http.ListenAndServe(":" + strconv.Itoa(port), router)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}