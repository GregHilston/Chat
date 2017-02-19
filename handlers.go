package main

import (
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fs := http.FileServer(http.Dir("html"))
        http.Handle("/", fs)
}