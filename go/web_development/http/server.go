package main

import (
    "fmt"
    "net/http"
)

func greeting(w http.ResponseWriter, r *http.Request){
    
    fmt.Fprintf(w, "HTTP World")

}

func main(){

    http.HandleFunc("/", greeting)
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.ListenAndServe(":8080", nil)
}
