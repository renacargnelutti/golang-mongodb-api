package main

import (
	"fmt"
	"log"
	"monitoringapi/handler"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	sub := router.PathPrefix("/api/v1").Subrouter()

	// endpoints
	sub.Methods("GET").Path("/status").HandlerFunc(handler.GetOSData)

	fmt.Println("Monitoring API listening on port: 50100")
	log.Fatal(http.ListenAndServe(":50100", router))
}

// log request
func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
