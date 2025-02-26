package main

import (
	"fmt"
	"log"
	"net/http"

	"go-todo/routers"
	"go-todo/middlewares"
)

func main() {
    log.Println("main Method called")
	r := routers.InitializeRouter()
	handler := middlewares.CORS(r)

	port := ":8080"
	fmt.Println("Todo List service running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
