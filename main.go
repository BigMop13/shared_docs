package main

import (
	"fmt"
	"net/http"
)

var (
	channelId = 123
)

func main() {
	http.HandleFunc("/", sendRequest)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")

	if err != nil {
		panic(err)
	}
}

func sendRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("test")

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)
}
