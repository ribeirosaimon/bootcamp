package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		if err := json.NewEncoder(writer).Encode("pong"); err != nil {
			json.NewEncoder(writer).Encode("error")
		}
	})
	http.ListenAndServe(":8080", nil)
}
