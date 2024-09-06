package routing

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRouter(t *testing.T) {
    router := NewRouter()
    router.Handle("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Test successful"))
    })

    req := httptest.NewRequest("GET", "/test", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Fatalf("Expected status %v but got %v", http.StatusOK, res.StatusCode)
    }
    body := w.Body.String()
    if body != "Test successful" {
        t.Fatalf("Expected body 'Test successful' but got '%v'", body)
    }
}
