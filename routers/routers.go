package routers

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"go-todo/controllers"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,"health check passed!!!")
}

func InitializeRouter() *mux.Router {
    r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthcheckHandler).Methods("GET");
    r.HandleFunc("/todo-list-service/task/create", controllers.CreateTask).Methods("POST");
    r.HandleFunc("/todo-list-service/task/cancel", controllers.CancelTask).Methods("POST");
    r.HandleFunc("/todo-list-service/task/markCompleted", controllers.MarkTaskCompleted).Methods("POST");
    r.HandleFunc("/todo-list-service/task/getPending", controllers.GetPendingTasks).Methods("POST");
    
    return r
}
