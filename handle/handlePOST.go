package handle

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

var mimeTypes = map[string]string{
	"image/png": ".png",
	"video/mp4": ".mp4",
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func HandlePOST(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("authorization")
	if auth != os.Getenv("AUTHORIZATION") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ext := mimeTypes[r.Header.Get("content-type")]
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	random := randomString(3)
	ioutil.WriteFile("files/"+random+ext, data, 0644)
	var response = map[string]string{
		"filename": random + ext,
		ext:        ext,
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}
