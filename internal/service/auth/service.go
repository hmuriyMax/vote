package auth

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"github.com/hmuriyMax/vote/internal/repo"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"time"
)

type Service struct {
	repository repo.Repository
	TokenTTL   time.Duration
}

func NewAuthService(repo repo.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (d *Service) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := d.repository.SelectUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNoUser
		}
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	hashPass := sha256.Sum256([]byte(user.Password))
	if string(hashPass[:]) != password {
		return "", ErrWrongPassword
	}

	token, err := d.repository.NewToken(ctx, user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}
	return token.Token, nil
}

func (d *Service) Register(ctx context.Context, username string, password string) (string, error) {
	_, err := d.repository.SelectUser(ctx, username)
	if err == nil {
		return "", ErrAccountAlreadyExist
	}
	if err != sql.ErrNoRows {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	hashPass := sha256.Sum256([]byte(password))
	userId, err := d.repository.InsertUser(ctx, &model.UserAuth{
		Name:      username,
		Password:  string(hashPass[:]),
		IsBlocked: false,
	})
	if err != nil {
		return "", fmt.Errorf("failed to insert user: %w", err)
	}

	token, err := d.repository.NewToken(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}
	return token.Token, nil
}
