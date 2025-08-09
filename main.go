package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"reinder/internal/httpserver"
	"reinder/internal/realtime"
)

func main() {
	hub := realtime.NewHub()
	go hub.Run()

	http.HandleFunc("/", httpserver.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		realtime.ServeWs(hub, w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("🚀 Starting Shared Docs server on port %s\n", port)
	fmt.Printf("📱 Open your browser and navigate to: http://localhost:%s\n", port)
	fmt.Println("💡 Multiple users can connect simultaneously for real-time collaboration")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
