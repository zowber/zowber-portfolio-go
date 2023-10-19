package main

import (
	"log"
	"net/http"

	"github.com/zowber/zowber-portfolio-go/internal/conf"
	"github.com/zowber/zowber-portfolio-go/internal/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":"+conf.DotEnv("HTTP_PORT"), router))
}
