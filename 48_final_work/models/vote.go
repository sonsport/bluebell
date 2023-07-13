package models

type VoteData struct {
	PostID string `json:"post_id" binding:"required"`
	Vote   int8   `json:"vote" binding:"oneof=1 0 -1"`
}
