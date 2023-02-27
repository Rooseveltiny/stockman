package stockmanrestapiserver

import "net/http"

func ThrowError(w http.ResponseWriter, status int, err error) {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
	}
}

func Respond(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	w.Write(body)
}
