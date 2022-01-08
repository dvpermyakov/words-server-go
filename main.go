package main

import (
	"log"
	"net/http"
)

func main() {
	addr := "localhost:8090"
	log.Printf("start server in addr: %s", addr)

	http.HandleFunc("/word/list", handleGetWords)
	http.HandleFunc("/word/add", handleAddWord)

	http.ListenAndServe(addr, nil)
}
