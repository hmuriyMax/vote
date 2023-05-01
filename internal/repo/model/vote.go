package model

import (
	pb "github.com/hmuriyMax/vote/internal/pb/vote_service"
	"time"
)

type (
	Vote struct {
		ID         int64     `db:"id"`
		Name       string    `db:"name"`
		FinishesAt time.Time `db:"finishes"`
	}
	VoteUser struct {
		ID     int64 `db:"id"`
		VoteID int64 `db:"vote_id"`
		UserID int64 `db:"voter_id"`
	}
	Variant struct {
		ID            int64  `db:"id"`
		VoteID        int64  `db:"vote_id"`
		Name          string `db:"name"`
		CurrentResult int64  `db:"current_result"`
	}

	Votes    []*Vote
	Variants []*Variant
)

func (v *Vote) ToPB() *pb.ShortVoteInfo {
	return &pb.ShortVoteInfo{
		Id:       v.ID,
		Name:     v.Name,
		Finishes: v.FinishesAt.Unix(),
	}
}

func (v Votes) ToPB() []*pb.ShortVoteInfo {
	pbVotes := make([]*pb.ShortVoteInfo, len(v))
	for i, v := range v {
		pbVotes[i] = v.ToPB()
	}
	return pbVotes
}

func (v Variant) ToPB() *pb.Variant {
	return &pb.Variant{
		Id:   v.ID,
		Name: v.Name,
	}
}

func (v Variants) ToPB() []*pb.Variant {
	pbVariants := make([]*pb.Variant, len(v))
	for i, v := range v {
		pbVariants[i] = v.ToPB()
	}
	return pbVariants
}
