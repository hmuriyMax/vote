package vote

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"github.com/hmuriyMax/vote/internal/service/auth"
	"github.com/hmuriyMax/vote/internal/service/cypher"
)

type Service struct {
	repo          repo.VoteRepository
	authenticator *auth.Service
	cypherService *cypher.Service
}

func NewVoteService(
	cypherService *cypher.Service,
	authService *auth.Service,
	repo repo.VoteRepository,
) *Service {
	return &Service{
		cypherService: cypherService,
		authenticator: authService,
		repo:          repo,
	}
}

func (s *Service) GetVotesByUserID(ctx context.Context, userID int64) (model.Votes, error) {
	return s.repo.GetVotesByUserID(ctx, userID)
}

func (s *Service) GetVoteInfoByID(ctx context.Context, id int64) (*model.Vote, error) {
	return s.repo.GetVoteInfoByID(ctx, id)
}

func (s *Service) GetVoteVariantsByID(ctx context.Context, id int64) (model.Variants, error) {
	return s.repo.GetVariantsByVoteID(ctx, id)
}
