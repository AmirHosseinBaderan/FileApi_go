package server

import (
	"fmt"
	"net/http"
)

func RunServer() {
	err := http.ListenAndServe(":2500", nil)
	if err != nil {
		fmt.Println("Error listening", err)
	} else {
		fmt.Println("Server Run Success")
	}
}

func AddRoutes() {
	http.HandleFunc("/SaveImage", SaveImage)
}
