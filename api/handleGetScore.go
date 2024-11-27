package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	model "github.com/ficoback/api/model/handleGetScore"
	"github.com/ficoback/domain"
	"github.com/ficoback/repository"
)

func HandleGetScore(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	repo := repository.NewRepositoryScore()
	service := domain.NewServiceScore(repo)

	score, err := service.GetScore(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.Response{
		ID:    score.ID,
		Score: score.Score,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
}
