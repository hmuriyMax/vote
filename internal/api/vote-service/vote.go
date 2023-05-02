package vote_service

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	vote "github.com/hmuriyMax/vote/internal/pb/vote_service"
	"github.com/hmuriyMax/vote/internal/service/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Implementation) Vote(ctx context.Context, req *vote.VoteRequest) (*vote.VoteResponse, error) {
	if err := validateVoteRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.authMiddle(ctx, req.GetAuth()); err != nil {
		return nil, err
	}

	return nil, status.Error(codes.Unimplemented, "/vote is not implemented")
}

func (s *Implementation) GetVotesForUser(ctx context.Context, req *vote.GetVotesRequest) (*vote.GetVotesResponse, error) {
	if err := validateGetVotesRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.authMiddle(ctx, req.GetAuth()); err != nil {
		return nil, err
	}

	votes, err := s.voteService.GetVotesByUserID(ctx, req.GetAuth().UserID)
	if err != nil {
		return nil, err
	}

	return &vote.GetVotesResponse{
		Votes: votes.ToPB(),
	}, nil
}

func (s *Implementation) GetVoteInfo(ctx context.Context, req *vote.GetVoteInfoRequest) (*vote.GetVoteInfoResponse, error) {
	if err := validateGetVoteInfoRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.authMiddle(ctx, req.GetAuth()); err != nil {
		return nil, err
	}

	info, err := s.voteService.GetVoteInfoByID(ctx, req.GetVoteId())
	if err != nil {
		return nil, err
	}

	vars, err := s.voteService.GetVoteVariantsByID(ctx, req.GetVoteId())
	if err != nil {
		return nil, err
	}

	return &vote.GetVoteInfoResponse{
		Info: &vote.ExtendedVoteInfo{
			Short:    info.ToPB(),
			Variants: vars.ToPB(),
		},
	}, nil
}

func validateVoteRequest(req *vote.VoteRequest) error {
	err := validation.ValidateStruct(req,
		validation.Field(&req.VoteId, validation.Required, validation.Min(1)),
		validation.Field(&req.CypherVote, validation.Required),
	)
	if err != nil {
		return err
	}
	return validateAuthInfo(req.GetAuth())
}

func validateGetVotesRequest(req *vote.GetVotesRequest) error {
	if err := validation.ValidateStruct(req); err != nil {
		return err
	}
	return validateAuthInfo(req.GetAuth())
}

func validateGetVoteInfoRequest(req *vote.GetVoteInfoRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.VoteId, validation.Min(1)),
	); err != nil {
		return err
	}
	return validateAuthInfo(req.GetAuth())
}

func validateAuthInfo(auth *vote.AuthInfo) error {
	return validation.ValidateStruct(auth,
		validation.Field(&auth.Token, validation.Required),
		validation.Field(&auth.UserID, validation.Min(1)),
	)
}

func (s *Implementation) authMiddle(ctx context.Context, authInfo *vote.AuthInfo) error {
	err := s.authService.CheckToken(ctx, authInfo.GetUserID(), authInfo.GetToken())
	if err != nil {
		if errors.Is(err, auth.ErrNoUser) {
			return status.Errorf(codes.NotFound, "not found user %d: %v", authInfo.GetUserID(), err)
		}
		if errors.Is(err, auth.ErrUserBlocked) {
			return status.Errorf(codes.PermissionDenied, "user %d is blocked: %v", authInfo.GetUserID(), err)
		}
		if errors.Is(err, auth.ErrUnauthorized) {
			return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}
		return status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return nil
}
