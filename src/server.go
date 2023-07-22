package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server listening on port 3001")
	http.ListenAndServe(":3001", http.FileServer(http.Dir("public/")))
}
