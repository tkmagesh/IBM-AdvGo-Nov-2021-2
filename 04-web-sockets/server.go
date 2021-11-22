package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var wsConnections []*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func streamTime(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Clinet connected")
	timer := time.Tick(5 * time.Second)
	for t := range timer {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(t.String())); err != nil {
			log.Println(err)
			return
		}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsConnections = append(wsConnections, ws)
	err = ws.WriteMessage(websocket.TextMessage, []byte("Welcome to the chat!"))
	if err != nil {
		log.Println(err)
		return
	}
	go reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		broadcast(string(message))
	}
}

func broadcast(message string) {
	for _, conn := range wsConnections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/time", streamTime)
	http.HandleFunc("/chat", chat)
	log.Println(http.ListenAndServe(":8080", nil))
}
