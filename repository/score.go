package repository

import (
	"context"
	"fmt"

	"github.com/ficoback/domain"
)

type ScoreRepository struct{}

func NewRepositoryScore() ScoreRepository {
	return ScoreRepository{}
}

func (r ScoreRepository) GetScore(ctx context.Context, ID string) (domain.Score, error) {
	bankAccount, ok := mockScores[ID]
	if !ok {
		return domain.Score{}, fmt.Errorf("current ID doesn't have any score")
	}

	financialScore := calculFinancialScore(ctx, bankAccount)

	score := domain.Score{
		ID:    ID,
		Score: financialScore,
	}

	return score, nil
}

func calculFinancialScore(_ context.Context, bankAccount BankAccount) int {
	return bankAccount.Income - bankAccount.Outcome
}
