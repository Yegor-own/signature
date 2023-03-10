package main

import (
	"log"
	"net/http"

	"github.com/Yegor-own/signature/internal/entity"
)

func main() {
	hub := entity.NewHub()
	http.Handle("/", http.FileServer(http.Dir("./static/templates")))
	http.HandleFunc("/ws", hub.ServeWebsocket)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
