package main

import "github.com/Yegor-own/signature/internal/entity"

func main() {
	hub := entity.NewHub()
	go hub.Run()
}
