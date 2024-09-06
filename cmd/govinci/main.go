package main

import (
    "fmt"
    "net/http"
    "github.com/govinci/govinci/pkg/routing"
)

func main() {
    router := routing.NewRouter()

    // Define some routes
    router.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
    })

    router.Handle("GET", "/goodbye", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Goodbye, world!")
    })

    // Start the server
    http.ListenAndServe(":8080", router)
}
