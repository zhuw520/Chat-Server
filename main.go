package main

import (
	"log"
	"net/http"
	"chat-server/server"
)

func main() {
	chatServer := server.NewChatServer()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/ws", chatServer.HandleWebSocket)
	
	port := ":8092"
	log.Printf("server started :%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}