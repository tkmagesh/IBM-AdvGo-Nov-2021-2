package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

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

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/time", streamTime)
	log.Println(http.ListenAndServe(":8080", nil))
}
