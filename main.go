package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, мир!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Сервер запущен на http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
