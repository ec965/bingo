package response

import (
	"net/http"
)

func Json(w http.ResponseWriter, j []byte, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
