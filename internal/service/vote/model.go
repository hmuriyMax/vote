package vote

type Vote struct {
	VoteID    int64 `json:"vote_id" db:"vote_id"`
	VariantID int64 `json:"variant_id" db:"variant_id"`
}
