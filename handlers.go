package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleGetWords(w http.ResponseWriter, req *http.Request) {
	log.Printf("method: %s", req.Method)
	if req.Method == http.MethodGet {
		for _, word := range readWordsFromFile() {
			fmt.Fprintf(w, "origin: %s\n", word.origin)
			fmt.Fprintf(w, "translated: %s\n\n", word.translated)
		}
	}
}

func handleAddWord(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err == nil {
			words := parseWords(string(body))
			appendWords(words)
		} else {
			log.Printf("handleAddWord err = %s", err.Error())
		}
	}
}
