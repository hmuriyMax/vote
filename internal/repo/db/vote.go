package db

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"time"
)

var (
	testRes1 int64
	testRes2 int64
	testRes3 int64
)

func (r *Postgres) GetVotesByUserID(ctx context.Context, userID int64) (model.Votes, error) {
	return []*model.Vote{
		{
			ID:         1,
			Name:       "test vote",
			FinishesAt: time.Date(2033, 04, 30, 23, 31, 24, 12, time.Local),
		},
	}, nil
}

func (r *Postgres) GetVoteInfoByID(ctx context.Context, id int64) (*model.Vote, error) {
	votes, _ := r.GetVotesByUserID(ctx, 0)
	return votes[0], nil
}

func (r *Postgres) GetVariantsByVoteID(ctx context.Context, voteID int64) ([]*model.Variant, error) {
	return []*model.Variant{
		{
			ID:            1,
			VoteID:        voteID,
			Name:          "test variant 1",
			CurrentResult: testRes1,
		},
		{
			ID:            2,
			VoteID:        voteID,
			Name:          "test variant 2",
			CurrentResult: testRes2,
		},
		{
			ID:            3,
			VoteID:        voteID,
			Name:          "test variant 3",
			CurrentResult: testRes3,
		},
	}, nil
}

func (r *Postgres) IncrementVote(ctx context.Context, variantID int64) (int64, error) {
	switch variantID {
	case 1:
		testRes1++
	case 2:
		testRes2++
	case 3:
		testRes3++
	}
	return -1, nil
}

func (r *Postgres) AssertVote(ctx context.Context, voteID int64, variantID int64) error {
	return nil
}
