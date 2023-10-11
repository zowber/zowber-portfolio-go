package main

import (
	"log"
	"net/http"

	"github.com/zowber/zowber-portfolio-go/internal/routes"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServeTLS(":8080", "server.pem", "server.key", router))

}
