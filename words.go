package main

import (
	"log"
	"os"
	"strings"
)

type Word struct {
	origin     string
	translated string
}

const WordsFileName = "resources/words"

func readWordsFromFile() []Word {
	fileData, fileError := os.ReadFile(WordsFileName)
	if fileError == nil {
		fileString := string(fileData)
		return parseWords(fileString)
	} else {
		log.Printf("fileError = %s", fileError.Error())
	}

	return []Word{}
}

func appendWords(words []Word) {
	wordsFromFile := readWordsFromFile()
	newWords := append(wordsFromFile, words...)
	os.WriteFile(WordsFileName, toBytes(newWords), 0644)
}

func toBytes(words []Word) []byte {
	var wordStrings []string
	for _, word := range words {
		wordStr := word.origin + ":" + word.translated
		wordStrings = append(wordStrings, wordStr)
	}
	return []byte(strings.Join(wordStrings, ","))
}

func parseWords(str string) []Word {
	var words []Word
	for _, wordString := range strings.Split(str, ",") {
		wordFields := strings.Split(wordString, ":")
		word := Word{
			origin:     wordFields[0],
			translated: wordFields[1],
		}
		words = append(words, word)
	}
	return words
}
