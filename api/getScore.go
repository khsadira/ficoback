package api

import (
	"net/http"
)

func HandleGetScore(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
}
