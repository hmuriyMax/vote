package vote

import "github.com/hmuriyMax/vote/internal/service/cypher"

type Service struct {
	cypherService *cypher.Service
}

func NewVoteService(cypherService *cypher.Service) *Service {
	return &Service{
		cypherService: cypherService,
	}
}

func (s *Service) CountVote(vote *Vote) error {
	// TODO: implement
	return nil
}
