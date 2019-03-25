package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var greetings = []string{
	"Hello",
	"Good morning",
	"Good afternoon",
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	randomNumber := randomGenerator.Intn(len(greetings))

	itemEnvValue := os.Getenv("ITEM")

	_,_ = fmt.Fprintf(w, "%s, I love %s!", greetings[randomNumber], itemEnvValue )
}
