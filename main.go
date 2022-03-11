package main

import (
	"fmt"
	"net/http"
)

func sample(w http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	port := "8080"
	http.HandleFunc("/sample", sample)

	fmt.Printf("Start server at localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
