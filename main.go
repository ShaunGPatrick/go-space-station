package main

import (
	"go-space-station/server"
	"log"
	"net/http"
)

func main() {
	server := server.NewSpaceStationServer()
	log.Fatal(http.ListenAndServe(":8080", server))
}
