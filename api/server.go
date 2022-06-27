package api

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Start() {
	router := Controllers()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	handlerServer := cors.AllowAll().Handler(router)
	println("Server run at port: " + PORT)
	log.Fatalln(http.ListenAndServe(":"+PORT, handlerServer))
}
