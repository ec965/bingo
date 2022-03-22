package response

import (
	"net/http"
)

func JsonDecodeError(w http.ResponseWriter, errStr string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte(errStr))
}

func ValidationError(w http.ResponseWriter, msg []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func Standard(w http.ResponseWriter, j []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
