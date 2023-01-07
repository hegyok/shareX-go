package server

import "net/http"

func CreateServer(port string) http.Server {
	Server := http.Server{
		Addr: "0.0.0.0" + port,
	}
	return Server
}
