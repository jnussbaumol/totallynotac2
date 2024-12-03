package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	serv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/health", checkHealth)
	mux.HandleFunc("/c2", getBotNet)
	fmt.Println("Starting server...")
	log.Fatal(serv.ListenAndServe())
}

type botBody struct {
	Test string `json:"test"`
}

func checkHealth(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Health Queried!")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	int, err := io.WriteString(w, "OK")
	if err != nil {
		log.Fatal(err, int)
	}
}

func getBotNet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Bot Phoning Home!")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	var err error
	var int int
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		int, err = io.WriteString(w, "Use POST!")
	} else {
		w.WriteHeader(http.StatusOK)
		int, err = io.WriteString(w, "OK")
	}
	if err != nil {
		log.Fatal(err, int)
	}
	var bb botBody
	err = json.NewDecoder(req.Body).Decode(&bb)
	if err != nil {
		log.Fatal(err, int)
	}
	fmt.Printf("%+v\n", bb)
}
