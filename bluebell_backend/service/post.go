package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/model"
	"bluebell_backend/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {
	// 生成post_id
	p.ID, err = snowflake.GetID()
	if err != nil {
		err = mysql.ErrorInvalidID
		return
	}
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	return redis.CreatePost(p.ID)
}

func GetPostDetial(id uint64) (data *model.ApiPostDetail, err error) {
	// 查询并拼接
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(id) failed", zap.Uint64("post_id", id), zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error(" mysql.GetUserByID(id) failed", zap.Uint64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error(" mysql.GetCommunityDetailByID(id) failed", zap.Uint64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	data = &model.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page, size int64) (data []*model.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	data = make([]*model.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error(" mysql.GetUserByID(id) failed", zap.Uint64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error(" mysql.GetCommunityDetailByID(id) failed", zap.Uint64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &model.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
