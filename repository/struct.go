package repository

type BankAccount struct {
	Income  int
	Outcome int
}

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
