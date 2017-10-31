package main

import (
  "log"
  "net/http"
)

// go run *.go
// Ref: https://thenewstack.io/make-a-restful-json-api-go/

func main() {
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
