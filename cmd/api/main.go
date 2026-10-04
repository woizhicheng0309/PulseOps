package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("PulseOps API listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
