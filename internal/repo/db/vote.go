package db

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"log"
)

var (
	testRes1 int64
	testRes2 int64
	testRes3 int64
)

func (r *InMemory) GetVotesByUserID(ctx context.Context, userID int64) (model.Votes, error) {
	return r.votes, nil
}

func (r *InMemory) GetVoteInfoByID(ctx context.Context, id int64) (*model.Vote, error) {
	votes, _ := r.GetVotesByUserID(ctx, 0)
	return votes[0], nil
}

func (r *InMemory) GetVariantsByVoteID(ctx context.Context, voteID int64) ([]*model.Variant, error) {
	return r.variants[voteID], nil
}

func (r *InMemory) IncrementVote(ctx context.Context, variantID int64) (int64, error) {
	switch variantID {
	case 1:
		testRes1++
	case 2:
		testRes2++
	case 3:
		testRes3++
	}
	log.Printf("received new vote, current result: %d: %d, %d: %d, %d: %d", 1, testRes1, 2, testRes2, 3, testRes3)
	return -1, nil
}

func (r *InMemory) AssertVote(ctx context.Context, voteID int64, variantID int64) error {
	return nil
}
