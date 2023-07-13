package models

import "time"

type CommunityList struct {
	CommunityID   int64  `json:"id" db:"community_id"`
	CommunityName string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	CommunityID   int64     `json:"id" db:"community_id"`
	CommunityName string    `json:"name" db:"community_name"`
	Introduction  string    `json:"intro" db:"introduction"`
	CreateTime    time.Time `json:"create" db:"create_time"`
	UpdateTime    time.Time `json:"update" db:"update_time"`
}
