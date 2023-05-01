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
	repository repo.AuthRepository
	TokenTTL   time.Duration
}

func NewAuthService(repo repo.AuthRepository) *Service {
	return &Service{
		repository: repo,
		TokenTTL:   24 * time.Hour,
	}
}

func (d *Service) Login(ctx context.Context, login string, password string) (*model.UserAuth, *model.Token, error) {
	user, err := d.repository.SelectUserByLogin(ctx, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, ErrNoUser
		}
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	hashPass := sha256.Sum256([]byte(user.Password))
	if string(hashPass[:]) != password {
		return nil, nil, ErrWrongPassword
	}

	token, err := d.repository.NewToken(ctx, user.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token: %w", err)
	}
	return user, token, nil
}

func (d *Service) Register(ctx context.Context, username, login, password string) (*model.UserAuth, *model.Token, error) {
	_, err := d.repository.SelectUserByLogin(ctx, username)
	if err == nil {
		return nil, nil, ErrAccountAlreadyExist
	}
	if err != sql.ErrNoRows {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	hashPass := sha256.Sum256([]byte(password))
	newUser := &model.UserAuth{
		Name:      username,
		Login:     login,
		Password:  string(hashPass[:]),
		IsBlocked: false,
	}

	userId, err := d.repository.InsertUser(ctx, newUser)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to insert user: %w", err)
	}

	token, err := d.repository.NewToken(ctx, userId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token: %w", err)
	}
	return newUser, token, nil
}

func (d *Service) CheckToken(ctx context.Context, userID int64, token string) error {
	user, err := d.repository.SelectUserByID(ctx, userID)
	if err != nil {
		return ErrNoUser
	}
	if user.IsBlocked {
		return ErrUserBlocked
	}
	success, err := d.repository.AssertToken(ctx, userID, token)
	if err != nil {
		return fmt.Errorf("failed to check token: %w", err)
	}
	if !success {
		return ErrUnauthorized
	}
	return nil
}
