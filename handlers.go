package main

import (
    "fmt"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}

func Message(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Message!")
}