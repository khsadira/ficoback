package domain

import (
	"context"
	"fmt"
)

type Score struct {
	ID    string
	Score int
}

type ServiceScore struct {
	scoreRepository ScoreRepository
}

type ScoreRepository interface {
	GetScore(ctx context.Context, ID string) (Score, error)
}

func NewServiceScore(scoreRepository ScoreRepository) ServiceScore {
	return ServiceScore{
		scoreRepository: scoreRepository,
	}
}

func (s ServiceScore) GetScore(ctx context.Context, ID string) (Score, error) {
	score, err := s.scoreRepository.GetScore(ctx, ID)
	if err != nil {
		return Score{}, fmt.Errorf("failed to get score: %w", err)
	}

	return score, nil
}
