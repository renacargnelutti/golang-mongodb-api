package router

import (
	"net/http"
	"os"

	"github.com/renacargnelutti/golang-mongodb-api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	sub := router.PathPrefix("/api/v1").Subrouter()

	// endpoints
	sub.Methods("GET").Path("/tasks").HandlerFunc(routes.GetAllTask)
	sub.Methods("POST").Path("/tasks").HandlerFunc(routes.CreateTask)

	// router.HandleFunc("/api/task/{id}", handler.TaskComplete).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/undoTask/{id}", handler.UndoTask).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/deleteTask/{id}", handler.DeleteTask).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/deleteAllTask", handler.DeleteAllTask).Methods("DELETE", "OPTIONS")

	return router
}

// log request
func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
