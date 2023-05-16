package vote

import (
	"context"
	"github.com/hmuriyMax/vote/internal/repo"
	"github.com/hmuriyMax/vote/internal/repo/model"
	"github.com/hmuriyMax/vote/internal/service/auth"
)

type Service struct {
	repo          repo.VoteRepository
	authenticator *auth.Service
}

func NewVoteService(
	authService *auth.Service,
	repo repo.VoteRepository,
) *Service {
	return &Service{
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

func (s *Service) Vote(ctx context.Context, variantID int64) (int64, error) {
	return s.repo.IncrementVote(ctx, variantID)
}

func (s *Service) AssertVoteVariant(ctx context.Context, voteID int64, variantID int64) error {
	return s.repo.AssertVote(ctx, voteID, variantID)
}
