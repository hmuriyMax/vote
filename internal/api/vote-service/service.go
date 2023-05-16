package vote_service

import (
	desc "github.com/hmuriyMax/vote/internal/pb/vote_service"
	"github.com/hmuriyMax/vote/internal/service/auth"
	"github.com/hmuriyMax/vote/internal/service/cypher"
	"github.com/hmuriyMax/vote/internal/service/vote"
)

type Implementation struct {
	desc.UnimplementedVoteServiceServer

	voteService   *vote.Service
	authService   *auth.Service
	cypherService *cypher.Service
}

func NewImplementation(
	voteService *vote.Service,
	authService *auth.Service,
	cypherService *cypher.Service,
) *Implementation {
	return &Implementation{
		voteService:   voteService,
		authService:   authService,
		cypherService: cypherService,
	}
}
