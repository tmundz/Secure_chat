package main

import (
	"fmt"
	"net/http"
)

func msg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello all")
}

func main() {
	http.HandleFunc("/", msg)
	port := 5051
	fmt.Printf("Server is listening on port %d....\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
