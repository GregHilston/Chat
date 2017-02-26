package main

import (
    "fmt"
    "net/http"
    "sync"
    "github.com/gorilla/websocket"
    "io/ioutil"
)

type msg struct {
	Message string
}

var mu sync.Mutex
var sockets []*websocket.Conn

func Index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	sockets = append(sockets, conn)

	go echo(conn)
}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("message: %#v\n", m)

		for index,element := range sockets {
			fmt.Printf("alerting: %#v\n", index)
			go alert(element, m)
		}

		//if err = conn.WriteJSON(m); err != nil {
		//	fmt.Println(err)
		//}
	}
}

func alert(conn *websocket.Conn, m msg) {
	mu.Lock()
	defer mu.Unlock()
	conn.WriteJSON(m)
}