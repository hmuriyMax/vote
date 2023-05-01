package auth_service

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	auth "github.com/hmuriyMax/vote/internal/pb/auth_service"
	authService "github.com/hmuriyMax/vote/internal/service/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	if err := validateAuthRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user, token, err := i.authService.Login(ctx, req.GetLogin(), req.GetPass())
	if err != nil {
		if errors.Is(err, authService.ErrNoUser) || errors.Is(err, authService.ErrWrongPassword) {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &auth.AuthResponse{
		Status: auth.AuthResponse_Success,
		User: &auth.UserInfo{
			UserID:   user.ID,
			Username: user.Name,
			Token: &auth.Token{
				Token:   token.Token,
				Expires: token.ExpiresAt.Unix(),
			},
		},
	}, nil
}

func (i *Implementation) Register(ctx context.Context, req *auth.RegRequest) (*auth.RegResponse, error) {
	if err := validateRegRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user, token, err := i.authService.Register(ctx, req.GetUsername(), req.GetLogin(), req.GetPass())
	if err != nil {
		if errors.Is(err, authService.ErrAccountAlreadyExist) {
			return &auth.RegResponse{Status: auth.RegResponse_AlreadyExists}, nil
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &auth.RegResponse{
		Status: auth.RegResponse_Success,
		User: &auth.UserInfo{
			UserID:   user.ID,
			Username: user.Name,
			Token: &auth.Token{
				Token:   token.Token,
				Expires: token.ExpiresAt.Unix(),
			},
		},
	}, nil
}

func validateAuthRequest(req *auth.AuthRequest) error {
	return validation.ValidateStruct(
		validation.Field(&req.Login, validation.Required, validation.Length(3, 20)),
		validation.Field(&req.Pass, validation.Required),
	)
}

func validateRegRequest(req *auth.RegRequest) error {
	return validation.ValidateStruct(
		validation.Field(&req.Username, validation.Required, validation.Length(3, 20)),
		validation.Field(&req.Login, validation.Required, validation.Length(3, 20)),
		validation.Field(&req.Pass, validation.Required),
	)
}
