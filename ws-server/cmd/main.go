package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		log.Print("upgrade sucessful! ")

		for {

			bytes := []byte("hello: " + time.Now().String())
			err = conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
			time.Sleep(time.Second)
			fmt.Println("msg sent!")
		}
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
