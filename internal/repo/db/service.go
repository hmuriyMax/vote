package db

import (
	"github.com/hmuriyMax/vote/internal/repo/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type InMemory struct {
	votes    model.Votes
	variants map[int64]model.Variants
}

func NewPostgres(db *sqlx.DB) *InMemory {
	return &InMemory{
		votes: []*model.Vote{
			{
				ID:         1,
				Name:       "test vote",
				FinishesAt: time.Date(2033, 04, 30, 23, 31, 24, 12, time.Local),
			},
		},
		variants: map[int64]model.Variants{
			1: {
				{
					ID:            1,
					VoteID:        1,
					Name:          "test variant 1",
					CurrentResult: testRes1,
				},
				{
					ID:            2,
					VoteID:        1,
					Name:          "test variant 2",
					CurrentResult: testRes2,
				},
				{
					ID:            3,
					VoteID:        1,
					Name:          "test variant 3",
					CurrentResult: testRes3,
				},
			},
		},
	}
}
