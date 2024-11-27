package repository

import (
	"context"
	"fmt"

	"github.com/ficoback/domain"
)

var mockScores = map[string]BankAccount{
	"1": {
		Income:  1000,
		Outcome: 500,
	},
	"2": {
		Income:  2000,
		Outcome: 1000,
	},
	"toto": {
		Income:  3000,
		Outcome: 5000,
	},
	"picsou": {
		Income:  10000,
		Outcome: 0,
	},
}

type ScoreRepository struct{}

func NewRepositoryScore() ScoreRepository {
	return ScoreRepository{}
}

func (r ScoreRepository) GetScore(_ context.Context, ID string) (domain.Score, error) {
	bankAccount, ok := mockScores[ID]
	if !ok {
		return domain.Score{}, fmt.Errorf("current ID doesn't have any score")
	}

	financialScore := calculFinancialScore(bankAccount)

	score := domain.Score{
		ID:    ID,
		Score: financialScore,
	}

	return score, nil
}

func calculFinancialScore(bankAccount BankAccount) int {
	return bankAccount.Income - bankAccount.Outcome
}
