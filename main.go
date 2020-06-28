package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/renacargnelutti/golang-mongodb-api/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 50100...")
	log.Fatal(http.ListenAndServe(":50100", r))
}
