package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    i, err := strconv.ParseInt(r.URL.Path[1:], 10, 64)
    if( err == nil && ( i % 3 ) == 0){
        fmt.Fprintf(w, "Fizz")
    }
    if( err == nil && ( i % 5 ) == 0){
        fmt.Fprintf(w, "Buzz")
    }
    fmt.Fprintf(w, "\nHi there, I love %s!", r.URL.Path[1:])
}

func main(){
    http.HandleFunc("/",handler)
    http.ListenAndServe(":8080",nil)
}
