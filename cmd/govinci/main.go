package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/govinci/govinci/pkg/middleware"
	"github.com/govinci/govinci/pkg/routing"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Safely retrieve the parameters from the context
    params, ok := r.Context().Value(routing.ParamsKey).(map[string]string)
    if !ok {
        http.Error(w, "Invalid parameters", http.StatusBadRequest)
        return
    }
    
    // Check if the user ID exists in the parameters
    userID, exists := params["id"]
    if !exists {
        http.Error(w, "User ID not found", http.StatusBadRequest)
        return
    }
    
    // Respond with the user ID
    w.Write([]byte("My User ID: " + userID))
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