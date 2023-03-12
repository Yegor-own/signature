package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yegor-own/signature/internal/entity"
)

func main() {

	hub := entity.NewHub()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", hub.ServeWebsocket)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("hi")
}
