package main

import (
	"log"
	"net/http"

	"zowber-portfolio-go/internal/conf"
	"zowber-portfolio-go/internal/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":"+conf.DotEnv("HTTP_PORT"), router))
}
