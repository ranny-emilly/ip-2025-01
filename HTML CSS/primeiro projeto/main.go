package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
