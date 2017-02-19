package main

import (
    "fmt"
    //"io/ioutil"
    //"time"
    //"bytes"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    //r.ParseForm()
    //StoredAs := r.Form.Get("index.html") // file name
    // data, err := ioutil.ReadFile("static/" + StoredAs)
    // if err != nil { fmt.Fprint(w, err) }
    // http.ServeContent(w, r, StoredAs, time.Now(),   bytes.NewReader(data))

    // http.ServeFile(w, r, r.URL.Path[1:])
    fs := http.FileServer(http.Dir("static"))
        http.Handle("/index", fs)
}

func Message(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Message!")
}