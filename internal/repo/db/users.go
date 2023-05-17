package db

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"time"
)

func (r *InMemory) SelectUserByLogin(ctx context.Context, login string) (*model.UserAuth, error) {
	return &model.UserAuth{
		ID:            1,
		Name:          "test",
		Login:         "test@vote.ru",
		Password:      "",
		IsBlocked:     false,
		LastLoginTime: time.Now(),
	}, nil
}

func (r *InMemory) SelectUserByID(ctx context.Context, id int64) (*model.UserAuth, error) {
	return &model.UserAuth{
		ID:            1,
		Name:          "test",
		Login:         "test@vote.ru",
		Password:      "",
		IsBlocked:     false,
		LastLoginTime: time.Now(),
	}, nil
}

func (r *InMemory) InsertUser(ctx context.Context, user *model.UserAuth) (int64, error) {
	return 1, nil
}

func (r *InMemory) NewToken(ctx context.Context, userID int64) (*model.Token, error) {
	return &model.Token{
		ID:        1,
		UserID:    userID,
		Token:     "test",
		ExpiresAt: time.Now().Add(time.Hour * 24 * 365),
	}, nil
}

func (r *InMemory) AssertToken(ctx context.Context, userID int64, token string) (bool, error) {
	return token == "test", nil
}
