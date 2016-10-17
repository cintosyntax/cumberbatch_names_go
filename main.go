package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "cumberbatch"
    "os"
)

func Generate(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, cumberbatch.GenerateName())
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Generate)

  port := os.Getenv("PORT")
  fmt.Println("Starting server on port", port)
  log.Fatal(http.ListenAndServe(":" + port, router))
}