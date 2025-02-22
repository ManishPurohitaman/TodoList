package main

import (
    "fmt"
    "net/http"
    "go-todo/routers"
)

func main() {
    r := routers.InitializeRouter() 

    fmt.Println("Todo List service running on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
