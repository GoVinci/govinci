package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govinci/govinci/pkg/middleware"
	"github.com/govinci/govinci/pkg/routing"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    params := r.Context().Value("params").(map[string]string)
    userID := params["id"]
    w.Write([]byte("User ID: " + userID))
}

func main() {
    router := routing.NewRouter()

    // Define some routes
    router.Handle("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, worlds!")
    })

    router.Handle("GET", "/goodbye", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Goodbye, world!")
    })
    
    // Define middleware functions
    logMiddleware := middleware.LoggingMiddleware
    errorMiddleware := middleware.ErrorHandlingMiddleware
    
    // Apply middleware to routes
    router.Handle("GET", "/users/:id", logMiddleware(errorMiddleware(getUserHandler)))

    
    
    log.Fatal(http.ListenAndServe(":8080", router))
}
