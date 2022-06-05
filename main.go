package main

import (
	"fmt"
	"net/http"
	"uploader/handle"
	"uploader/server"

	"github.com/joho/godotenv"
)

var (
	PORT = ":80"
)

func main() {
	godotenv.Load(".env")
	fmt.Println("Listening on port " + PORT)
	server.CreateServer(PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handle.HandlePOST(w, r)
		} else if r.Method == "GET" {
			handle.HandleGET(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(PORT, nil)
}
