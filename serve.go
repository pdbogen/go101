///usr/bin/env go run "$0"; exit $?

package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
  log.Print("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
