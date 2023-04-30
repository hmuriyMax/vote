package repo

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
)

type Repository interface {
	SelectUser(ctx context.Context, user string) (*model.UserAuth, error)
	InsertUser(ctx context.Context, user *model.UserAuth) (int64, error)
	NewToken(ctx context.Context, userID int64) (*model.Token, error)
	AssertToken(ctx context.Context, userID int64, token string) (bool, error)
}
