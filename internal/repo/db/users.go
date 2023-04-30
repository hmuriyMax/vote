package db

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"time"
)

func (r *Postgres) SelectUser(ctx context.Context, user string) (*model.UserAuth, error) {
	return &model.UserAuth{
		ID:            0,
		Name:          "test",
		Password:      "",
		IsBlocked:     false,
		LastLoginTime: time.Now(),
	}, nil
}

func (r *Postgres) InsertUser(ctx context.Context, user *model.UserAuth) (int64, error) {
	return 0, nil
}

func (r *Postgres) NewToken(ctx context.Context, userID int64) (*model.Token, error) {
	return &model.Token{
		ID:        0,
		UserID:    userID,
		Token:     "test",
		CreatedAt: time.Now(),
	}, nil
}

func (r *Postgres) AssertToken(ctx context.Context, userID int64, token string) (bool, error) {
	return token == "test", nil
}
