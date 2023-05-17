package vote_service

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
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
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var (
		keyBytes    = req.GetPublicKey()
		receivedKey rsa.PublicKey
		serverKey   = s.cypherService.GetPublicKey()
	)
	err := json.Unmarshal(keyBytes, &receivedKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unmarshal public key: %w", err)
	}

	if !serverKey.Equal(&receivedKey) {
		return nil, status.Errorf(codes.PermissionDenied, "invalid public key")
	}

	var voteVariant vote.VoteVariantPair
	err = s.cypherService.DecryptProto(req.GetCypherVote(), &voteVariant)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Errorf("failed to decrypt vote: %w", err).Error())
	}

	err = s.voteService.AssertVoteVariant(ctx, voteVariant.VoteID, voteVariant.VariantID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Errorf("invalid vote variant: %w", err).Error())
	}

	_, err = s.voteService.Vote(ctx, voteVariant.VariantID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to vote: %w", err).Error())
	}

	return &vote.VoteResponse{
		Status: vote.VoteResponse_Accepted,
	}, nil
}

func (s *Implementation) GetVotesForUser(ctx context.Context, req *vote.GetVotesRequest) (*vote.GetVotesResponse, error) {
	if err := validateGetVotesRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.authMiddle(ctx, req.GetAuth()); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
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
		return nil, status.Error(codes.Unauthenticated, err.Error())
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

func (s *Implementation) GetVotePublicKey(ctx context.Context, req *vote.Empty) (*vote.GetVotePublicKeyResponse, error) {
	publicKey := s.cypherService.GetPublicKey()
	bytes, err := json.Marshal(publicKey)
	if err != nil {
		return nil, fmt.Errorf("marshal public key: %w", err)
	}
	return &vote.GetVotePublicKeyResponse{
		PublicKeyJson: bytes,
	}, nil
}

func validateVoteRequest(req *vote.VoteRequest) error {
	err := validation.ValidateStruct(req,
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
