package repository

import (
	"context"

	"github.com/ficoback/domain"
)

type ScoreRepository struct{}

func NewRepositoryScore() ScoreRepository {
	return ScoreRepository{}
}

func (r ScoreRepository) GetScore(_ context.Context, ID string) (domain.Score, error) {
	score := domain.Score{
		ID:    ID,
		Score: 100,
	}
	return score, nil
}
