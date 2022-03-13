package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func sample(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to sample")

	// http.ResponseWriterはio.Writer interfaceを満たしているので、
	// fmt.Fprintを利用できる
	fmt.Fprint(w, "test")

	// 他の記述方法として、http.ResponseWriterのWriteを使う場合は[]byteに変換する
	// w.Write([]byte("test"))
}

func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to echo")
	body := req.Body
	defer body.Close()

	// request bodyを読む際に、fileや標準入力を読むようにscannerを利用できる。
	// Bodyがio.Reader interfaceを満たすため。
	scanner := bufio.NewScanner(req.Body)
	for scanner.Scan() {
		fmt.Fprint(w, scanner.Text())
	}
}

// jsonのレスポンスに利用する構造体
type randomResponse struct {
	// 大文字の公開されている変数のみjsonの変換対象になる
	Value int `json:"value"`
}

func random(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to random")
	resp := randomResponse{rand.Intn(100)}

	// レスポンスヘッダーの設定はbodyに書き込む前に実施する必要がある
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoderは指定したio.Writerにjson encode結果を書き込むencoderを作成する。
	// http.ResponseWriterはio.Writerを満たすため、json.NewEncoderの書き込み先に指定できる。
	// json encodeした結果をresponse bodyに書き込む。
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type random2Response struct {
	Results []map[string]int `json:"results"`
}

// 10回randして下記メッセージの配列を返す
// - 50以上なら {success: "70"}
// - 50未満なら {fail: "31"}
func random2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to random2")
	threshold := 50
	results := make([]map[string]int, 10)
	for i := range results {
		v := rand.Intn(100)
		if v >= threshold {
			results[i] = map[string]int{"success": v}
		} else {
			results[i] = map[string]int{"fail": v}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(random2Response{results})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type echoBody struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

// request bodyでechoBody型のjsonを受け取り、そのままjsonをレスポンス
func echoJson(w http.ResponseWriter, req *http.Request) {
	fmt.Println("request to echo json")

	b := echoBody{}

	// json.NewDecoderは指定したio.Readerからjson文字列を読むDecoderを作成する。
	// Bodyはio.Readerを満たすため、json.NewDecoderの読み込み先に指定できる。
	// json decodeした結果をechoBodyインスタンスへ書き込む。
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		// decodeのエラーメッセージをそのままレスポンス
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	port := "8080"
	http.HandleFunc("/sample", sample)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/random", random)
	http.HandleFunc("/random2", random2)
	http.HandleFunc("/echojson", echoJson)

	fmt.Printf("Start server at localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
