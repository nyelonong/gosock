package main

import (
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func wsHandler(ws *websocket.Conn) {
	if err := websocket.Message.Send(ws, "Hey, you pushed me!"); err != nil {
		log.Println(err)
	}
}

func indexHanlder(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHanlder)
	http.Handle("/ws", websocket.Handler(wsHandler))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
