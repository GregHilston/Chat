package main

import (
    "net/http"
    "html/template"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Index")
}

func Messages(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    x := r.Form.Get("message")
    fmt.Println(x)
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	
}