package vote_service

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hmuriyMax/vote/internal/pb/vote_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) Vote(ctx context.Context, req *vote_service.VoteRequest) (*vote_service.VoteResponse, error) {
	if err := validateVoteRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Errorf(codes.Unimplemented, "vote method unimplemented")
}

func validateVoteRequest(req *vote_service.VoteRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserId, validation.Required, validation.Min(1)),
		validation.Field(&req.CypherVote, validation.Required),
	)
}
