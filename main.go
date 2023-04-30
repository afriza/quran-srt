package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
