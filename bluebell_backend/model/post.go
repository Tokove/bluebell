package model

import "time"

type Post struct {
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content"`
	ID          uint64    `json:"id" db:"post_id"`
	AuthorID    uint64    `json:"author_id" db:"author_id"`
	CommunityID uint64    `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}
