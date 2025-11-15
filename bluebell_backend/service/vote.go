package service

import (
	"bluebell_backend/dao/redis"
	"bluebell_backend/model"
	"strconv"

	"go.uber.org/zap"
)

// 投票限制：7 天
// 1. 到期之后将 redis 中 赞同票数 和 反对票数 存在 mysql 中
// 2. 到期之后删除 KeyPostVotedZSetPrefix
func VoteForPost(userID uint64, p *model.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Uint64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.FormatUint(userID, 10), p.PostID, float64(p.Direction))
}
