package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("The server is listening...")
	http.ListenAndServe(":8081", nil)
}

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}
