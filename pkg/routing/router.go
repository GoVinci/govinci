package routing

import (
    "net/http"
    "strings"
)

// Route defines a single route in the router with optional parameters
type Route struct {
    Method  string
    Path    string
    Handler http.HandlerFunc
}

// Router holds the routes and manages routing requests
type Router struct {
    routes []Route
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
    return &Router{}
}

// Handle registers a new route with the router
func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
    r.routes = append(r.routes, Route{Method: method, Path: path, Handler: handler})
}

// ServeHTTP is the main routing function that directs requests to the appropriate handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if route.Method == req.Method && strings.HasPrefix(req.URL.Path, route.Path) {
            route.Handler.ServeHTTP(w, req)
            return
        }
    }
    http.NotFound(w, req)
}
