package main

import (
    "log"
    "net/http"
    "central-control/internal/api"
)

func main() {
    mux := http.NewServeMux()
    api.RegisterRoutes(mux)
    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", mux)
}
