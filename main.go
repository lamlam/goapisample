package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to echo")
	body := req.Body
	defer body.Close()

	scanner := bufio.NewScanner(req.Body)
	for scanner.Scan() {
		fmt.Fprint(w, scanner.Text())
	}
}

func sample(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to sample")
	fmt.Fprintf(w, "test")
}

func main() {
	port := "8080"
	http.HandleFunc("/sample", sample)
	http.HandleFunc("/echo", echo)

	fmt.Printf("Start server at localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
