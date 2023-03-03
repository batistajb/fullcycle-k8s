package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/secret", Secret)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

    name := os.Getenv("NAME")
    age := os.Getenv("AGE")

    fmt.Fprintf(w, "Hello, %s you are %s years old! V1", name, age)
}

func Secret(w http.ResponseWriter, r *http.Request) {

    pass := os.Getenv("PASS")
    user := os.Getenv("USER")

    fmt.Fprintf(w, "User: %s. Password: %s", user, pass)
}