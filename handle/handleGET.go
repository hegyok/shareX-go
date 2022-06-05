package handle

import (
	"net/http"
	"strings"
)

func HandleGET(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	if strings.Contains(path, "..") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if path == "/" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, "files"+path)
}
