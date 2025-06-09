package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/greetings", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == "POST" {
			type myStruct struct {
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
			}
			marshal, err := json.Marshal(myStruct{
				FirstName: request.FormValue("firstName"),
				LastName:  request.FormValue("lastName"),
			})

			if err != nil {
				writer.Write([]byte(err.Error()))
			}
			writer.Write(marshal)
		}

	})
	http.ListenAndServe(":8080", nil)
}
