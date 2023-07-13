package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	Author           string `json:"author"`
	VoteData         int    `json:"vote"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}

type ApiTableDetail struct {
	ID         int64     `json:"id" db:"id"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	Username   string    `json:"username" db:"username"`
}
