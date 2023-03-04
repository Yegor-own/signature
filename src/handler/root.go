package handler

import (
	"log"
	"net/http"
)

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.URL)
	if request.URL.Path != "/" {
		http.Error(writer, "Not found", http.StatusNotFound)
		return
	}

	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(writer, request, "../template/root.html")
}
