package server

import "net/http"

func CreateServer(port string) http.Server {
	Server := http.Server{
		Addr: "127.0.0.1" + port,
	}
	return Server
}
