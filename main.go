package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type successMessage struct {
	Message int `json:"success"`
}
type failMessage struct {
	Message int `json:"fail"`
}
type random2Response struct {
	Results []interface{} `json:"results"`
}

// 10回randして下記メッセージの配列を返す
// - 50以上なら {success: "70"}
// - 50未満なら {fail: "31"}
func random2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to random2")
	threshold := 50
	results := make([]interface{}, 10)
	for i := range results {
		v := rand.Intn(100)
		if v >= threshold {
			results[i] = successMessage{v}
		} else {
			results[i] = failMessage{v}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(random2Response{results})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type randomResponse struct {
	Value int `json:"value"`
}

func random(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to random")
	resp := randomResponse{rand.Intn(100)}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

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
	http.HandleFunc("/random", random)
	http.HandleFunc("/random2", random2)

	fmt.Printf("Start server at localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
