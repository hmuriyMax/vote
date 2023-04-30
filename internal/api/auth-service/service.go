package auth_service

import (
	desc "github.com/hmuriyMax/vote/internal/pb/auth_service"
	"github.com/hmuriyMax/vote/internal/service/auth"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer

	authService *auth.Service
}

func NewImplementation(authService *auth.Service) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
