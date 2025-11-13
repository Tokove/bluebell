package mysql

import (
	"bluebell_backend/model"
	"database/sql"
)

func CreatePost(post *model.Post) error {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return err
}

func GetPostDetailByID(id uint64) (p *model.Post, err error) {
	p = new(model.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ? `

	if err = db.Get(p, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

func GetPostList(page, size int64) (posts []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post limit ?, ?`
	// 从第几条读，读多少
	posts = make([]*model.Post, 0, size)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
