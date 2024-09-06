package routing

import (
	"context"
	"net/http"
	"strings"
)

type Route struct {
    Method  string
    Path    string
    Handler http.HandlerFunc
    Params  map[string]int // Parameter names and their positions
}

type contextKey string

type Router struct {
    routes []Route
}

func NewRouter() *Router {
    return &Router{}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
    params := make(map[string]int)
    pathSegments := strings.Split(path, "/")
    
    // Collect parameters
    for i, segment := range pathSegments {
        if strings.HasPrefix(segment, ":") {
            paramName := segment[1:]
            params[paramName] = i
        }
    }
    
    r.routes = append(r.routes, Route{
        Method:  method,
        Path:    path,
        Handler: handler,
        Params:  params,
    })
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if req.Method == route.Method {
            routeSegments := strings.Split(route.Path, "/")
            requestSegments := strings.Split(req.URL.Path, "/")
            
            if len(routeSegments) == len(requestSegments) {
                params := make(map[string]string)
                match := true
                for paramName, index := range route.Params {
                    if index >= len(requestSegments) || routeSegments[index] != requestSegments[index] {
                        match = false
                        break
                    }
                    params[paramName] = requestSegments[index]
                }
                 if match {
                    req = req.WithContext(context.WithValue(req.Context(), contextKey("params"), params)) // Use the custom type
                    route.Handler(w, req)
                    return
                }
            }
        }
    }
    http.NotFound(w, req)
}
