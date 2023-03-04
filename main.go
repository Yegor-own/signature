package main

import (
	"github.com/Yegor-own/signature/src/entity"
	"github.com/Yegor-own/signature/src/handler"
	"log"
	"net/http"
)

//var addr = flag.String("addr", ":8080", "http service address")

func main() {
	//flag.Parse()
	hub := entity.NewHub()
	go hub.Run()

	http.HandleFunc("/", handler.RootHandler)

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		entity.ServeWS(hub, writer, request)
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalln("Listen and serve:", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Fatalln("Listen and serve:", err)
		}
	}()

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalln("Listen and serve:", err)
	}
}
