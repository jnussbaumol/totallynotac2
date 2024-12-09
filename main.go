package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/health", getBotNet)
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

type botBody struct {
	Test string `json:"test"`
}

func getBotNet(w http.ResponseWriter, req *http.Request) {
	log.Println("Bot Phoning Home!")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if req.Method == "GET" {
		log.Println("Bot Checking Home!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "C2 is Alive")
		return
	} else if req.Method == "POST" {
		log.Println("Bot Sending Data!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Received")
	} else {
		log.Println("Wrong Method Hit! DANGER WE'VE BEEN MADE")
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Wrong Method")
		return
	}
	var bb botBody
	err := json.NewDecoder(req.Body).Decode(&bb)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", bb)
}
