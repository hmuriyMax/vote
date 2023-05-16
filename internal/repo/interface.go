package repo

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
)

type AuthRepository interface {
	SelectUserByLogin(ctx context.Context, login string) (*model.UserAuth, error)
	SelectUserByID(ctx context.Context, id int64) (*model.UserAuth, error)
	InsertUser(ctx context.Context, user *model.UserAuth) (int64, error)
	NewToken(ctx context.Context, userID int64) (*model.Token, error)
	AssertToken(ctx context.Context, userID int64, token string) (bool, error)
}

type VoteRepository interface {
	GetVotesByUserID(ctx context.Context, userID int64) (model.Votes, error)
	GetVoteInfoByID(ctx context.Context, voteID int64) (*model.Vote, error)
	GetVariantsByVoteID(ctx context.Context, voteID int64) ([]*model.Variant, error)
	IncrementVote(ctx context.Context, variantID int64) (int64, error)
	AssertVote(ctx context.Context, voteID int64, variantID int64) error
}
