package main

import (
    "net/http"
    "html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Index")
}