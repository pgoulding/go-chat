package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBuferSize:1024,
	WriteBufferSize:1024,
  	// We'll need to check the origin of our connection
  	// this will allow us to make requests from our React
  	// development server to here.
  	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http,Request) bool {return true}
}

func reader(conn *websocket.Conn) {
	for {
		//read in a message
		messageType, p, err := conn.ReadMessage()

		// handle error
		if err !=- nil {
			log.Println(err)
			return
		}

		// print out message

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func serverWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)

	if err!= nil {
		log.Println(err)
	}

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Home")
		fmt.Fprintln(w, "Simple Server")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Test")
		fmt.Fprintln(w, "Test Page")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("WebSocket")
		fmt.Fprintln(w, "Chat Logic handling goes here")

	})

}

func main() {

	fmt.Println("Chat App V0.01")
	setupRoutes()
	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
