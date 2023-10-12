package main

import (
	"log"
	"net/http"

	"github.com/zowber/zowber-portfolio-go/internal/routes"
)

func main() {

	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", router))

}
